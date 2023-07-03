package keeper_test

import (
	"strconv"
	"testing"

	keepertest "EVMStoreChain/testutil/keeper"
	"EVMStoreChain/testutil/nullify"
	"EVMStoreChain/x/evmstorechain/keeper"
	"EVMStoreChain/x/evmstorechain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNBlockstoragestate(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Blockstoragestate {
	items := make([]types.Blockstoragestate, n)
	for i := range items {
		items[i].Blocknumber = strconv.Itoa(i)

		keeper.SetBlockstoragestate(ctx, items[i])
	}
	return items
}

func TestBlockstoragestateGet(t *testing.T) {
	keeper, ctx := keepertest.EvmstorechainKeeper(t)
	items := createNBlockstoragestate(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetBlockstoragestate(ctx,
			item.Blocknumber,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestBlockstoragestateRemove(t *testing.T) {
	keeper, ctx := keepertest.EvmstorechainKeeper(t)
	items := createNBlockstoragestate(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBlockstoragestate(ctx,
			item.Blocknumber,
		)
		_, found := keeper.GetBlockstoragestate(ctx,
			item.Blocknumber,
		)
		require.False(t, found)
	}
}

func TestBlockstoragestateGetAll(t *testing.T) {
	keeper, ctx := keepertest.EvmstorechainKeeper(t)
	items := createNBlockstoragestate(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBlockstoragestate(ctx)),
	)
}
