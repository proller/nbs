UNITTEST_FOR(cloud/blockstore/libs/storage/disk_registry)

INCLUDE(${ARCADIA_ROOT}/cloud/storage/core/tests/recipes/medium.inc)

SRCS(
    disk_registry_ut_lifecycle.cpp
)

PEERDIR(
    cloud/blockstore/config
    cloud/blockstore/libs/storage/api
    cloud/blockstore/libs/storage/disk_registry/testlib
    cloud/blockstore/libs/storage/testlib
    library/cpp/testing/unittest
    contrib/ydb/core/testlib/basics
)

END()
