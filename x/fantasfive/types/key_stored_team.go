package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoredTeamKeyPrefix is the prefix to retrieve all StoredTeam
	StoredTeamKeyPrefix = "StoredTeam/value/"
)

// StoredTeamKey returns the store key to retrieve a StoredTeam from the index fields
func StoredTeamKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
