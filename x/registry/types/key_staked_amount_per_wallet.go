package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StakedAmountPerWalletKeyPrefix is the prefix to retrieve all StakedAmountPerWallet
	StakedAmountPerWalletKeyPrefix = "StakedAmountPerWallet/value/"
)

// StakedAmountPerWalletKey returns the store key to retrieve a StakedAmountPerWallet from the index fields
func StakedAmountPerWalletKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
