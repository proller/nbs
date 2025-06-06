RESOURCES_LIBRARY()
OWNER(
    g:yatest
    heretic
)

IF (TEST_TOOL_HOST_LOCAL)
    MESSAGE(WARNING Host test tool $TEST_TOOL_HOST_LOCAL will be used)
ENDIF()
IF (TEST_TOOL3_HOST_LOCAL)
    MESSAGE(WARNING Host test tool3 $TEST_TOOL3_HOST_LOCAL will be used)
ENDIF()

IF (OPENSOURCE)
    INCLUDE(host_os.ya.make.inc)
ELSE()
    INCLUDE(host.ya.make.inc)
ENDIF()

IF (TEST_TOOL_TARGET_LOCAL)
    MESSAGE(WARNING Target test tool $TEST_TOOL_TARGET_LOCAL will be used)
ENDIF()
IF (OS_IOS AND NOT BUILD_IOS_APP)
    DECLARE_EXTERNAL_RESOURCE(TEST_TOOL_TARGET sbr:707351393)
    INCLUDE(${ARCADIA_ROOT}/build/platform/xcode/tools/ya.make.inc)
ENDIF()

END()
