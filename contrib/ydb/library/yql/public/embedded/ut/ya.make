UNITTEST()

SIZE(MEDIUM)

TIMEOUT(300)

PEERDIR(
    contrib/ydb/library/yql/public/udf/service/exception_policy
    contrib/ydb/library/yql/public/embedded
)

YQL_LAST_ABI_VERSION()

SRCS(
    yql_embedded_ut.cpp
)

END()

