LIBRARY()

YQL_ABI_VERSION(
    2
    28
    0
)

SRCS(
    static_udf.cpp
    topfreq.cpp
)

PEERDIR(
    contrib/ydb/library/yql/public/udf
)

END()
