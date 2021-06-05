package nist_sp800_22

import (
	"crypto/rand"
	"fmt"
	"reflect"
	"testing"

	"github.com/mjibson/go-dsp/fft"
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
	err := Prepare_CONSTANT_E_asEpsilon()
	if err != nil {
		t.Errorf("FAILED TO GET CONSTANT E")
	}
	// fmt.Println(epsilon[0:100])

	err = Prepare_CONSTANT_PI_asEpsilon()
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
	InputEpsilon(testArray)

	P_value, pass, err := BlockFrequency(3, 10)
	fmt.Println(P_value, pass, err)
}

func TestRuns(t *testing.T) {
	InputEpsilonAsString("1100100100001111110110101010001000100001011010001100001000110100110001001100011001100010100010111000")

	P_value, _, _ := Runs(uint64(len(epsilon)))
	fmt.Println(P_value)
}

func TestLongestRunOfOnes(t *testing.T) {
	InputEpsilonAsString("11001100000101010110110001001100111000000000001001001101010100010001001111010110100000001101011111001100111001101101100010110010")

	P_value, _, _ := LongestRunOfOnes(128)
	fmt.Println(P_value)
}

func TestRank(t *testing.T) {
	readERR := Prepare_CONSTANT_E_asEpsilon()
	if readERR != nil {
		t.Error("FAILED TO GET CONSTANT E")
	}
	epsilon = epsilon[0:100000]

	P_value, _, _ := Rank(100000)
	fmt.Printf("P-value : %f\n", P_value)
}

func TestNonOverlappingTemplateMatching(t *testing.T) {
	InputEpsilonAsString_NonRevert("10100100101110010110")

	theTemplate := []uint8{0, 0, 1}

	P_value, _, _ := NonOverlappingTemplateMatching(theTemplate, 10)
	fmt.Printf("P-value : %f\n", P_value)
}

func TestOverlappingTemplateMatching(t *testing.T) {
	readERR := Prepare_CONSTANT_E_asEpsilon()
	if readERR != nil {
		t.Error("FAILED TO GET CONSTANT E")
	}
	epsilon = epsilon[0:1000000]

	theTemplate := []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1}

	P_value, _, _ := OverlappingTemplateMatching(theTemplate, 1032)
	fmt.Printf("P-value : %f\n", P_value)
}

/*
func TestOverlappingTemplateMatchingExample(t *testing.T) {
	InputEpsilonAsString_NonRevert("10111011110010110100011100101110111110000101101001")
	P_value, _, _ := OverlappingTemplateMatching([]uint8{1, 1}, 10)
	fmt.Println("P_value", P_value)
}
*/
func TestUniversal(t *testing.T) {
	InputEpsilonAsString_NonRevert("01011010011101010111")
	P_value, _, _ := Universal(2, 4, uint64(len(epsilon)))
	fmt.Printf("P-value : %f\n", P_value)
}

func TestLinearComplexity(t *testing.T) {
	readERR := Prepare_CONSTANT_E_asEpsilon()
	if readERR != nil {
		t.Error("FAILED TO GET CONSTANT E")
	}
	epsilon = epsilon[0:1000000]
	P_value, _, _ := LinearComplexity(1000, uint64(len(epsilon)))
	fmt.Printf("P-value : %f\n", P_value)
}

func TestSerial(t *testing.T) {
	readERR := Prepare_CONSTANT_E_asEpsilon()
	if readERR != nil {
		t.Error("FAILED TO GET CONSTANT E")
	}
	epsilon = epsilon[0:1000000]

	P_values, _, _ := Serial(2, uint64(len(epsilon)))
	fmt.Printf("P-value : %f\n", P_values)
}

