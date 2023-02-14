package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgProposePrice = "propose_price"

var _ sdk.Msg = &MsgProposePrice{}

func NewMsgProposePrice(creator string, registryId string, storeName string, storeAddr string, purchaseTime string, prodName string, price string, receiptCode string, reserved string) *MsgProposePrice {
	return &MsgProposePrice{
		Creator:      creator,
		RegistryId:   registryId,
		StoreName:    storeName,
		StoreAddr:    storeAddr,
		PurchaseTime: purchaseTime,
		ProdName:     prodName,
		Price:        price,
		ReceiptCode:  receiptCode,
		Reserved:     reserved,
	}
}

func (msg *MsgProposePrice) Route() string {
	return RouterKey
}

func (msg *MsgProposePrice) Type() string {
	return TypeMsgProposePrice
}

func (msg *MsgProposePrice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgProposePrice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProposePrice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
