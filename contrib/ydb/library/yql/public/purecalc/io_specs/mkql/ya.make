LIBRARY()

INCLUDE(ya.make.inc)

PEERDIR(
    contrib/ydb/library/yql/public/purecalc/common
)

END()

RECURSE(
    no_llvm
)

RECURSE_FOR_TESTS(
    ut
    ut/no_llvm
)

