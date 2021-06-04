// From NIST SP800-22 Revision 1a.
// 2.13.1 Test Purpose
// The focus of this test is the maximal excursion (from zero) of the random walk defined by the cumulative sum of adjusted (-1, +1) digits in the sequence.
// The purpose of the test is to determine
// whether the cumulative sum of the partial sequences occurring in the tested sequence is too large or too small
// relative to the expected behavior of that cumulative sum for random sequences.
// This cumulative sum may be considered as a random walk.
// For a random sequence, the excursions of the random walk should be near zero.
// For certain types of non-random sequences, the excursions of this random walk from zero will be large.

package nist_sp800_22

import (
	"math"
)

// param mode: A switch for applying the test either forward through the input sequence
//             mode = 0 : forward through the input sequence
//             mode = 1 : backward through the sequence
func CumulativeSums(mode int, n uint64) (float64, bool, error) {

	if n < 2 {
		panic("input n is too small. should be larger than 2")
	}

	var X []int8 = make([]int8, n)
	var S []int64 = make([]int64, n)

	// (1) Form a normalized sequence: The zeros and ones of the input sequence (ε) are converted to values X[i] of –1 and +1 using Xi = 2εi – 1.
	for i := range epsilon {
		X[i] = 2*int8(epsilon[i]) - 1
	}

	// (2) Compute partial sums S[i] of successively larger subsequences
	if mode == 0 {
		// Forward
		var _n int64 = int64(n)
		var index_S int64 = 0
		S[index_S] = int64(X[index_S])
		for index_S = 1; index_S < _n; index_S++ {
			S[index_S] = S[index_S-1] + int64(X[index_S])
		}
	} else if mode == 1 {
		// Backward
		var index_S uint64 = 0
		var index_X uint64 = n - 1
		S[index_S] = int64(X[index_X])
		for index_S = 1; index_S < n; index_S++ {
			index_X--
			S[index_S] = S[index_S-1] + int64(X[index_X])
		}
	} else {
		panic("Mode value is neither 0 nor 1")
	}

	// (3) Compute the test statistic z
	var z float64 = math.Abs(float64(S[0]))
	var now float64
	var index uint64 = 0
	for index = 1; index < n; index++ {
		now = math.Abs(float64(S[index]))
		if z < now {
			z = now
		}
	}

	// (4) Compute P-value (Refer 5.5.3)
	var P_value float64
	var term1, term2 float64
	var _n_float64 float64 = float64(n)

	var k int64
	var sqrt_n float64 = math.Sqrt(_n_float64)
	for k = int64((-1.0*_n_float64/z + 1.0) / 4.0); k <= int64((_n_float64/z-1.0)/4.0); k++ {
		term1 += CumulativeDistribution(float64(4*k+1)*z/sqrt_n) - CumulativeDistribution(float64(4*k-1)*z/sqrt_n)
	}
	for k = int64((-1.0*_n_float64/z - 3.0) / 4.0); k <= int64((_n_float64/z-1.0)/4.0); k++ {
		term2 += CumulativeDistribution(float64(4*k+3)*z/sqrt_n) - CumulativeDistribution(float64(4*k+1)*z/sqrt_n)
	}
	P_value = 1 - term1 + term2

	return P_value, DecisionRule(P_value, LEVEL), nil
}

func CumulativeSums_All() ([]float64, []bool, error) {
	forward_P, forward_R, _ := CumulativeSums(0, uint64(len(epsilon)))
	backward_P, backward_R, _ := CumulativeSums(1, uint64(len(epsilon)))
	return []float64{forward_P, backward_P}, []bool{forward_R, backward_R}, nil
}
