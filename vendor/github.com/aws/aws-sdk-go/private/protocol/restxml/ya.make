GO_LIBRARY()

LICENSE(Apache-2.0)

SRCS(restxml.go)

GO_XTEST_SRCS(
    build_test.go
    unmarshal_test.go
)

END()

RECURSE(gotest)
