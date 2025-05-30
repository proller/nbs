#pragma once

#include "command.h"

namespace NCloud::NBlockStore::NClient {

////////////////////////////////////////////////////////////////////////////////

TCommandPtr NewExecuteActionCommand(IBlockStorePtr client);

}   // namespace NCloud::NBlockStore::NClient
