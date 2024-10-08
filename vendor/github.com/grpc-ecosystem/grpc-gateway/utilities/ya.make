GO_LIBRARY()

LICENSE(BSD-3-Clause)

SRCS(
    doc.go
    pattern.go
    readerfactory.go
    trie.go
)

GO_XTEST_SRCS(trie_test.go)

END()

RECURSE(gotest)
