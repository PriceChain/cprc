package keeper

import (
	"encoding/binary"

	"github.com/PriceChain/cprc/x/prcibc/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetIbcMsgCount get the total number of ibcMsg
func (k Keeper) GetIbcMsgCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.IbcMsgCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetIbcMsgCount set the total number of ibcMsg
func (k Keeper) SetIbcMsgCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.IbcMsgCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendIbcMsg appends a ibcMsg in the store with a new id and update the count
func (k Keeper) AppendIbcMsg(
	ctx sdk.Context,
	ibcMsg types.IbcMsg,
) uint64 {
	// Create the ibcMsg
	count := k.GetIbcMsgCount(ctx)

	// Set the ID of the appended value
	ibcMsg.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IbcMsgKey))
	appendedValue := k.cdc.MustMarshal(&ibcMsg)
	store.Set(GetIbcMsgIDBytes(ibcMsg.Id), appendedValue)

	// Update ibcMsg count
	k.SetIbcMsgCount(ctx, count+1)

	return count
}

// SetIbcMsg set a specific ibcMsg in the store
func (k Keeper) SetIbcMsg(ctx sdk.Context, ibcMsg types.IbcMsg) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IbcMsgKey))
	b := k.cdc.MustMarshal(&ibcMsg)
	store.Set(GetIbcMsgIDBytes(ibcMsg.Id), b)
}

// GetIbcMsg returns a ibcMsg from its id
func (k Keeper) GetIbcMsg(ctx sdk.Context, id uint64) (val types.IbcMsg, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IbcMsgKey))
	b := store.Get(GetIbcMsgIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveIbcMsg removes a ibcMsg from the store
func (k Keeper) RemoveIbcMsg(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IbcMsgKey))
	store.Delete(GetIbcMsgIDBytes(id))
}

// GetAllIbcMsg returns all ibcMsg
func (k Keeper) GetAllIbcMsg(ctx sdk.Context) (list []types.IbcMsg) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IbcMsgKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.IbcMsg
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetIbcMsgIDBytes returns the byte representation of the ID
func GetIbcMsgIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetIbcMsgIDFromBytes returns ID in uint64 format from a byte array
func GetIbcMsgIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
