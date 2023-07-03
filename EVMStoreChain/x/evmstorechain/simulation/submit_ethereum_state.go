package simulation

import (
	"math/rand"

	"EVMStoreChain/x/evmstorechain/keeper"
	"EVMStoreChain/x/evmstorechain/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgSubmitEthereumState(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSubmitEthereumState{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SubmitEthereumState simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SubmitEthereumState simulation not implemented"), nil, nil
	}
}
