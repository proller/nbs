#pragma once

#include <contrib/ydb/core/formats/arrow/arrow_helpers.h>
#include <contrib/ydb/core/formats/arrow/common/validation.h>
#include <contrib/ydb/core/formats/arrow/serializer/abstract.h>
#include <contrib/ydb/core/tx/columnshard/common/scalars.h>
#include <contrib/libs/apache/arrow/cpp/src/arrow/util/compression.h>
#include <util/generic/set.h>
#include <util/generic/hash_set.h>

namespace NKikimr::NOlap {

class TTierInfo {
private:
    YDB_READONLY_DEF(TString, Name);
    YDB_READONLY_DEF(TString, EvictColumnName);
    YDB_READONLY_DEF(TDuration, EvictDuration);

    ui32 TtlUnitsInSecond;
    YDB_READONLY_DEF(std::optional<NArrow::NSerialization::TSerializerContainer>, Serializer);
public:
    TTierInfo(const TString& tierName, TDuration evictDuration, const TString& column, ui32 unitsInSecond = 0)
        : Name(tierName)
        , EvictColumnName(column)
        , EvictDuration(evictDuration)
        , TtlUnitsInSecond(unitsInSecond)
    {
        Y_ABORT_UNLESS(!!Name);
        Y_ABORT_UNLESS(!!EvictColumnName);
    }

    TInstant GetEvictInstant(const TInstant now) const {
        return now - EvictDuration;
    }

    TTierInfo& SetSerializer(const NArrow::NSerialization::TSerializerContainer& value) {
        Serializer = value;
        return *this;
    }

    std::shared_ptr<arrow::Field> GetEvictColumn(const std::shared_ptr<arrow::Schema>& schema) const {
        return schema->GetFieldByName(EvictColumnName);
    }

    std::optional<TInstant> ScalarToInstant(const std::shared_ptr<arrow::Scalar>& scalar) const;

    static std::shared_ptr<TTierInfo> MakeTtl(const TDuration evictDuration, const TString& ttlColumn, ui32 unitsInSecond = 0) {
        return std::make_shared<TTierInfo>("TTL", evictDuration, ttlColumn, unitsInSecond);
    }

    TString GetDebugString() const {
        TStringBuilder sb;
        sb << "name=" << Name << ";duration=" << EvictDuration << ";column=" << EvictColumnName << ";serializer=";
        if (Serializer) {
            sb << Serializer->DebugString();
        } else {
            sb << "NOT_SPECIFIED(Default)";
        }
        sb << ";";
        return sb;
    }
};

class TTierRef {
public:
    TTierRef(const std::shared_ptr<TTierInfo>& tierInfo)
        : Info(tierInfo)
    {
        Y_ABORT_UNLESS(tierInfo);
    }

    bool operator < (const TTierRef& b) const {
        if (Info->GetEvictDuration() > b.Info->GetEvictDuration()) {
            return true;
        } else if (Info->GetEvictDuration() == b.Info->GetEvictDuration()) {
            return Info->GetName() > b.Info->GetName(); // add stability: smaller name is hotter
        }
        return false;
    }

    bool operator == (const TTierRef& b) const {
        return Info->GetEvictDuration() == b.Info->GetEvictDuration()
            && Info->GetName() == b.Info->GetName();
    }

    const TTierInfo& Get() const {
        return *Info;
    }

    std::shared_ptr<TTierInfo> GetPtr() const {
        return Info;
    }

private:
    std::shared_ptr<TTierInfo> Info;
};

class TTiering {
    using TTiersMap = THashMap<TString, std::shared_ptr<TTierInfo>>;
    TTiersMap TierByName;
    TSet<TTierRef> OrderedTiers;
public:

    std::shared_ptr<TTierInfo> Ttl;

    const TTiersMap& GetTierByName() const {
        return TierByName;
    }

    const TSet<TTierRef>& GetOrderedTiers() const {
        return OrderedTiers;
    }

    bool HasTiers() const {
        return !OrderedTiers.empty();
    }

    void Add(const std::shared_ptr<TTierInfo>& tier) {
        if (HasTiers()) {
            // TODO: support different ttl columns
            Y_ABORT_UNLESS(tier->GetEvictColumnName() == OrderedTiers.begin()->Get().GetEvictColumnName());
        }

        TierByName.emplace(tier->GetName(), tier);
        OrderedTiers.emplace(tier);
    }

    TString GetHottestTierName() const {
        if (OrderedTiers.size()) {
            return OrderedTiers.rbegin()->Get().GetName(); // hottest one
        }
        return {};
    }

    std::optional<NArrow::NSerialization::TSerializerContainer> GetSerializer(const TString& name) const {
        auto it = TierByName.find(name);
        if (it != TierByName.end()) {
            Y_ABORT_UNLESS(!name.empty());
            return it->second->GetSerializer();
        }
        return {};
    }

    THashSet<TString> GetTtlColumns() const {
        THashSet<TString> out;
        if (Ttl) {
            out.insert(Ttl->GetEvictColumnName());
        }
        for (auto& [tierName, tier] : TierByName) {
            out.insert(tier->GetEvictColumnName());
        }
        return out;
    }

    TString GetDebugString() const {
        TStringBuilder sb;
        if (Ttl) {
            sb << Ttl->GetDebugString() << "; ";
        }
        for (auto&& i : OrderedTiers) {
            sb << i.Get().GetDebugString() << "; ";
        }
        return sb;
    }
};

}
