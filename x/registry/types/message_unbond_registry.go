package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUnbondRegistry = "unbond_registry"

var _ sdk.Msg = &MsgUnbondRegistry{}

func NewMsgUnbondRegistry(creator string, registryId string, unbondAmount string, reason string) *MsgUnbondRegistry {
	return &MsgUnbondRegistry{
		Creator:      creator,
		RegistryId:   registryId,
		UnbondAmount: unbondAmount,
		Reason:       reason,
	}
}

func (msg *MsgUnbondRegistry) Route() string {
	return RouterKey
}

func (msg *MsgUnbondRegistry) Type() string {
	return TypeMsgUnbondRegistry
}

func (msg *MsgUnbondRegistry) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnbondRegistry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnbondRegistry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
