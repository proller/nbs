#pragma once
#include "remove.h"
#include "write.h"
#include "read.h"
#include "gc.h"

#include <contrib/ydb/core/tx/columnshard/blobs_action/blob_manager_db.h>
#include <contrib/ydb/core/tx/columnshard/blobs_action/counters/storage.h>
#include <contrib/ydb/core/tx/columnshard/blobs_action/counters/remove_gc.h>
#include <contrib/ydb/library/accessor/accessor.h>

namespace NKikimr::NColumnShard {
class TTiersManager;
}

namespace NKikimr::NOlap {

class TCommonBlobsTracker: public IBlobInUseTracker {
private:
    // List of blobs that are used by in-flight requests
    THashMap<TUnifiedBlobId, i64> BlobsUseCount;
protected:
    virtual bool DoUseBlob(const TUnifiedBlobId& blobId) override;
    virtual bool DoFreeBlob(const TUnifiedBlobId& blobId) override;
public:
    virtual bool IsBlobInUsage(const NOlap::TUnifiedBlobId& blobId) const override;
    virtual void OnBlobFree(const TUnifiedBlobId& blobId) = 0;
};

class IBlobsStorageOperator {
private:
    YDB_READONLY_DEF(TString, StorageId);

    std::shared_ptr<IBlobsGCAction> CurrentGCAction;
    YDB_READONLY(bool, Stopped, false);
    std::shared_ptr<NBlobOperations::TStorageCounters> Counters;
protected:
    virtual std::shared_ptr<IBlobsDeclareRemovingAction> DoStartDeclareRemovingAction() = 0;
    virtual std::shared_ptr<IBlobsWritingAction> DoStartWritingAction() = 0;
    virtual std::shared_ptr<IBlobsReadingAction> DoStartReadingAction() = 0;
    virtual bool DoLoad(NColumnShard::IBlobManagerDb& dbBlobs) = 0;
    virtual bool DoStop() {
        return true;
    }

    virtual void DoOnTieringModified(const std::shared_ptr<NColumnShard::TTiersManager>& tiers) = 0;
    virtual TString DoDebugString() const {
        return "";
    }

    virtual std::shared_ptr<IBlobsGCAction> DoStartGCAction(const std::shared_ptr<NBlobOperations::TRemoveGCCounters>& counters) const = 0;
    std::shared_ptr<IBlobsGCAction> StartGCAction(const std::shared_ptr<NBlobOperations::TRemoveGCCounters>& counters) const {
        return DoStartGCAction(counters);
    }

public:
    IBlobsStorageOperator(const TString& storageId)
        : StorageId(storageId)
    {
        Counters = std::make_shared<NBlobOperations::TStorageCounters>(storageId);
    }

    void Stop();

    virtual std::shared_ptr<IBlobInUseTracker> GetBlobsTracker() const = 0;

    virtual ~IBlobsStorageOperator() = default;

    TString DebugString() const {
        return TStringBuilder() << "(storage_id=" << StorageId << ";details=(" << DoDebugString() << "))";
    }

    bool Load(NColumnShard::IBlobManagerDb& dbBlobs) {
        return DoLoad(dbBlobs);
    }
    void OnTieringModified(const std::shared_ptr<NColumnShard::TTiersManager>& tiers) {
        return DoOnTieringModified(tiers);
    }

    std::shared_ptr<IBlobsDeclareRemovingAction> StartDeclareRemovingAction(const TString& consumerId) {
        auto result = DoStartDeclareRemovingAction();
        result->SetCounters(Counters->GetConsumerCounter(consumerId)->GetRemoveDeclareCounters());
        return result;
    }
    std::shared_ptr<IBlobsWritingAction> StartWritingAction(const TString& consumerId) {
        auto result = DoStartWritingAction();
        result->SetCounters(Counters->GetConsumerCounter(consumerId)->GetWriteCounters());
        return result;
    }
    std::shared_ptr<IBlobsReadingAction> StartReadingAction(const TString& consumerId) {
        auto result = DoStartReadingAction();
        result->SetCounters(Counters->GetConsumerCounter(consumerId)->GetReadCounters());
        return result;
    }
    bool StartGC() {
        if (CurrentGCAction && CurrentGCAction->IsInProgress()) {
            return false;
        }
        if (Stopped) {
            return false;
        }
        auto task = StartGCAction(Counters->GetConsumerCounter("GC")->GetRemoveGCCounters());
        if (!task) {
            return false;
        }
        CurrentGCAction = task;
        return true;
    }
};

}
