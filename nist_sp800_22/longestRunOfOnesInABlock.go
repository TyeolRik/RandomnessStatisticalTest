/**
* From NIST SP800-22 Revision 1a. Page 29.
* 2.4.1. Test Purpose
* The focus of the test is the longest run of ones within M-bit blocks.
* The purpose of this test is to determine
* whether the length of the longest run of ones within the tested sequence is consistent
* with the length of the longest run of ones that would be expected in a random sequence.
* Note that an irregularity in the expected length of the longest run of ones implies
* that there is also an irregularity in the expected length of the longest run of zeroes.
* Therefore, only a test for ones is necessary.
 */

package nist_sp800_22

import "fmt"

var _PI_K3_M8 [4]float64 = [4]float64{0.2148, 0.3672, 0.2305, 0.1875}
var _PI_K5_M128 [6]float64 = [6]float64{0.1174, 0.2430, 0.2493, 0.1752, 0.1027, 0.1124}
var _PI_K5_M512 [6]float64 = [6]float64{0.1170, 0.2460, 0.2523, 0.1755, 0.1027, 0.1124}
var _PI_K5_M1000 [6]float64 = [6]float64{0.1307, 0.2437, 0.2452, 0.1714, 0.1002, 0.1088}
var _PI_K6_M10000 [7]float64 = [7]float64{0.0882, 0.2092, 0.2483, 0.1933, 0.1208, 0.0675, 0.0727}

func LongestRunOfOnes(n uint64) (float64, error) {
	var M uint64 // The length of each block.
	var N uint64 // The number of blocks; selected in accordance with the value of M.
	var K uint64

	if n < 128 {
		err := fmt.Errorf("input length of sequence is too small. (n = %d < 128)", n)
		return __ERROR_float64__, err
	} else if n < 6272 {
		M = 8
		N = n / 8
		K = 3
	} else if n < 750000 {
		M = 128
		N = n / 128
		K = 5
	} else {
		M = 10000
		N = n / 10000
		K = 6
	}

	// Divide the sequence into M-bit blocks.
	// sub_epsilons := [][]uint8{}
	sliceBoundary_start := 0
	sliceBoundary_end := 8
	v := [7]uint64{0, 0, 0, 0, 0, 0, 0}
	for {
		sub := epsilon[sliceBoundary_start:sliceBoundary_end]
		var longest uint64 = 0
		var count uint64 = 0
		for _, value := range sub {
			if value == 0 {
				longest = Max(longest, count)
				count = 0
			} else {
				count++
			}
		}
		longest = Max(longest, count)

		// Tabulate the frequencies νi of the longest runs of ones in each block into categories,
		// where each cell contains the number of runs of ones of a given length.
		switch K {
		case 3:
			if longest <= 1 {
				v[0]++
			} else if longest == 2 {
				v[1]++
			} else if longest == 3 {
				v[2]++
			} else {
				v[3]++
			}
		case 5:
			if longest <= 4 {
				v[0]++
			} else if longest == 5 {
				v[1]++
			} else if longest == 6 {
				v[2]++
			} else if longest == 7 {
				v[3]++
			} else if longest == 8 {
				v[4]++
			} else {
				v[5]++
			}
		case 6:
			if longest <= 10 {
				v[0]++
			} else if longest == 11 {
				v[1]++
			} else if longest == 12 {
				v[2]++
			} else if longest == 13 {
				v[3]++
			} else if longest == 14 {
				v[4]++
			} else if longest == 15 {
				v[5]++
			} else {
				v[6]++
			}
		default:
			panic("__MODE__ value is unexpected!")
		}

		// sub_epsilons = append(sub_epsilons, sub)
		sliceBoundary_start = sliceBoundary_start + 8
		sliceBoundary_end = sliceBoundary_end + 8
		if sliceBoundary_end > len(epsilon) {
			break
		}
	}
	// (3) Compute Test Statistic and Reference Distribution χ^2
	var chi_square float64 = 0
	var i uint64
	switch K {
	case 3:
		for i = 0; i <= K; i++ {
			var __v float64 = float64(v[i])
			var __N float64 = float64(N)
			var __PI float64 = _PI_K3_M8[i]
			var __temp float64 = (__v - __N*__PI) * (__v - __N*__PI) / (__N * __PI)
			chi_square = chi_square + __temp
		}
	case 5:
		for i = 0; i <= K; i++ {
			var __v float64 = float64(v[i])
			var __N float64 = float64(N)
			var __PI float64
			switch M {
			case 128:
				__PI = _PI_K5_M128[i]
			case 512:
				__PI = _PI_K5_M512[i]
			case 1000:
				__PI = _PI_K5_M1000[i]
			default:
				panic("LongestRunOfOnes :: ERROR :: M is unexpected")
			}
			var __temp float64 = (__v - __N*__PI) * (__v - __N*__PI) / (__N * __PI)
			chi_square = chi_square + __temp
		}
	case 6:
		for i = 0; i <= K; i++ {
			var __v float64 = float64(v[i])
			var __N float64 = float64(N)
			var __PI float64 = _PI_K6_M10000[i]
			var __temp float64 = (__v - __N*__PI) * (__v - __N*__PI) / (__N * __PI)
			chi_square = chi_square + __temp
		}
	}

	// (4) Compute P-value
	// var P_value float64 = igamc(float64(K)/2.0, chi_square/2.0)
	P_value := igamc(float64(K)/2.0, chi_square/2.0)

	return P_value, nil

	/**
	* 2.4.5. Decision Rule (at the 1% Level)
	* If the computed P-value is < 0.01, then conclude that the sequence is non-random.
	* Otherwise, conclude that the sequence is random.
	 */
}
