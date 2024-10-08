#pragma once
#include <contrib/ydb/library/yql/minikql/computation/mkql_computation_node.h>
#include <contrib/ydb/library/yql/ast/yql_expr_types.h>

namespace NKikimr {
namespace NMiniKQL {

template <NYql::ETypeAnnotationKind Kind>
IComputationNode* WrapMakeType(TCallable& callable, const TComputationNodeFactoryContext& ctx, ui32 exprCtxMutableIndex);

}
}
