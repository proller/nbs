PY3TEST()

INCLUDE(${ARCADIA_ROOT}/cloud/storage/core/tests/recipes/medium.inc)

TEST_SRCS(
    test.py
)

DEPENDS(
)

PEERDIR(
    cloud/blockstore/public/sdk/python/client
)

END()
