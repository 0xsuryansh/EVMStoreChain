package evmstorechain

import (
	"EVMStoreChain/x/evmstorechain/keeper"
	"EVMStoreChain/x/evmstorechain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the vote
	for _, elem := range genState.VoteList {
		k.SetVote(ctx, elem)
	}

	// Set vote count
	k.SetVoteCount(ctx, genState.VoteCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.VoteList = k.GetAllVote(ctx)
	genesis.VoteCount = k.GetVoteCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
