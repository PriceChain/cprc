package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgJoinRegistryMember = "join_registry_member"

var _ sdk.Msg = &MsgJoinRegistryMember{}

func NewMsgJoinRegistryMember(creator string, registryId string, stakeAmount string) *MsgJoinRegistryMember {
	return &MsgJoinRegistryMember{
		Creator:     creator,
		RegistryId:  registryId,
		StakeAmount: stakeAmount,
	}
}

func (msg *MsgJoinRegistryMember) Route() string {
	return RouterKey
}

func (msg *MsgJoinRegistryMember) Type() string {
	return TypeMsgJoinRegistryMember
}

func (msg *MsgJoinRegistryMember) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgJoinRegistryMember) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgJoinRegistryMember) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
