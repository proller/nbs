GO_TEST_FOR(cloud/disk_manager/internal/pkg/facade)

SET_APPEND(RECIPE_ARGS --creation-and-deletion-allowed-only-for-disks-with-id-prefix "Test")
INCLUDE(${ARCADIA_ROOT}/cloud/disk_manager/internal/pkg/facade/testcommon/common.inc)

GO_XTEST_SRCS(
    snapshot_service_test.go
)

END()
