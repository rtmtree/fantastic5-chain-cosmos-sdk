package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoredMWKeyPrefix is the prefix to retrieve all StoredMW
	StoredMWKeyPrefix = "StoredMW/value/"
)

// StoredMWKey returns the store key to retrieve a StoredMW from the index fields
func StoredMWKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
