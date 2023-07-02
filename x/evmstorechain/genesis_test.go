package evmstorechain_test

import (
	"testing"

	keepertest "EVMStoreChain/testutil/keeper"
	"EVMStoreChain/testutil/nullify"
	"EVMStoreChain/x/evmstorechain"
	"EVMStoreChain/x/evmstorechain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		VoteList: []types.Vote{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		VoteCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EvmstorechainKeeper(t)
	evmstorechain.InitGenesis(ctx, *k, genesisState)
	got := evmstorechain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.VoteList, got.VoteList)
	require.Equal(t, genesisState.VoteCount, got.VoteCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
