// From NIST SP800-22 Revision 1a.
// 2.14.1 Test Purpose
// The focus of this test is the number of cycles having exactly K visits in a cumulative sum random walk.
// The cumulative sum random walk is derived from partial sums after the (0,1) sequence is transferred to the appropriate (-1, +1) sequence.
// A cycle of a random walk consists of a sequence of steps of unit length taken at random that begin at and return to the origin.
// The purpose of this test is to determine if the number of visits to a particular state within a cycle deviates from what one would expect for a random sequence.
// This test is actually a series of eight tests (and conclusions), one test and conclusion for each of the states: -4, -3, -2, -1 and +1, +2, +3, +4.

package nist_sp800_22

import (
	"math"
)

func RandomExcursions(n uint64) ([]float64, []bool, error) {

	var State_X []int64 = []int64{-4, -3, -2, -1, 1, 2, 3, 4}

	var X []int64 = make([]int64, n)

	// (1) Form a normalized (-1, +1) sequence X
	for i := range epsilon {
		X[i] = 2*int64(epsilon[i]) - 1
	}

	// (2) Compute the partial sums S[i] of successively larger subsequences.
	var S []int64 = make([]int64, n)
	var index_S uint64
	S[0] = X[0]
	for index_S = 1; index_S < n; index_S++ {
		S[index_S] = S[index_S-1] + X[index_S]
	}

	// (3) Form a new sequence S' by attaching zeros before and after the set S.
	var S_Prime []int64 = []int64{0}
	S_Prime = append(S_Prime, S...)
	S_Prime = append(S_Prime, 0)
	S = nil // nil will release the underlying memory to the garbage collector.

	// TyeolRik Note.
	// (4) ~ (7) Could be more effective code.
	// But, to showing like schoolbook, I coded as document steps.

	// (4) Calculate J = the total number of zero crossings in S', where a zero crossing is a value of zero in S ' that occurs after the starting zero.
	var J uint64 = 0 // the Number of Cycles
	for _, value := range S_Prime {
		if value == 0 {
			J++
		}
	}
	J = J - 1 // Not consisting 1st zero.

	// (5) Drawing Tables
	var Cycles [][]uint64 = make([][]uint64, 8)
	var CycleIndex int64 = -1 // Due to omit 1st zero.
	for i := range Cycles {
		Cycles[i] = make([]uint64, J)
	}
	for _, stateX := range S_Prime {
		switch stateX {
		case -4:
			Cycles[0][CycleIndex]++
		case -3:
			Cycles[1][CycleIndex]++
		case -2:
			Cycles[2][CycleIndex]++
		case -1:
			Cycles[3][CycleIndex]++
		case 0:
			CycleIndex++
		case 1:
			Cycles[4][CycleIndex]++
		case 2:
			Cycles[5][CycleIndex]++
		case 3:
			Cycles[6][CycleIndex]++
		case 4:
			Cycles[7][CycleIndex]++
		}
	}

	// (6) Count v_k(x) = the total number of cycles in which state x occurs exactly k times among all cycles.
	// What I understand is sum of count of State x (Cycle) row. (= Cycle Row)
	var v [8][6]uint64
	for rowIndex, CyclesRow := range Cycles {
		for _, occur := range CyclesRow {
			if occur < 5 {
				v[rowIndex][occur]++
			} else {
				v[rowIndex][5]++
			}
		}
	}

	/*
		// Print Log
		fmt.Println("J", J)
		for row, value := range v {
			fmt.Println(State_X[row], value, sumArray(value))
		}
	*/

	// (7) For each of the eight states of x, compute the test statistic χ^2
	var chi_square []float64 = make([]float64, len(State_X))
	for chi_square_Index, x := range State_X {
		// (7) - 1. Calculate theoretical probabilities π_0 ... π_5 (Page 85. Section 3.14)
		var _x float64 = float64(x)
		var pi [6]float64
		var tempArg_1_divided_by_2_abs_x float64 = 1.0 / (2.0 * math.Abs(_x))
		pi[0] = 1.0 - tempArg_1_divided_by_2_abs_x
		for k := 1; k <= 4; k++ {
			pi[k] = tempArg_1_divided_by_2_abs_x * tempArg_1_divided_by_2_abs_x * math.Pow((1.0-tempArg_1_divided_by_2_abs_x), float64(k-1))
		}
		pi[5] = tempArg_1_divided_by_2_abs_x * (1.0 - tempArg_1_divided_by_2_abs_x) * (1.0 - tempArg_1_divided_by_2_abs_x) * (1.0 - tempArg_1_divided_by_2_abs_x) * (1.0 - tempArg_1_divided_by_2_abs_x)

		// (7) - 2. compute the test statistic χ^2
		var sum float64 = 0.0
		var J_pi float64
		for k := 0; k <= 5; k++ {
			J_pi = float64(J) * pi[k]
			sum += (float64(v[chi_square_Index][k]) - J_pi) * (float64(v[chi_square_Index][k]) - J_pi) / J_pi
		}
		chi_square[chi_square_Index] = sum
	}

	var P_value []float64 = make([]float64, 8)
	var randomness []bool = make([]bool, 8)
	// fmt.Println("State=x", "\tCHI_SQUARE", "\t P-value", "\t\t Conclusion")
	for i := range P_value {
		P_value[i] = igamc(5.0/2.0, chi_square[i]/2.0)
		randomness[i] = DecisionRule(P_value[i], LEVEL)
		// fmt.Println(State_X[i], "\t", chi_square[i], "\t", P_value[i], "\t", DecisionRule(P_value[i], LEVEL))
	}

	return P_value, randomness, nil
}
