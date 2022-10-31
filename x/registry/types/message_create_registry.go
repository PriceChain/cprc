package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateRegistry = "create_registry"

var _ sdk.Msg = &MsgCreateRegistry{}

func NewMsgCreateRegistry(creator string, name string, stakeAmount string, quorum string, consensusExpiringTime string) *MsgCreateRegistry {
	return &MsgCreateRegistry{
		Creator:               creator,
		Name:                  name,
		StakeAmount:           stakeAmount,
		Quorum:                quorum,
		ConsensusExpiringTime: consensusExpiringTime,
	}
}

func (msg *MsgCreateRegistry) Route() string {
	return RouterKey
}

func (msg *MsgCreateRegistry) Type() string {
	return TypeMsgCreateRegistry
}

func (msg *MsgCreateRegistry) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRegistry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRegistry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
