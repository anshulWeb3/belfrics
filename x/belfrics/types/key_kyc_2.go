package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// Kyc2KeyPrefix is the prefix to retrieve all Kyc2
	Kyc2KeyPrefix = "Kyc2/value/"
)

// Kyc2Key returns the store key to retrieve a Kyc2 from the index fields
func Kyc2Key(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
