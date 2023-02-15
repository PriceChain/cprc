package keeper

import (
	"github.com/PriceChain/cprc/x/registry/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetRegistryStakedAmount set a specific registryStakedAmount in the store from its index
func (k Keeper) SetRegistryStakedAmount(ctx sdk.Context, registryStakedAmount types.RegistryStakedAmount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryStakedAmountKeyPrefix))
	b := k.cdc.MustMarshal(&registryStakedAmount)
	store.Set(types.RegistryStakedAmountKey(
		registryStakedAmount.Index,
	), b)
}

// GetRegistryStakedAmount returns a registryStakedAmount from its index
func (k Keeper) GetRegistryStakedAmount(
	ctx sdk.Context,
	index string,

) (val types.RegistryStakedAmount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryStakedAmountKeyPrefix))

	b := store.Get(types.RegistryStakedAmountKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRegistryStakedAmount removes a registryStakedAmount from the store
func (k Keeper) RemoveRegistryStakedAmount(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryStakedAmountKeyPrefix))
	store.Delete(types.RegistryStakedAmountKey(
		index,
	))
}

// GetAllRegistryStakedAmount returns all registryStakedAmount
func (k Keeper) GetAllRegistryStakedAmount(ctx sdk.Context) (list []types.RegistryStakedAmount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryStakedAmountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RegistryStakedAmount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
