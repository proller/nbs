UNITTEST_FOR(contrib/ydb/core/kqp/rm_service)

FORK_SUBTESTS()

IF (SANITIZER_TYPE OR WITH_VALGRIND)
    SIZE(MEDIUM)
ENDIF()

SRCS(
    kqp_resource_estimation_ut.cpp
    kqp_rm_ut.cpp
)

PEERDIR(
    contrib/ydb/core/kqp/ut/common
    contrib/ydb/library/yql/sql/pg_dummy
)

YQL_LAST_ABI_VERSION()

END()
