UNITTEST_FOR(cloud/blockstore/libs/client)

INCLUDE(${ARCADIA_ROOT}/cloud/storage/core/tests/recipes/small.inc)

SRCS(
    client_ut.cpp
    durable_ut.cpp
    session_ut.cpp
)

PEERDIR(
)

END()
