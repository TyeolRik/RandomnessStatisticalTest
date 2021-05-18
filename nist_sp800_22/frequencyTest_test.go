package nist_sp800_22

import (
	"crypto/rand"
	"fmt"
	"testing"
)

// generateRandomBitArray() is using Package rand, which implements a cryptographically secure random number generator.
func generateRandomBitArray() []uint8 {
	b := make([]byte, 8)
	n, err := rand.Read(b)
	if n != 8 {
		panic(n)
	} else if err != nil {
		panic(err)
	}

	ret := make([]uint8, 64)
	retIndex := 0
	for _, value := range b {
		convert_uint8_to_bitString := fmt.Sprintf("%08b", value)
		for i := 0; i < 8; i++ {
			ret[retIndex] = convert_uint8_to_bitString[i] - 48
			retIndex = retIndex + 1
		}
	}

	return ret
}

func TestFrequencyTest(t *testing.T) {
	for i := 0; i < 100; i++ {
		randomBitsArray := generateRandomBitArray()
		P_value := FrequencyTest(randomBitsArray)
		fmt.Println(P_value)
		if P_value < 0.01 {
			t.Errorf("Non-random. Value was %f", P_value)
		}
	}
}
