/**
* From NIST SP800-22 Revision 1a. Page 32.
* 2.5.1 Test Purpose
* The focus of the test is the rank of disjoint sub-matrices of the entire sequence.
* The purpose of this test is to check for linear dependence among fixed length substrings of the original sequence.
* Note that this test also appears in the DIEHARD battery of tests.
 */

package nist_sp800_22

import (
	"math"
)

func Rank(n uint64) (float64, bool, error) {
	var M uint64 = 32 // The number of rows in each matrix.
	var Q uint64 = 32 // The number of columns in each matrix.
	var R []uint64    // Rank
	var F []uint64    // the Number of Matrices with R_l = index (index means, rank)

	// (1) Sequentially divide the sequence into M•Q-bit disjoint blocks
	N := n / (M * Q)
	R = make([]uint64, N)
	F = make([]uint64, M+1)

	epsilonIndex := 0
	matrices := make([][][]uint8, N)
	for i := 0; i < int(N); i++ {
		matrices[i] = make([][]uint8, M)
		for j := 0; j < int(M); j++ {
			matrices[i][j] = make([]uint8, Q)
			for k := 0; k < int(Q); k++ {
				matrices[i][j][k] = epsilon[epsilonIndex]
				epsilonIndex++
			}
		}
	}

	for index, matrix := range matrices {
		// (2) Determine the binary rank ( R ) of each matrix, where l = 1,...,N.
		// The method for determining the rank is described in Appendix A. - Page 33.
		R[index] = RankComputationOfBinaryMatrices(matrix)

		// (3) Let F_M = number of matrices with R_l = M (full rank),
		F[R[index]]++
	}

	// (4) Compute χ^2
	var __F_M_float64 = float64(F[M])
	var __F_M_minus_one_float64 = float64(F[M-1])
	var __N_float64 = float64(N)
	var chi_square float64 = (__F_M_float64-0.2888*__N_float64)*(__F_M_float64-0.2888*__N_float64)/(0.2888*__N_float64) + (__F_M_minus_one_float64-0.5776*__N_float64)*(__F_M_minus_one_float64-0.5776*__N_float64)/(0.5776*__N_float64) + (__N_float64-__F_M_float64-__F_M_minus_one_float64-0.1336*__N_float64)*(__N_float64-__F_M_float64-__F_M_minus_one_float64-0.1336*__N_float64)/(0.1336*__N_float64)

	//fmt.Println("N", N, __N_float64)
	//fmt.Println("F_M", F[M], __F_M_float64)
	//fmt.Println("F_M-1", F[M-1], __F_M_minus_one_float64)
	//fmt.Println("X^2 : ", chi_square)
	// (5) Compute P_Value
	var P_value float64 = math.Pow(math.E, -1*chi_square/2)

	/**
	* 2.5.5 Decision Rule (at the 1% Level)
	* If the computed P-value is < 0.01, then conclude that the sequence is non-random.
	* Otherwise, conclude that the sequence is random.
	 */

	return P_value, DecisionRule(P_value, LEVEL), nil
}
