/**
* From NIST SP800-22 Revision 1a.
* 2.8.1 Test Purpose
* The focus of the Overlapping Template Matching test is the number of occurrences of pre-specified target strings.
* Both this test and the Non-overlapping Template Matching test of Section 2.7 use an m-bit window to search for a specific m-bit pattern.
* As with the test in Section 2.7, if the pattern is not found, the window slides one bit position.
* If the pattern is not found, the window slides one bit position.
* The difference between this test and the test in Section 2.7 is that when the pattern is found, the window slides only one bit before resuming the search.
 */

package nist_sp800_22

import (
	"fmt"
	"math"
)

// Original Function in Official Document was func NonOverlappingTemplateMatching(m uint64, n uint64, eachBlockSize uint64)
// However, this parameter look slightly odd to Golang. So, I changed.
//
// Variable B
// The m-bit template to be matched
// B is a string of ones and zeros (of length m)
// which is defined in a template library of non-periodic patterns contained within the test code.
func OverlappingTemplateMatching(B []uint8, eachBlockSize uint64) (float64, bool, error) {

	// Original Parameter
	var m int = len(B)
	var n int = len(epsilon)

	var M uint64 = eachBlockSize   // The length in bits of the substring of ε to be tested.
	var N uint64 = (uint64(n) / M) // The number of independent blocks. N has been fixed at 8 in the test code.

	// (1) Partition the sequence into N independent blocks of length M.
	var blocks [][]uint8 = make([][]uint8, N)
	var v []float64 = make([]float64, 6) // the number of occurrences of B in each block by incrementing an array v[i]
	var partitionStart uint64 = 0
	var partitionEnd uint64 = M
	for j := range blocks {
		blocks[j] = epsilon[partitionStart:partitionEnd]
		partitionStart = partitionEnd
		partitionEnd = partitionEnd + M
	}
	fmt.Println("N", N)

	var hit uint64 = 0
	// (2) Search for matches
	for j := range blocks {
		hit = 0
		for bitPosition := 0; bitPosition <= int(M)-m; bitPosition++ {
			for i := range B {
				if blocks[j][bitPosition+i] != B[i] {
					goto UN_HIT
				}
			}
			// Hit
			hit++
		UN_HIT:
		}
		// Misprint
		// In this part, There is example error. Page 40.
		if hit > 0 {
			if hit > 5 {
				v[5]++
			} else {
				v[hit]++
			}
		}
	}

	// (3) Compute values for λand η
	// that will be used to compute the theoretical probabilities π_i corresponding to the classes of v0:
	var _float64_m float64 = float64(m)
	var lambda float64 = (float64(M) - _float64_m + 1) / math.Pow(2, _float64_m)
	var eta float64 = lambda / 2.0
	fmt.Println("lambda\t", lambda)
	fmt.Println("eta\t", eta)

	// Page 40.
	// (4) Compute χ^2 as specified in Section 3.8 (Page. 74)
	var pi []float64 = []float64{0.364091, 0.185659, 0.139381, 0.100571, 0.070432, 0.139865} // On page 74
	// var pi []float64 = []float64{0.324652, 0.182617, 0.142670, 0.106645, 0.077147, 0.166269}
	var p float64 = math.Exp(-1 * eta)
	fmt.Println("P(U=0)\t", p)

	// Compute Probabilities
	sum := 0.0
	K := 5
	for i := 0; i < K; i++ {
		pi[i] = Pr(i, eta)
		sum += pi[i]
	}
	pi[K] = 1 - sum
	fmt.Println(pi)

	var chi_square float64 = 0
	for i := range v {
		var temp float64 = 5.0 * pi[i]
		chi_square += (v[i] - temp) * (v[i] - temp) / temp
	}
	fmt.Println("chi_square\t", chi_square)

	// (5) Compute P-value
	var P_value float64 = igamc(2.5, chi_square/2.0)
	// Misprint report : in Page 41. P-value = igamc(5.0/2.0, 3.167729/2.0) = 0.274932
	// But igamc(5.0/2.0, 3.167729/2.0) = 0.6741449650657756 in Cephes.

	return P_value, DecisionRule(P_value, 0.01), nil
}

/*
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func Sum(x uint64) float64 {
	if x == 0 {
		return 0.0
	}
	var ret float64 = 0.0
	var _x float64 = float64(x)
	var i float64
	for i = 0.0; i < _x; i = i + 1.0 {
		ret = ret + i
	}
	return ret
}

// https://encyclopediaofmath.org/wiki/Confluent_hypergeometric_function
func KummerFunction(a float64, lambda float64, z float64) float64 {
	var result float64 = 1.0
	var current float64 = 1.0
	var index uint64 = 1
	for {
		current *= (a + Sum(index-1)) / (lambda + Sum(index-1)) * z / float64(Factorial(index))
		result += current
		// fmt.Println(index, current)
		index++
		if current < 0.000000001 {
			break
		}
	}
	return result
}

// Reference : Page 74. function P(U = u)
func Pr_ver2(u int, eta float64) float64 {
	if u == 0 {
		return math.Exp(-1 * eta)
	} else {
		return eta * math.Exp(-2*eta) / math.Exp2(float64(u)) * KummerFunction(float64(u+1), 2, eta)
	}
}
*/

// Reference : https://github.com/terrillmoore/NIST-Statistical-Test-Suite/blob/master/sts/src/overlappingTemplateMatchings.c#L95-L110
func Pr(u int, eta float64) float64 {
	var l int
	var sum, p float64

	if u == 0 {
		p = math.Exp(-1 * eta)
	} else {
		sum = 0.0
		for l = 1; l <= u; l++ {
			lgam_u, _ := math.Lgamma(float64(u))
			lgam_l, _ := math.Lgamma(float64(l))
			lgam_u_l_plus1, _ := math.Lgamma(float64(u - l + 1))
			sum += math.Exp(-1*eta - float64(u)*math.Log(2) + float64(l)*math.Log(eta) - lgam_u - lgam_l - lgam_u_l_plus1)
		}
		p = sum
	}
	return p
}
