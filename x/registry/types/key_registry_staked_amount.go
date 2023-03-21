package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RegistryStakedAmountKeyPrefix is the prefix to retrieve all RegistryStakedAmount
	RegistryStakedAmountKeyPrefix = "RegistryStakedAmount/value/"
)

// RegistryStakedAmountKey returns the store key to retrieve a RegistryStakedAmount from the index fields
func RegistryStakedAmountKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
