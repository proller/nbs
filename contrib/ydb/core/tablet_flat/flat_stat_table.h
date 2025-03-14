#pragma once

#include "flat_part_laid.h"
#include "flat_stat_part.h"
#include "flat_table_subset.h"

#include <util/generic/queue.h>
#include <util/generic/hash_set.h>

#include <contrib/ydb/core/scheme/scheme_tablecell.h>

namespace NKikimr {
namespace NTable {

// Iterates over all parts and maintains total row count and data size
class TStatsIterator {
public:
    explicit TStatsIterator(TIntrusiveConstPtr<TKeyCellDefaults> keyColumns)
        : KeyColumns(keyColumns)
        , Heap(TIterKeyGreater{ this })
    {}

    void Add(THolder<TScreenedPartIndexIterator> pi) {
        Y_ABORT_UNLESS(pi->IsValid());
        Iterators.PushBack(std::move(pi));
        TScreenedPartIndexIterator* it = Iterators.back();
        Heap.push(it);
    }

    EReady Next(TPartDataStats& stats) {
        ui64 lastRowCount = stats.RowCount;
        ui64 lastDataSize = stats.DataSize.Size;

        TCellsStorage cellsStorage;

        while (!Heap.empty()) {
            TScreenedPartIndexIterator* it = Heap.top();
            Heap.pop();

            // makes key copy
            cellsStorage.Reset({it->GetCurrentKey().Columns, it->GetCurrentKey().ColumnCount});
            TDbTupleRef key(KeyColumns->BasicTypes().data(), cellsStorage.GetCells().data(), cellsStorage.GetCells().size());

            auto ready = it->Next(stats);
            if (ready == EReady::Page) {
                return ready;
            } else if (ready == EReady::Data) {
                Heap.push(it);
            }

            // guarantees that all results will be different
            while (!Heap.empty() && CompareKeys(key, Heap.top()->GetCurrentKey()) == 0) {
                it = Heap.top();
                Heap.pop();

                ready = it->Next(stats);
                if (ready == EReady::Page) {
                    return ready;
                } else if (ready == EReady::Data) {
                    Heap.push(it);
                }
            }

            if (stats.RowCount != lastRowCount && stats.DataSize.Size != lastDataSize) {
                break;
            }
        }

        return Heap.empty() ? EReady::Gone : EReady::Data;
    }

    TDbTupleRef GetCurrentKey() const {
        Y_ABORT_UNLESS(!Heap.empty());
        return Heap.top()->GetCurrentKey();
    }

private:
    int CompareKeys(const TDbTupleRef& a, const TDbTupleRef& b) const noexcept {
        return ComparePartKeys(a.Cells(), b.Cells(), *KeyColumns);
    }

    struct TIterKeyGreater {
        const TStatsIterator* Self;

        bool operator ()(const TScreenedPartIndexIterator* a, const TScreenedPartIndexIterator* b) const {
            return Self->CompareKeys(a->GetCurrentKey(), b->GetCurrentKey()) > 0;
        }
    };

    TIntrusiveConstPtr<TKeyCellDefaults> KeyColumns;
    THolderVector<TScreenedPartIndexIterator> Iterators;
    TPriorityQueue<TScreenedPartIndexIterator*, TSmallVec<TScreenedPartIndexIterator*>, TIterKeyGreater> Heap;
};

struct TBucket {
    TString EndKey;
    ui64 Value;
};

using THistogram = TVector<TBucket>;

struct TStats {
    ui64 RowCount = 0;
    TPartDataSize DataSize = { };
    TPartDataSize IndexSize = { };
    THistogram RowCountHistogram;
    THistogram DataSizeHistogram;

    void Clear() {
        RowCount = 0;
        DataSize = { };
        IndexSize = { };
        RowCountHistogram.clear();
        DataSizeHistogram.clear();
    }

    void Swap(TStats& other) {
        std::swap(RowCount, other.RowCount);
        std::swap(DataSize, other.DataSize);
        std::swap(IndexSize, other.IndexSize);
        RowCountHistogram.swap(other.RowCountHistogram);
        DataSizeHistogram.swap(other.DataSizeHistogram);
    }
};

class TKeyAccessSample {
public:
    enum class EAccessKind {
        Read = 1,
        Update = 2,
        Delete = 3
    };

    using TSample = TVector<std::pair<TString, EAccessKind>>;

public:
    explicit TKeyAccessSample(ui64 sampleCount = 100)
        : SampleCount(sampleCount)
        , TotalCount(0)
    {}

    void Add(TArrayRef<const TCell> key, EAccessKind accessKind = EAccessKind::Read) {
        ui64 idx = TotalCount;
        ++TotalCount;
        if (idx >= SampleCount) {
            idx = RandomNumber<ui64>(TotalCount);
        }

        if (idx >= SampleCount) {
            return;
        }

        TString serializedKey = TSerializedCellVec::Serialize(key);
        ++KeyRefCount[serializedKey];

        if (Sample.size() < SampleCount) {
            Sample.emplace_back(std::make_pair(serializedKey, accessKind));
            return;
        }

        TString old = Sample[idx].first;
        auto oit = KeyRefCount.find(old);
        Y_ABORT_UNLESS(oit != KeyRefCount.end());
        --oit->second;

        // Delete the key if this was the last reference
        if (oit->second == 0) {
            KeyRefCount.erase(oit);
        }

        Sample[idx] = std::make_pair(serializedKey, accessKind);
    }

    const TSample& GetSample() const {
        return Sample;
    }

    void Clear() {
        Sample.clear();
        TotalCount = 0;
        KeyRefCount.clear();
    }

private:
    TSample Sample;
    const ui64 SampleCount;
    ui64 TotalCount;
    // Store only unique keys and their ref counts to save memory
    THashMap<TString, ui64> KeyRefCount;
};

bool BuildStats(const TSubset& subset, TStats& stats, ui64 rowCountResolution, ui64 dataSizeResolution, IPages* env);
void GetPartOwners(const TSubset& subset, THashSet<ui64>& partOwners);

}}
