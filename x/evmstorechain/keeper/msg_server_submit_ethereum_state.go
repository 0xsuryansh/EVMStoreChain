package keeper

import (
	"context"

	"EVMStoreChain/x/evmstorechain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitEthereumState(goCtx context.Context, msg *types.MsgSubmitEthereumState) (*types.MsgSubmitEthereumStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitEthereumStateResponse{}, nil
}
