/**
* From NIST SP800-22 Revision 1a.
* 2.7.1 Test Purpose
* The focus of this test is the number of occurrences of pre-specified target strings.
* The purpose of this test is to detect generators that produce too many occurrences of a given non-periodic (aperiodic) pattern.
* For this test and for the Overlapping Template Matching test of Section 2.8, an m-bit window is used to search for a specific m-bit pattern.
* If the pattern is not found, the window slides one bit position.
* If the pattern is found, the window is reset to the bit after the found pattern, and the search resumes.
 */

package nist_sp800_22

import (
	"errors"
	"fmt"
	"math"
)

// Input Size Recommendation
// Additionally, be sure that M > 0.01 * n and N=floor(n/M).
// Original Function in Official Document was func NonOverlappingTemplateMatching(m uint64, n uint64, eachBlockSize uint64)
// However, this parameter look slightly odd to Golang. So, I changed.
//
// Variable B
// The m-bit template to be matched
// B is a string of ones and zeros (of length m)
// which is defined in a template library of non-periodic patterns contained within the test code.
func NonOverlappingTemplateMatching(B []uint8, eachBlockSize uint64) (float64, bool, error) {

	// Original Parameter
	var m int = len(B)
	var n int = len(epsilon)

	var M uint64 = eachBlockSize // The length in bits of the substring of ε to be tested.
	var N uint64                 // The number of independent blocks. N has been fixed at 8 in the test code.

	if uint64(n)%M != 0 {
		errorMessage := fmt.Sprintf("Input, eachBlockSize=%v, is wrong. %v mod %v remains %v", eachBlockSize, n, M, uint64(n)%M)
		return __ERROR_float64__, false, errors.New(errorMessage)
	}
	N = (uint64(n) / M)

	// (1) Partition the sequence into N independent blocks of length M.
	var blocks [][]uint8 = make([][]uint8, N)
	var W []uint64 = make([]uint64, N) // W[j] (j = 0, …, N-1) be the number of times that B (the template) occurs within the block j.
	var partitionStart uint64 = 0
	var partitionEnd uint64 = M
	for j := range blocks {
		blocks[j] = epsilon[partitionStart:partitionEnd]
		partitionStart = partitionEnd
		partitionEnd = partitionEnd + M
	}

	// (2) Search for matches
	for j := range blocks {
		for bitPosition := 0; bitPosition <= int(M)-m; bitPosition++ {
			for i := range B {
				if blocks[j][bitPosition+i] != B[i] {
					goto UN_HIT
				}
			}
			W[j]++
			bitPosition = bitPosition + len(B) - 1
		UN_HIT:
		}
	}

	// (3) Compute the theoretical mean μ and variance σ2
	var mu, sigma2 float64
	var _float64_m float64 = float64(m)
	mu = float64(M-uint64(m)+1) / float64(math.Pow(2, _float64_m))
	sigma2 = float64(M) * (1/math.Pow(2, _float64_m) - float64(2*m-1)/float64(math.Pow(2, 2*_float64_m)))

	// (4) Compute χ2
	var chi_square float64 = 0
	for _, value := range W {
		chi_square = chi_square + math.Pow((float64(value)-mu), 2)/sigma2
	}

	// (5) Compute P-value
	var P_value float64 = igamc(float64(N)/2.0, chi_square/2.0)

	return P_value, DecisionRule(P_value, 0.01), nil
}
