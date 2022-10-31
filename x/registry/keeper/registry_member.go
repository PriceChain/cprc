package keeper

import (
	"encoding/binary"

	"github.com/PriceChain/rd_net/x/registry/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetRegistryMemberCount get the total number of registryMember
func (k Keeper) GetRegistryMemberCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RegistryMemberCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetRegistryMemberCount set the total number of registryMember
func (k Keeper) SetRegistryMemberCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RegistryMemberCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendRegistryMember appends a registryMember in the store with a new id and update the count
func (k Keeper) AppendRegistryMember(
	ctx sdk.Context,
	registryMember types.RegistryMember,
) uint64 {
	// Create the registryMember
	count := k.GetRegistryMemberCount(ctx)

	// Set the ID of the appended value
	registryMember.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryMemberKey))
	appendedValue := k.cdc.MustMarshal(&registryMember)
	store.Set(GetRegistryMemberIDBytes(registryMember.Id), appendedValue)

	// Update registryMember count
	k.SetRegistryMemberCount(ctx, count+1)

	return count
}

// SetRegistryMember set a specific registryMember in the store
func (k Keeper) SetRegistryMember(ctx sdk.Context, registryMember types.RegistryMember) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryMemberKey))
	b := k.cdc.MustMarshal(&registryMember)
	store.Set(GetRegistryMemberIDBytes(registryMember.Id), b)
}

// GetRegistryMember returns a registryMember from its id
func (k Keeper) GetRegistryMember(ctx sdk.Context, id uint64) (val types.RegistryMember, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryMemberKey))
	b := store.Get(GetRegistryMemberIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRegistryMember removes a registryMember from the store
func (k Keeper) RemoveRegistryMember(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryMemberKey))
	store.Delete(GetRegistryMemberIDBytes(id))
}

// GetAllRegistryMember returns all registryMember
func (k Keeper) GetAllRegistryMember(ctx sdk.Context) (list []types.RegistryMember) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistryMemberKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RegistryMember
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetRegistryMemberIDBytes returns the byte representation of the ID
func GetRegistryMemberIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetRegistryMemberIDFromBytes returns ID in uint64 format from a byte array
func GetRegistryMemberIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
