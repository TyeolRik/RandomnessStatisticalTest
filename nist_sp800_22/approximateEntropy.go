// From NIST SP800-22 Revision 1a.
// 2.12.1 Test Purpose
// As with the Serial test of Section 2.11, the focus of this test is the frequency of all possible overlapping m-bit patterns across the entire sequence.
// The purpose of the test is to compare the frequency of overlapping blocks of two consecutive/adjacent lengths (m and m+1) against the expected result for a random sequence.

package nist_sp800_22

import (
	"math"
)

// param m: The length of each block
// param n: len(epsilon) :: The length of the entire bit sequence.
func ApproximateEntropy(m uint64, n uint64) (float64, bool, error) {
	var psi [2]float64 // (5) Repeat twice
	var original_m = m

	for indexPSI := range psi {
		// (1) Augment the n-bit sequence to create n overlapping m-bit sequences by appending m-1 bits from the beginning of the sequence to the end of the sequence.
		appendedEpsilon := append(epsilon, epsilon[0:m-1]...)
		var two_raise_power_to_m uint64 = 1
		var tempVar uint64
		for tempVar = 0; tempVar < m; tempVar++ {
			two_raise_power_to_m = two_raise_power_to_m * 2
		}

		var possible_m_bits_blocks [][]uint8 = make([][]uint8, two_raise_power_to_m)
		var C []float64 = make([]float64, two_raise_power_to_m)
		var tempLastIndex uint64 = uint64(len(appendedEpsilon)) - m + 1
		var index uint64

		for bit := range possible_m_bits_blocks {
			possible_m_bits_blocks[bit] = Uint_To_BitsArray_size_N(uint64(bit), m)
			for index = 0; index < tempLastIndex; index++ {
				if isEqualBetweenBitsArray(appendedEpsilon[index:index+m], possible_m_bits_blocks[bit]) {
					C[bit] = C[bit] + 1
				}
			}
		}

		// (3) Compute C_{i}^{m}
		for indexC := range C {
			C[indexC] = C[indexC] / float64(n)
		}

		// (4) Compute PSI
		var sum float64 = 0.0
		for _, value := range C {
			if value > 0 {
				sum += value * math.Log(value)
			}
		}
		psi[indexPSI] = sum

		// (5) replacing m by m+1.
		m++
	}
	m = original_m // Return to original Block size

	// (6) Compute the test statistic χ^2
	var chi_square float64 = 2.0 * float64(n) * (math.Log(2) - (psi[0] - psi[1]))

	// (7) Compute P-value
	var P_value float64 = igamc(math.Pow(2.0, float64(m-1)), chi_square/2.0)
	return P_value, DecisionRule(P_value, 0.01), nil
}
