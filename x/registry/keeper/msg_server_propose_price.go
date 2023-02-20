package keeper

import (
	"context"

	"github.com/PriceChain/cprc/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ProposePrice(goCtx context.Context, msg *types.MsgProposePrice) (*types.MsgProposePriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_ = ctx
	// Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// RegistryId   string `protobuf:"bytes,2,opt,name=registryId,proto3" json:"registryId,omitempty"`
	// StoreName    string `protobuf:"bytes,3,opt,name=storeName,proto3" json:"storeName,omitempty"`
	// StoreAddr    string `protobuf:"bytes,4,opt,name=storeAddr,proto3" json:"storeAddr,omitempty"`
	// PurchaseTime string `protobuf:"bytes,5,opt,name=purchaseTime,proto3" json:"purchaseTime,omitempty"`
	// ProdName     string `protobuf:"bytes,6,opt,name=prodName,proto3" json:"prodName,omitempty"`
	// Price        string `protobuf:"bytes,7,opt,name=price,proto3" json:"price,omitempty"`
	// ReceiptCode  string `protobuf:"bytes,8,opt,name=receiptCode,proto3" json:"receiptCode,omitempty"`
	// Reserved     string `protobuf:"bytes,9,opt,name=reserved,proto3" json:"reserved,omitempty"`

	return &types.MsgProposePriceResponse{}, nil
}
