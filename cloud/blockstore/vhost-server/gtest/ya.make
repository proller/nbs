GTEST()

INCLUDE(${ARCADIA_ROOT}/cloud/storage/core/tests/recipes/medium.inc)

SRCS(
    ../backend.cpp
    ../backend_aio.cpp
    ../critical_event.cpp
    ../histogram.cpp
    ../options.cpp
    ../request_aio.cpp
    ../server.cpp
    ../stats.cpp

    ../request_aio_ut.cpp
    ../server_ut.cpp
)

ADDINCL(
    cloud/contrib/vhost
)

PEERDIR(
    cloud/blockstore/libs/common
    cloud/blockstore/libs/encryption
    cloud/blockstore/libs/encryption/model
    cloud/contrib/vhost

    cloud/storage/core/libs/common
    cloud/storage/core/libs/diagnostics
    cloud/storage/core/libs/vhost-client

    library/cpp/getopt
    library/cpp/getopt/small

    contrib/libs/libaio
)

END()
