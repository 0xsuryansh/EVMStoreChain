package keeper

import (
	"EVMStoreChain/x/evmstorechain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetBlockstoragestate set a specific blockstoragestate in the store from its index
func (k Keeper) SetBlockstoragestate(ctx sdk.Context, blockstoragestate types.Blockstoragestate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockstoragestateKeyPrefix))
	b := k.cdc.MustMarshal(&blockstoragestate)
	store.Set(types.BlockstoragestateKey(
		blockstoragestate.Blocknumber,
	), b)
}

// GetBlockstoragestate returns a blockstoragestate from its index
func (k Keeper) GetBlockstoragestate(
	ctx sdk.Context,
	blocknumber string,

) (val types.Blockstoragestate, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockstoragestateKeyPrefix))

	b := store.Get(types.BlockstoragestateKey(
		blocknumber,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBlockstoragestate removes a blockstoragestate from the store
func (k Keeper) RemoveBlockstoragestate(
	ctx sdk.Context,
	blocknumber string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockstoragestateKeyPrefix))
	store.Delete(types.BlockstoragestateKey(
		blocknumber,
	))
}

// GetAllBlockstoragestate returns all blockstoragestate
func (k Keeper) GetAllBlockstoragestate(ctx sdk.Context) (list []types.Blockstoragestate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockstoragestateKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Blockstoragestate
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
