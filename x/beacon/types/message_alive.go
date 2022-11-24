package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAlive = "alive"

var _ sdk.Msg = &MsgAlive{}

func NewMsgAlive(creator string) *MsgAlive {
	return &MsgAlive{
		Creator: creator,
	}
}

func (msg *MsgAlive) Route() string {
	return RouterKey
}

func (msg *MsgAlive) Type() string {
	return TypeMsgAlive
}

func (msg *MsgAlive) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAlive) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAlive) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
