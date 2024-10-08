#pragma once
#include <contrib/ydb/core/base/appdata.h>

#include <contrib/ydb/public/api/protos/ydb_table.pb.h>
#include <contrib/ydb/services/metadata/service.h>

namespace NKikimr::NMetadata::NModifications {

class TBaseObject {
public:
    static Ydb::Table::CreateTableRequest AddHistoryTableScheme(const Ydb::Table::CreateTableRequest& baseScheme, const TString& tableName);

};

template <class TDerived>
class TObject: public TBaseObject {
public:
    static Ydb::Table::CreateTableRequest AddHistoryTableScheme(const Ydb::Table::CreateTableRequest& baseScheme) {
        return TBaseObject::AddHistoryTableScheme(baseScheme, TDerived::GetBehaviour()->GetStorageHistoryTablePath());
    }
};

}
