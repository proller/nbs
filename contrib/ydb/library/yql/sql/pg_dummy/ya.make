LIBRARY()

PROVIDES(
    yql_pg_sql_translator
    yql_pg_runtime
)

PEERDIR(
    contrib/ydb/library/yql/parser/pg_wrapper/interface
    contrib/ydb/library/yql/minikql
)

SRCS(
    pg_sql_dummy.cpp
)

YQL_LAST_ABI_VERSION()

END()
