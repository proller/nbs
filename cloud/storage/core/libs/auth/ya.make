LIBRARY()

GENERATE_ENUM_SERIALIZATION(auth_counters.h)

SRCS(
    auth_counters.cpp
    auth_scheme.cpp
    authorizer.cpp
)

PEERDIR(
    cloud/storage/core/libs/api
    cloud/storage/core/protos
    contrib/ydb/library/actors/core
    library/cpp/monlib/dynamic_counters
    contrib/ydb/core/security
)

END()

RECURSE_FOR_TESTS(
    ut
)
