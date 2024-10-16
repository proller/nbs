LIBRARY()

SRCS(
    GLOBAL stock_workload.cpp
    GLOBAL kv_workload.cpp
    workload_factory.cpp
)

PEERDIR(
    library/cpp/getopt
    contrib/ydb/public/api/protos
    contrib/ydb/public/sdk/cpp/client/ydb_table
)

GENERATE_ENUM_SERIALIZATION_WITH_HEADER(kv_workload.h)
GENERATE_ENUM_SERIALIZATION_WITH_HEADER(stock_workload.h)

END()
