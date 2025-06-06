LIBRARY()

OWNER(g:yql)

SRCS(
    sql_format.cpp
)

RESOURCE(DONT_PARSE ../SQLv1.g.in SQLv1.g.in)

PEERDIR(
    contrib/ydb/library/yql/parser/lexer_common
    contrib/ydb/library/yql/sql/settings
    contrib/ydb/library/yql/sql/v1/lexer
    contrib/ydb/library/yql/sql/v1/proto_parser
    contrib/ydb/library/yql/core/sql_types
    library/cpp/protobuf/util
    library/cpp/resource
)

END()

RECURSE_FOR_TESTS(
    ut
)
