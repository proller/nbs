#pragma once
#include <contrib/ydb/core/tx/columnshard/common/portion.h>
#include <contrib/ydb/core/tx/columnshard/common/snapshot.h>
#include <contrib/ydb/core/formats/arrow/replace_key.h>
#include <contrib/ydb/core/formats/arrow/special_keys.h>
#include <contrib/ydb/core/protos/tx_columnshard.pb.h>
#include <contrib/ydb/library/accessor/accessor.h>
#include <util/stream/output.h>

namespace NKikimr::NOlap {

struct TIndexInfo;

struct TPortionMeta {
private:
    std::shared_ptr<NArrow::TFirstLastSpecialKeys> ReplaceKeyEdges; // first and last PK rows
    YDB_ACCESSOR_DEF(TString, TierName);
public:
    using EProduced = NPortion::EProduced;

    std::optional<NArrow::TReplaceKey> IndexKeyStart;
    std::optional<NArrow::TReplaceKey> IndexKeyEnd;

    std::optional<TSnapshot> RecordSnapshotMin;
    std::optional<TSnapshot> RecordSnapshotMax;
    EProduced Produced{EProduced::UNSPECIFIED};
    ui32 FirstPkColumn = 0;

    bool IsChunkWithPortionInfo(const ui32 columnId, const ui32 chunkIdx) const {
        return columnId == FirstPkColumn && chunkIdx == 0;
    }

    bool DeserializeFromProto(const NKikimrTxColumnShard::TIndexPortionMeta& portionMeta, const TIndexInfo& indexInfo);

    std::optional<NKikimrTxColumnShard::TIndexPortionMeta> SerializeToProto(const ui32 columnId, const ui32 chunk) const;

    void FillBatchInfo(const NArrow::TFirstLastSpecialKeys& primaryKeys, const NArrow::TMinMaxSpecialKeys& snapshotKeys, const TIndexInfo& indexInfo);

    EProduced GetProduced() const {
        return Produced;
    }

    TString DebugString() const;

    friend IOutputStream& operator << (IOutputStream& out, const TPortionMeta& info) {
        out << info.DebugString();
        return out;
    }

    bool HasSnapshotMinMax() const {
        return !!RecordSnapshotMax && !!RecordSnapshotMin;
    }

    bool HasPrimaryKeyBorders() const {
        return !!IndexKeyStart && !!IndexKeyEnd;
    }
};

class TPortionAddress {
private:
    YDB_READONLY(ui64, PathId, 0);
    YDB_READONLY(ui64, PortionId, 0);
public:
    TPortionAddress(const ui64 pathId, const ui64 portionId)
        : PathId(pathId)
        , PortionId(portionId)
    {

    }

    TString DebugString() const;

    bool operator<(const TPortionAddress& item) const {
        return std::tie(PathId, PortionId) < std::tie(item.PathId, item.PortionId);
    }

    bool operator==(const TPortionAddress& item) const {
        return std::tie(PathId, PortionId) == std::tie(item.PathId, item.PortionId);
    }
};

} // namespace NKikimr::NOlap

template<>
struct THash<NKikimr::NOlap::TPortionAddress> {
    inline ui64 operator()(const NKikimr::NOlap::TPortionAddress& x) const noexcept {
        return CombineHashes(x.GetPortionId(), x.GetPathId());
    }
};

