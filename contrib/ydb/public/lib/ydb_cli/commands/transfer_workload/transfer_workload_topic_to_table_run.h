#pragma once

#include <contrib/ydb/public/lib/ydb_cli/commands/ydb_workload.h>
#include <contrib/ydb/public/lib/ydb_cli/commands/topic_readwrite_scenario.h>

namespace NYdb::NConsoleClient {

class TCommandWorkloadTransferTopicToTableRun : public TWorkloadCommand {
public:
    TCommandWorkloadTransferTopicToTableRun();

    void Config(TConfig& config) override;
    void Parse(TConfig& config) override;
    int Run(TConfig& config) override;

private:
    TTopicReadWriteScenario Scenario;
};

}
