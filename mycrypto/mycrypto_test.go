package mycrypto

import (
	"testing"
)

func TestKeccak256(t *testing.T) {
	testSHA3 := keccak256("FFFF")
	var correct = [32]byte{224, 134, 154, 112, 160, 228, 47, 39, 26, 158, 90, 128, 41, 197, 180, 81, 206, 184, 245, 109, 204, 188, 128, 30, 66, 40, 99, 77, 133, 116, 163, 68}
	for idx, correctValue := range correct {
		if testSHA3[idx] != correctValue {
			t.Errorf("Wrong")
		}
	}
}
