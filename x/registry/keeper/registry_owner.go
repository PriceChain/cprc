package keeper

import (
	"encoding/binary"

	"github.com/PriceChain/cprc/x/registry/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetRegistryOwnerCount get the total number of registryOwner
func (k Keeper) GetRegistryOwnerCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RegistryOwnerCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetRegistryOwnerCount set the total number of registryOwner
func (k Keeper) SetRegistryOwnerCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RegistryOwnerCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendRegistryOwner appends a registryOwner in the store with a new id and update the count
func (k Keeper) AppendRegistryOwner(
	ctx sdk.Context,
	registryOwner types.RegistryOwner,
) uint64 {
	// Create the registryOwner
	count := k.GetRegistryOwnerCount(ctx)

	// Set the ID of the appended value
	registryOwner.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryOwnerKey))
	appendedValue := k.cdc.MustMarshal(&registryOwner)
	store.Set(GetRegistryOwnerIDBytes(registryOwner.Id), appendedValue)

	// Update registryOwner count
	k.SetRegistryOwnerCount(ctx, count+1)

	return count
}

// SetRegistryOwner set a specific registryOwner in the store
func (k Keeper) SetRegistryOwner(ctx sdk.Context, registryOwner types.RegistryOwner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryOwnerKey))
	b := k.cdc.MustMarshal(&registryOwner)
	store.Set(GetRegistryOwnerIDBytes(registryOwner.Id), b)
}

// GetRegistryOwner returns a registryOwner from its id
func (k Keeper) GetRegistryOwner(ctx sdk.Context, id uint64) (val types.RegistryOwner, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryOwnerKey))
	b := store.Get(GetRegistryOwnerIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRegistryOwner removes a registryOwner from the store
func (k Keeper) RemoveRegistryOwner(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryOwnerKey))
	store.Delete(GetRegistryOwnerIDBytes(id))
}

// GetAllRegistryOwner returns all registryOwner
func (k Keeper) GetAllRegistryOwner(ctx sdk.Context) (list []types.RegistryOwner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryOwnerKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RegistryOwner
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetRegistryOwnerIDBytes returns the byte representation of the ID
func GetRegistryOwnerIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetRegistryOwnerIDFromBytes returns ID in uint64 format from a byte array
func GetRegistryOwnerIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
