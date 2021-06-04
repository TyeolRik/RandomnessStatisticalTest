// From NIST SP800-22 Revision 1a.
// 2.11.1 Test Purpose
// The focus of this test is the frequency of all possible overlapping m-bit patterns across the entire sequence.
// The purpose of this test is to determine whether the number of occurrences of the 2^m m-bit overlapping patterns is approximately the same as would be expected for a random sequence.
// Random sequences have uniformity; that is, every m-bit pattern has the same chance of appearing as every other m-bit pattern.
// Note that for m = 1, the Serial test is equivalent to the Frequency test of Section 2.1.

package nist_sp800_22

import (
	"math"
)

// Input Size Recommendation
// Choose m and n such that m < floor(log_2 (n))- 2.
func Serial(m uint64, n uint64) ([]float64, []bool, error) {

	var v [][]uint64 = make([][]uint64, 3)
	var section2_index uint64
	for section2_index = 0; section2_index <= 2; section2_index++ {
		// (1) Form an augmented sequence ε′:
		// Extend the sequence by appending the first m-1 bits to the end of the sequence for distinct values of n.
		if int64(m)-int64(section2_index)-1 < 0 {
			break
		}
		appendedEpsilon := append(epsilon, epsilon[0:m-section2_index-1]...)
		var blockSize uint64 = m - section2_index
		var blockIndex uint64
		v[section2_index] = make([]uint64, uint64(math.Pow(2.0, float64(blockSize))))

		// (2) Determine the frequency of all possible overlapping m-bit blocks
		// the frequency of all possible overlapping m-bit blocks
		for blockIndex = 0; blockIndex <= uint64(len(appendedEpsilon))-blockSize; blockIndex++ {
			for vIndex := range v[section2_index] {
				if isEqualBetweenBitsArray(appendedEpsilon[blockIndex:blockIndex+blockSize], Uint_To_BitsArray_size_N(uint64(vIndex), blockSize)) {
					v[section2_index][vIndex]++
				}
			}
		}
	}

	// (3) Compute ψ
	var psi [3]float64 = [3]float64{0, 0, 0} // ψ_m = psi[0] / ψ_{m-1} = psi[1] / ψ_{m-2} = psi[2]
	for i := range psi {
		if len(v[i]) == 0 {
			break
		}
		for _, value := range v[i] {
			psi[i] += float64(value) * float64(value)
		}
		psi[i] = math.Pow(2.0, float64(m)-float64(i))/float64(n)*psi[i] - float64(n)
		// CAUTION :: Possible to happen Floating-point error mitigation
	}
	//fmt.Println("PSI: ", psi)

	// (4) Compute ∇ψ^2 and ∇^2ψ^2
	delta1 := psi[0] - psi[1]
	delta2 := psi[0] - 2*psi[1] + psi[2]
	//fmt.Println("Delta1:", delta1)
	//fmt.Println("Delta2:", delta2)

	// (5) Compute P_value
	var tempArg float64 = math.Pow(2.0, float64(m-2))
	P_value1 := igamc(tempArg, delta1/2.0)
	P_value2 := igamc(tempArg/2.0, delta2/2.0)

	retP_value := []float64{P_value1, P_value2}
	retBools := []bool{DecisionRule(P_value1, LEVEL), DecisionRule(P_value2, LEVEL)}

	return retP_value, retBools, nil
}
