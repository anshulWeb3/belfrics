package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// KycKeyPrefix is the prefix to retrieve all Kyc
	KycKeyPrefix = "Kyc/value/"
)

// KycKey returns the store key to retrieve a Kyc from the index fields
func KycKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
