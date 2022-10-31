package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgModifyRegistry = "modify_registry"

var _ sdk.Msg = &MsgModifyRegistry{}

func NewMsgModifyRegistry(creator string, registryId string, stakeAmount string, name string, quorum string, consensusExpringTime string, reason string) *MsgModifyRegistry {
	return &MsgModifyRegistry{
		Creator:              creator,
		RegistryId:           registryId,
		StakeAmount:          stakeAmount,
		Name:                 name,
		Quorum:               quorum,
		ConsensusExpringTime: consensusExpringTime,
		Reason:               reason,
	}
}

func (msg *MsgModifyRegistry) Route() string {
	return RouterKey
}

func (msg *MsgModifyRegistry) Type() string {
	return TypeMsgModifyRegistry
}

func (msg *MsgModifyRegistry) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgModifyRegistry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgModifyRegistry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
