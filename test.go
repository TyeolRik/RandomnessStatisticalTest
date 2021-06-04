package main

import (
	. "github.com/tyeolrik/RandomnessStatisticalTest/nist_sp800_22"
)

func main() {
	//     Reference : NIST SP800-22 Revision 1a.
	// File Location : /reference/"NIST SP800-22 Revision 1a.pdf"

	// How to use?
	// Random bits should be in specific []uint8 variable, named epsilon, using SetEpsilon
	// There is Three types of method to input []uint8 (./nist_sp800_22/globalVariable.go)
	//   1. InputEpsilon(_input []uint8)
	//      - Just input Epsilon using _input, whose type is []uint8 slice.
	//   2. InputEpsilonAsString(_input string)
	//      - Easy Version. Input Epsilon using _input, whose type is String type.
	//      - In this function, parse _input String and make []uint8 slice.
	//      - Be aware that, this function revert array.
	//   3. InputEpsilonAsString_NonRevert(_input string)
	//      - Easy Version. Input Epsilon using _input, whose type is String type.
	//      - In this function, parse _input String and make []uint8 slice.
	//      - Be aware that, this function don't revert array.
	//   4. Prepare_CONSTANT_E_asEpsilon()
	//      - Source file : /assets/data.e
	//      - This function put Euler's number (= natural number = 2.718281828...) into Epsilon.
	//      - This E is composed of binary numbers, which is from NIST official code sts/data/data.e
	//      - In this function, get all 0 and 1 from file, and parse into []uint8 and put it into Epsilon
	//   5. Prepare_CONSTANT_PI_asEpsilon()
	//      - Source file : /assets/data.pi
	//      - Same as Prepare_CONSTANT_E_asEpsilon. The only difference is not E but PI.

	readERR := Prepare_CONSTANT_PI_asEpsilon()
	if readERR != nil {
		panic("FAILED to load natural PI")
	}
	SetEpsilon(GetEpsilon()[0:1000000]) // Put Your Data here.
	SetLevel(0.01)

	var P_values []float64
	var isRandoms []bool

	var P_value float64
	var isRandom bool
	var err error

	var n uint64 = uint64(len(GetEpsilon()))

	// Initialize Printer
	PrettyPrint_Init()

	// 2.1 Frequency Test (Page 24)
	P_value, isRandom, err = Frequency(n)
	if err != nil {
		panic(err)
	}
	PrettyPrint_Add("The Frequency (Monobit) Test", P_value, isRandom)

	// 2.2 Frequency Test within a Block (Page 26)
	P_value, isRandom, err = BlockFrequency(20, n)
	if err != nil {
		panic(err)
	}
	PrettyPrint_Add("Frequency Test within a Block", P_value, isRandom)

	// 2.3 The Runs Test (Page 27)
	P_value, isRandom, err = Runs(n)
	if err != nil {
		panic(err)
	}
	PrettyPrint_Add("The Runs Test", P_value, isRandom)

	// 2.4 Tests for the Longest-Run-of-Ones in a Block (Page 29)
	P_value, isRandom, err = LongestRunOfOnes(n)
	if err != nil {
		panic(err)
	}
	PrettyPrint_Add("Tests for the Longest-Run-of-Ones in a Block", P_value, isRandom)

	// 2.5 The Binary Matrix Rank Test (Page 32)
	P_value, isRandom, err = Rank(n)
	if err != nil {
		panic(err)
	}
	PrettyPrint_Add("The Binary Matrix Rank Test", P_value, isRandom)

	// 2.6 The Discrete Fourier Transform (Spectral) Test (Page 34)
	P_value, isRandom, err = DiscreteFourierTransform(n)
	if err != nil {
		panic(err)
	}
	PrettyPrint_Add("The Discrete Fourier Transform (Spectral) Test", P_value, isRandom)

	// 2.7 The Non-overlapping Template Matching Test (Page 36)
	var m uint64 = 10 // Block Size
	P_values = []float64{}
	isRandoms = []bool{}
	for i := 1; i < 500; i = i + 2 {
		P_value, isRandom, _ = NonOverlappingTemplateMatching(Uint_To_BitsArray_size_N(uint64(i), m), 10000)
		P_values = append(P_values, P_value)
		isRandoms = append(isRandoms, isRandom)
	}
	PrettyPrint_Add_Array("The Non-overlapping Template Matching Test", P_values, isRandoms)

	// 2.8 The Overlapping Template Matching Test (Page 39)
	// [Working now] There are lots of error in official document. I need time to cross check.
	//InputEpsilonAsString_NonRevert("10111011110010110100011100101110111110000101101001")
	//P_value, isRandom, err = OverlappingTemplateMatching([]uint8{1, 1}, 10)
	//P_value, isRandom, err = OverlappingTemplateMatching([]uint8{1, 1, 1, 1, 1, 1, 1, 1, 1}, 1032)
	// P_value, isRandom, err = OverlappingTemplateMatching([]uint8{0, 1, 0, 1, 0, 1, 0, 1, 0, 1}, 1000)
	//if err != nil {
	//	panic(err)
	//}
	//PrettyPrint_Add("The Discrete Fourier Transform (Spectral) Test", P_value, isRandom)

	// 2.9 Maurer's "Universal Statistical" Test
	P_value, isRandom, err = Universal_Recommended()
	if err != nil {
		panic(err)
	}
	PrettyPrint_Add("Maurer's \"Universal Statistical\" Test", P_value, isRandom)

	// 2.10 Linear Complexity Test
	P_value, isRandom, err = LinearComplexity(1000, n)
	if err != nil {
		panic(err)
	}
	PrettyPrint_Add("Linear Complexity Test", P_value, isRandom)

	// 2.11 Serial Test
	// Be aware that, if m > 8, Too SLOW. Because Time complexity is over O(2^n)
	P_values, isRandoms, err = Serial(2, n)
	if err != nil {
		panic(err)
	}
	PrettyPrint_Add_Array("Serial Test", P_values, isRandoms)

	// 2.12 Approximate Entropy Test
	// Recommend Size : m < floor(log_2 (n))- 5.
	// Be aware that, if m > 8, Too SLOW. Because Time complexity is over O(2^n)
	P_value, isRandom, err = ApproximateEntropy(5, n)
	if err != nil {
		panic(err)
	}
	PrettyPrint_Add("Approximate Entropy Test", P_value, isRandom)

	// 2.13 Cumulative Sums (Cusum) Test
	P_values, isRandoms, err = CumulativeSums_All()
	if err != nil {
		panic(err)
	}
	PrettyPrint_Add_Array("Cumulative Sums (Cusum) Test", P_values, isRandoms)

	// 2.14 Random Excursions Test
	P_values, isRandoms, err = RandomExcursions(n)
	if err != nil {
		panic(err)
	}
	PrettyPrint_Add_Array("Random Excursions Test", P_values, isRandoms)

	// 2.15 Random Excursions Variant Test
	P_values, isRandoms, err = RandomExcursionsVariant(n)
	if err != nil {
		panic(err)
	}
	PrettyPrint_Add_Array("Random Excursions Variant Test", P_values, isRandoms)

	PrettyPrint_Render()
}
