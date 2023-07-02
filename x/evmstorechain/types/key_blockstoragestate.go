package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BlockstoragestateKeyPrefix is the prefix to retrieve all Blockstoragestate
	BlockstoragestateKeyPrefix = "Blockstoragestate/value/"
)

// BlockstoragestateKey returns the store key to retrieve a Blockstoragestate from the index fields
func BlockstoragestateKey(
	blocknumber string,
) []byte {
	var key []byte

	blocknumberBytes := []byte(blocknumber)
	key = append(key, blocknumberBytes...)
	key = append(key, []byte("/")...)

	return key
}
