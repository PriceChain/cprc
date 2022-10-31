package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgJoinRegistryCoOperator = "join_registry_co_operator"

var _ sdk.Msg = &MsgJoinRegistryCoOperator{}

func NewMsgJoinRegistryCoOperator(creator string, registryId string, stakeAmount string) *MsgJoinRegistryCoOperator {
	return &MsgJoinRegistryCoOperator{
		Creator:     creator,
		RegistryId:  registryId,
		StakeAmount: stakeAmount,
	}
}

func (msg *MsgJoinRegistryCoOperator) Route() string {
	return RouterKey
}

func (msg *MsgJoinRegistryCoOperator) Type() string {
	return TypeMsgJoinRegistryCoOperator
}

func (msg *MsgJoinRegistryCoOperator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgJoinRegistryCoOperator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgJoinRegistryCoOperator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
