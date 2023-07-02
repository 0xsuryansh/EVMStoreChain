package keeper

import (
	"EVMStoreChain/x/evmstorechain/types"
)

var _ types.QueryServer = Keeper{}
