package keeper

import (
	"github.com/PriceChain/cprc/x/registry/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetStakedAmountPerWallet set a specific stakedAmountPerWallet in the store from its index
func (k Keeper) SetRegistryStakedAmountPerWallet(ctx sdk.Context, stakedAmountPerWallet types.RegistryStakedAmountPerWallet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StakedAmountPerWalletKeyPrefix))
	b := k.cdc.MustMarshal(&stakedAmountPerWallet)
	store.Set(types.StakedAmountPerWalletKey(
		stakedAmountPerWallet.Index,
	), b)
}

// GetStakedAmountPerWallet returns a stakedAmountPerWallet from its index
func (k Keeper) GetRegistryStakedAmountPerWallet(
	ctx sdk.Context,
	index string,

) (val types.RegistryStakedAmountPerWallet, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StakedAmountPerWalletKeyPrefix))

	b := store.Get(types.StakedAmountPerWalletKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStakedAmountPerWallet removes a stakedAmountPerWallet from the store
func (k Keeper) RemoveRegistryStakedAmountPerWallet(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StakedAmountPerWalletKeyPrefix))
	store.Delete(types.StakedAmountPerWalletKey(
		index,
	))
}

// GetAllStakedAmountPerWallet returns all stakedAmountPerWallet
func (k Keeper) GetAllRegistryStakedAmountPerWallet(ctx sdk.Context) (list []types.RegistryStakedAmountPerWallet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StakedAmountPerWalletKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RegistryStakedAmountPerWallet
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
