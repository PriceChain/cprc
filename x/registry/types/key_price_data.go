package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PriceDataKeyPrefix is the prefix to retrieve all PriceData
	PriceDataKeyPrefix = "PriceData/value/"
)

// PriceDataKey returns the store key to retrieve a PriceData from the index fields
func PriceDataKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
