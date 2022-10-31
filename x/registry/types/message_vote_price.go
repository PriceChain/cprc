package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgVotePrice = "vote_price"

var _ sdk.Msg = &MsgVotePrice{}

func NewMsgVotePrice(creator string, registryId string, answer string, reserved string) *MsgVotePrice {
	return &MsgVotePrice{
		Creator:    creator,
		RegistryId: registryId,
		Answer:     answer,
		Reserved:   reserved,
	}
}

func (msg *MsgVotePrice) Route() string {
	return RouterKey
}

func (msg *MsgVotePrice) Type() string {
	return TypeMsgVotePrice
}

func (msg *MsgVotePrice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVotePrice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVotePrice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
