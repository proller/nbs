GO_LIBRARY()

SUBSCRIBER(g:go-contrib)

LICENSE(Apache-2.0)

SRCS(
    baggage.go
    context.go
)

GO_TEST_SRCS(context_test.go)

END()

RECURSE(
    gotest
)
