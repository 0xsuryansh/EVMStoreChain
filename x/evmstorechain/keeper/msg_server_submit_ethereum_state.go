package keeper

import (
	"context"

	"EVMStoreChain/x/evmstorechain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

)

func (k msgServer) SubmitEthereumState(goCtx context.Context, msg *types.MsgSubmitEthereumState) (*types.MsgSubmitEthereumStateResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)

    // Get the validator set
    validators := k.StakingKeeper.GetAllValidators(ctx)

    // Convert the creator to a ValAddress
    creator, err := sdk.ValAddressFromBech32(msg.Creator)
    if err != nil {
        return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
    }

    // Check if the message creator is a validator
    isValidator := false
    for _, validator := range validators {
        if validator.GetOperator().Equals(creator) {
            isValidator = true
            break
        }
    }

    // If the message creator is not a validator, return an error
    if !isValidator {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "only validators can vote")
    }

    // Handling the message
    vote := types.Vote{
        Validator:   msg.Creator,
        Blocknumber: msg.Blocknumber,
        State:       msg.State,
    }
    k.AppendVote(ctx, vote)

    return &types.MsgSubmitEthereumStateResponse{}, nil
}

