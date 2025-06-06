#pragma once
#include <contrib/ydb/core/tx/columnshard/engines/reader/read_metadata.h>
#include <contrib/ydb/core/tx/columnshard/engines/reader/read_context.h>
#include <contrib/ydb/core/tx/columnshard/engines/portions/column_record.h>
#include <contrib/ydb/core/tx/columnshard/blobs_reader/task.h>
#include <contrib/ydb/core/tx/columnshard/blob.h>
#include "source.h"

namespace NKikimr::NOlap::NPlainReader {

class TBlobsFetcherTask: public NBlobOperations::NRead::ITask {
private:
    using TBase = NBlobOperations::NRead::ITask;
    const std::shared_ptr<IDataSource> Source;
    const std::shared_ptr<IFetchingStep> Step;
    const std::shared_ptr<TSpecialReadContext> Context;

    virtual void DoOnDataReady(const std::shared_ptr<NResourceBroker::NSubscribe::TResourcesGuard>& resourcesGuard) override;
    virtual bool DoOnError(const TBlobRange& range, const IBlobsReadingAction::TErrorStatus& status) override;
public:
    TBlobsFetcherTask(const std::vector<std::shared_ptr<IBlobsReadingAction>>& readActions, const std::shared_ptr<IDataSource>& sourcePtr,
        const std::shared_ptr<IFetchingStep>& step, const std::shared_ptr<TSpecialReadContext>& context, const TString& taskCustomer, const TString& externalTaskId)
        : TBase(readActions, taskCustomer, externalTaskId) 
        , Source(sourcePtr)
        , Step(step)
        , Context(context)
    {

    }
};

}
