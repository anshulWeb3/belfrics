package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// Kyc3KeyPrefix is the prefix to retrieve all Kyc3
	Kyc3KeyPrefix = "Kyc3/value/"
)

// Kyc3Key returns the store key to retrieve a Kyc3 from the index fields
func Kyc3Key(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
