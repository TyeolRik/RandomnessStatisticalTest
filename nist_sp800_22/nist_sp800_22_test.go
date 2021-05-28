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

func TestConstant(t *testing.T) {
	err := prepare_CONSTANT_E_asEpsilon()
	if err != nil {
		t.Errorf("FAILED TO GET CONSTANT E")
	}
	// fmt.Println(epsilon[0:100])

	err = prepare_CONSTANT_PI_asEpsilon()
	if err != nil {
		t.Errorf("FAILED TO GET CONSTANT PI")
	}
	// fmt.Println(epsilon[0:100])
}

func TestFrequency(t *testing.T) {
	var passCount uint64 = 0
	var P_value float64
	var pass bool
	var err error
	failed := make([][]uint8, 0)

	for i := 0; i < 100; i++ {
		epsilon = generateRandomBitArray()
		P_value, pass, err = Frequency(uint64(len(epsilon)))
		// fmt.Println(P_value)
		if err != nil {
			t.Error(err)
		}
		if pass {
			passCount++
		} else {
			failed = append(failed, epsilon)
		}
	}
	if passCount < 95 {
		t.Errorf("Non-random. Value was %f", P_value)
	} else {
		fmt.Printf("Failed test was total %d\n", len(failed))
		for _, value := range failed {
			fmt.Println(Frequency(uint64(len(epsilon))))
			fmt.Println(value)
		}
	}
}

func TestBlockFrequency(t *testing.T) {
	testArray := []uint8{0, 1, 1, 0, 0, 1, 1, 0, 1, 0}
	inputEpsilon(testArray)

	P_value, pass, err := BlockFrequency(3, 10)
	fmt.Println(P_value, pass, err)
}

func TestRuns(t *testing.T) {
	inputEpsilonAsString("1100100100001111110110101010001000100001011010001100001000110100110001001100011001100010100010111000")

	P_value, _, _ := Runs(uint64(len(epsilon)))
	fmt.Println(P_value)
}

func TestLongestRunOfOnes(t *testing.T) {
	inputEpsilonAsString("11001100000101010110110001001100111000000000001001001101010100010001001111010110100000001101011111001100111001101101100010110010")

	P_value, _, _ := LongestRunOfOnes(128)
	fmt.Println(P_value)
}

func TestRank(t *testing.T) {
	readERR := prepare_CONSTANT_E_asEpsilon()
	if readERR != nil {
		t.Error("FAILED TO GET CONSTANT E")
	}
	epsilon = epsilon[0:100000]

	P_value, _, _ := Rank(100000)
	fmt.Printf("P-value : %f\n", P_value)
}

func TestNonOverlappingTemplateMatching(t *testing.T) {
	inputEpsilonAsString_NonRevert("10100100101110010110")

	theTemplate := []uint8{0, 0, 1}

	P_value, _, _ := NonOverlappingTemplateMatching(theTemplate, 10)
	fmt.Printf("P-value : %f\n", P_value)
}

func TestOverlappingTemplateMatching(t *testing.T) {
	/*
		readERR := prepare_CONSTANT_E_asEpsilon()
		if readERR != nil {
			t.Error("FAILED TO GET CONSTANT E")
		}
		epsilon = epsilon[0:1000000]
	*/
	inputEpsilonAsString_NonRevert("10111011110010110100011100101110111110000101101001")

	// theTemplate := []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1}
	theTemplate := []uint8{1, 1}

	P_value, _, _ := OverlappingTemplateMatching(theTemplate, 10)
	fmt.Printf("P-value : %f\n", P_value)
}