func TestApproximateEntropy(t *testing.T) {
	//inputEpsilonAsString_NonRevert("1100100100001111110110101010001000100001011010001100001000110100110001001100011001100010100010111000")
	//P_value, _, _ := ApproximateEntropy(2, uint64(len(epsilon)))
	InputEpsilonAsString_NonRevert("0100110101")
	P_value, _, _ := ApproximateEntropy(3, uint64(len(epsilon)))
	fmt.Printf("P-value : %f\n", P_value)
}
func TestCumulativeSums(t *testing.T) {
	InputEpsilonAsString_NonRevert("1100100100001111110110101010001000100001011010001100001000110100110001001100011001100010100010111000")
	P_value_forward, _, _ := CumulativeSums(0, uint64(len(epsilon)))
	P_value_backward, _, _ := CumulativeSums(1, uint64(len(epsilon)))
	fmt.Printf("P-value  (forward) : %f\n", P_value_forward)
	fmt.Printf("P-value (backward) : %f\n", P_value_backward)
}

func TestRandomExcursions(t *testing.T) {
	//inputEpsilonAsString_NonRevert("0110110101")

	readERR := Prepare_CONSTANT_E_asEpsilon()
	if readERR != nil {
		t.Error("FAILED TO GET CONSTANT E")
	}
	epsilon = epsilon[0:1000000]

	P_value, _, _ := RandomExcursions(uint64(len(epsilon)))
	fmt.Printf("P-value : %f\n", P_value)
}
func TestRandomExcursionsVariant(t *testing.T) {
	// inputEpsilonAsString_NonRevert("0110110101")

	readERR := Prepare_CONSTANT_E_asEpsilon()
	if readERR != nil {
		t.Error("FAILED TO GET CONSTANT E")
	}
	epsilon = epsilon[0:1000000]

	P_value, _, _ := RandomExcursionsVariant(uint64(len(epsilon)))
	fmt.Printf("P-value : %f\n", P_value)
}

func TestFunctions(t *testing.T) {
	a := []uint8{1, 2, 3, 4, 5}
	var b []uint8 = nil
	fmt.Println("isEqual", isEqualBetweenBitsArray(a, b))
	fmt.Println("isEqual", reflect.DeepEqual(a, b))
}

// Run all Benchmark
// go test -bench=^Benchmark . -run=^$ . -v -benchtime 5s -benchmem

func BenchmarkMyFunction(b *testing.B) {
	// go test -bench=^BenchmarkMyFunction . -run=^$ . -v -benchtime 5s -benchmem
	a1 := Uint_To_BitsArray(^uint64(0))
	a2 := Uint_To_BitsArray(^uint64(0))
	for i := 0; i < b.N; i++ {
		isEqualBetweenBitsArray(a1, a2)
	}
}

func BenchmarkReflect(b *testing.B) {
	// go test -bench=^BenchmarkReflect . -run=^$ . -v -benchtime 5s -benchmem
	a1 := Uint_To_BitsArray(^uint64(0))
	a2 := Uint_To_BitsArray(^uint64(0))
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(a1, a2)
	}
}

func BenchmarkDFT(b *testing.B) {
	// go test -bench=^BenchmarkDFT . -run=^$ . -v -benchtime 5s -benchmem
	Prepare_CONSTANT_E_asEpsilon()
	epsilon = epsilon[0:10000]
	// InputEpsilonAsString_NonRevert("1001010011")

	var float64_epsilon []float64 = make([]float64, len(epsilon))
	for i := range float64_epsilon {
		float64_epsilon[i] = 2*float64(epsilon[i]) - 1
	}
	for i := 0; i < b.N; i++ {
		DFT(float64_epsilon)
	}
}

func BenchmarkFFT(b *testing.B) {
	// go test -bench=^BenchmarkFFT . -run=^$ . -v -benchtime 5s -benchmem
	Prepare_CONSTANT_E_asEpsilon()
	epsilon = epsilon[0:1000000]
	//InputEpsilonAsString_NonRevert("1001010011")

	var float64_epsilon []float64 = make([]float64, len(epsilon))
	for i := range float64_epsilon {
		float64_epsilon[i] = 2*float64(epsilon[i]) - 1
	}
	for i := 0; i < b.N; i++ {
		fft.FFTReal(float64_epsilon)
	}
}
