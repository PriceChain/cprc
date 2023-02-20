package keeper

import (
	"github.com/PriceChain/cprc/x/registry/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPriceData set a specific priceData in the store from its index
func (k Keeper) SetPriceData(ctx sdk.Context, priceData types.PriceData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceDataKeyPrefix))
	b := k.cdc.MustMarshal(&priceData)
	store.Set(types.PriceDataKey(
		priceData.Index,
	), b)
}

// GetPriceData returns a priceData from its index
func (k Keeper) GetPriceData(
	ctx sdk.Context,
	index string,

) (val types.PriceData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceDataKeyPrefix))

	b := store.Get(types.PriceDataKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePriceData removes a priceData from the store
func (k Keeper) RemovePriceData(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceDataKeyPrefix))
	store.Delete(types.PriceDataKey(
		index,
	))
}

// GetAllPriceData returns all priceData
func (k Keeper) GetAllPriceData(ctx sdk.Context) (list []types.PriceData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceDataKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PriceData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
