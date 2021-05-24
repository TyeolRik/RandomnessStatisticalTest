package nist_sp800_22

import (
	"crypto/rand"
	"fmt"
	"testing"
)

// generateRandomBitArray() is using Package rand, which implements a cryptographically secure random number generator.
func generateRandomBitArray() []uint8 {
	b := make([]byte, 100)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	ret := make([]uint8, len(b)*8)
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

func TestFrequency(t *testing.T) {
	var pass uint64 = 0
	var P_value float64
	failed := make([][]uint8, 0)

	for i := 0; i < 100; i++ {
		randomBitsArray := generateRandomBitArray()
		P_value = Frequency(randomBitsArray)
		// fmt.Println(P_value)
		if P_value >= 0.01 {
			pass++
		} else {
			failed = append(failed, randomBitsArray)
		}
	}
	if pass < 95 {
		t.Errorf("Non-random. Value was %f", P_value)
	} else {
		fmt.Printf("Failed test was total %d\n", len(failed))
		for _, value := range failed {
			fmt.Println(Frequency(value))
			fmt.Println(value)
		}
	}
}

func TestBlockFrequency(t *testing.T) {
	epsilon = []uint8{0, 1, 1, 0, 0, 1, 1, 0, 1, 0}

	P_value := blockFrequency(3, 10)
	fmt.Println(P_value)
}
