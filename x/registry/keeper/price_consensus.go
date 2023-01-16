package keeper

import (
	"encoding/binary"

	"github.com/PriceChain/cprc/x/registry/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetPriceConsensusCount get the total number of priceConsensus
func (k Keeper) GetPriceConsensusCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PriceConsensusCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPriceConsensusCount set the total number of priceConsensus
func (k Keeper) SetPriceConsensusCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PriceConsensusCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPriceConsensus appends a priceConsensus in the store with a new id and update the count
func (k Keeper) AppendPriceConsensus(
	ctx sdk.Context,
	priceConsensus types.PriceConsensus,
) uint64 {
	// Create the priceConsensus
	count := k.GetPriceConsensusCount(ctx)

	// Set the ID of the appended value
	priceConsensus.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceConsensusKey))
	appendedValue := k.cdc.MustMarshal(&priceConsensus)
	store.Set(GetPriceConsensusIDBytes(priceConsensus.Id), appendedValue)

	// Update priceConsensus count
	k.SetPriceConsensusCount(ctx, count+1)

	return count
}

// SetPriceConsensus set a specific priceConsensus in the store
func (k Keeper) SetPriceConsensus(ctx sdk.Context, priceConsensus types.PriceConsensus) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceConsensusKey))
	b := k.cdc.MustMarshal(&priceConsensus)
	store.Set(GetPriceConsensusIDBytes(priceConsensus.Id), b)
}

// GetPriceConsensus returns a priceConsensus from its id
func (k Keeper) GetPriceConsensus(ctx sdk.Context, id uint64) (val types.PriceConsensus, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceConsensusKey))
	b := store.Get(GetPriceConsensusIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePriceConsensus removes a priceConsensus from the store
func (k Keeper) RemovePriceConsensus(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceConsensusKey))
	store.Delete(GetPriceConsensusIDBytes(id))
}

// GetAllPriceConsensus returns all priceConsensus
func (k Keeper) GetAllPriceConsensus(ctx sdk.Context) (list []types.PriceConsensus) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceConsensusKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PriceConsensus
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPriceConsensusIDBytes returns the byte representation of the ID
func GetPriceConsensusIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetPriceConsensusIDFromBytes returns ID in uint64 format from a byte array
func GetPriceConsensusIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
