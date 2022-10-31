package keeper

import (
	"encoding/binary"

	"github.com/PriceChain/rd_net/x/registry/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetRegistryCount get the total number of registry
func (k Keeper) GetRegistryCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RegistryCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetRegistryCount set the total number of registry
func (k Keeper) SetRegistryCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RegistryCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendRegistry appends a registry in the store with a new id and update the count
func (k Keeper) AppendRegistry(
	ctx sdk.Context,
	registry types.Registry,
) uint64 {
	// Create the registry
	count := k.GetRegistryCount(ctx)

	// Set the ID of the appended value
	registry.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryKey))
	appendedValue := k.cdc.MustMarshal(&registry)
	store.Set(GetRegistryIDBytes(registry.Id), appendedValue)

	// Update registry count
	k.SetRegistryCount(ctx, count+1)

	return count
}

// SetRegistry set a specific registry in the store
func (k Keeper) SetRegistry(ctx sdk.Context, registry types.Registry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryKey))
	b := k.cdc.MustMarshal(&registry)
	store.Set(GetRegistryIDBytes(registry.Id), b)
}

// GetRegistry returns a registry from its id
func (k Keeper) GetRegistry(ctx sdk.Context, id uint64) (val types.Registry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryKey))
	b := store.Get(GetRegistryIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRegistry removes a registry from the store
func (k Keeper) RemoveRegistry(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryKey))
	store.Delete(GetRegistryIDBytes(id))
}

// GetAllRegistry returns all registry
func (k Keeper) GetAllRegistry(ctx sdk.Context) (list []types.Registry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Registry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetRegistryIDBytes returns the byte representation of the ID
func GetRegistryIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetRegistryIDFromBytes returns ID in uint64 format from a byte array
func GetRegistryIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
