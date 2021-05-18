package mycrypto

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

func keccak256(_hex string) (digest [32]byte) {
	input, _ := hex.DecodeString(_hex) // Hex
	return sha3.Sum256(input)
}
