// From NIST SP800-22 Revision 1a.
// 2.15.1 Test Purpose
// The focus of this test is the total number of times that a particular state is visited (i.e., occurs) in a cumulative sum random walk.
// The purpose of this test is to detect deviations from the expected number of visits to various states in the random walk.
// This test is actually a series of eighteen tests (and conclusions), one test and conclusion for each of the states: -9, -8, …, -1 and +1, +2, …, +9.

package nist_sp800_22

import (
	"math"
)

func RandomExcursionsVariant(n uint64) ([]float64, bool, error) {

	var State_X []int64 = []int64{-9, -8, -7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9}

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

	// From Here, There is difference between RandomExcursions and RandomExcursionsVariant

	// (4) For each of the eighteen non-zero states of x, compute ξ(x) = the total number of times that state x occurred across all J cycles.
	// (4) - 1. Calculate J, the number of Cycle = (the number of zero in S' - 1)	// The reason why -1 is omitting the 1st Zero.
	var J int64 = 0
	for _, value := range S_Prime {
		if value == 0 {
			J++
		}
	}
	J = J - 1 // Due to omit 1st zero.

	var ksi [18]int64
	// (4) - 2. Compute ξ
	for _, value := range S_Prime {
		if -9 <= value && value < 0 {
			ksi[value+9]++
		} else if 0 < value && value <= 9 {
			ksi[value+8]++
		}
	}

	// (5) For each ξ(x), Compute P-value
	var P_value []float64 = make([]float64, 18)
	var randomness []bool = make([]bool, 18)
	for i := range P_value {
		P_value[i] = math.Erfc(math.Abs(float64(ksi[i]-J)) / math.Sqrt(2.0*float64(J)*(4.0*math.Abs(float64(State_X[i]))-2.0)))
		randomness[i] = DecisionRule(P_value[i], LEVEL)
	}

	/*
		// Show Result in Terminal
		fmt.Println("J =", J)
		fmt.Println("--------------------------------------------------------------------------")
		fmt.Println("|    State(x)    |    Counts  ξ(x)    |    P_value    |    Conclusion    |")
		fmt.Println("--------------------------------------------------------------------------")
		for i := range P_value {
			if DecisionRule(P_value[i], LEVEL) {
				fmt.Printf("|      %2d        |        %04d        |   %.7f   |      Random      |\n", State_X[i], ksi[i], P_value[i])
			} else {
				fmt.Printf("|      %2d        |        %04d        |   %.7f   |    non-Random    |\n", State_X[i], ksi[i], P_value[i])
			}
		}
		fmt.Println("--------------------------------------------------------------------------")
	*/
	return P_value, false, nil
}
