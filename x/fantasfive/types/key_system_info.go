package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SystemInfoKeyPrefix is the prefix to retrieve all SystemInfo
	SystemInfoKeyPrefix = "SystemInfo/value/"
)

// SystemInfoKey returns the store key to retrieve a SystemInfo from the index fields
func SystemInfoKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
