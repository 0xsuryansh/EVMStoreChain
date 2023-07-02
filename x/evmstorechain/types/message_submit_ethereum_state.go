package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitEthereumState = "submit_ethereum_state"

var _ sdk.Msg = &MsgSubmitEthereumState{}

func NewMsgSubmitEthereumState(creator string, blocknumber uint64, state uint64) *MsgSubmitEthereumState {
	return &MsgSubmitEthereumState{
		Creator:     creator,
		Blocknumber: blocknumber,
		State:       state,
	}
}

func (msg *MsgSubmitEthereumState) Route() string {
	return RouterKey
}

func (msg *MsgSubmitEthereumState) Type() string {
	return TypeMsgSubmitEthereumState
}

func (msg *MsgSubmitEthereumState) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitEthereumState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitEthereumState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
