#pragma once
#include <contrib/ydb/library/yql/minikql/computation/mkql_computation_node.h>

namespace NKikimr {
namespace NMiniKQL {

IComputationNode* WrapTypeHandle(TCallable& callable, const TComputationNodeFactoryContext& ctx, ui32 exprCtxMutableIndex);

}
}
