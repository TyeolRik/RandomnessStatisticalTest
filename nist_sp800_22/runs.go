/**
* From NIST SP800-22 Revision 1a.
* Test Purpose
* The focus of this test is the total number of runs in the sequence,
* where a run is an uninterrupted sequence of identical bits.
* A run of length k consists of exactly k identical bits and is bounded before and after with a bit of the opposite value.
* The purpose of the runs test is to determine
* whether the number of runs of ones and zeros of various lengths is as expected for a random sequence.
* In particular, this test determines whether the oscillation between such zeros and ones is too fast or too slow.
 */

package nist_sp800_22

import (
	"fmt"
	"math"
)

// Runs function returns "The total number of runs" across all n bits.
// the total number of zero runs + the total number of one-runs
func Runs(n uint64) (float64, bool, error) {
	var pi float64 = 0
	var _n_float64 = float64(n) // For Speed

	for _, value := range epsilon {
		pi = pi + float64(value)
	}
	pi = pi / _n_float64

	// (2) Determine if the prerequisite Frequency test is passed
	var tau float64 = 2.0 / math.Sqrt(_n_float64) // Note that for this test, var τ(tau) has been pre-defined in the test code.
	if math.Abs(pi-(1.0/2.0)) >= tau {
		// then the Runs test need not be performed
		return __ERROR_float64__, false, fmt.Errorf("the Runs test need not be performed! Because (%f) >= (tau = %f)", math.Abs(pi-(1.0/2.0)), tau)
	}

	// Compute the test statistic V_n
	var V_n float64 = 0
	for i := 0; i < len(epsilon)-1; i++ {
		if epsilon[i] == epsilon[i+1] {
			V_n = V_n + 0
		} else {
			V_n = V_n + 1
		}
	}
	V_n = V_n + 1

	var P_value float64 = math.Erfc(math.Abs(V_n-2*_n_float64*pi*(1-pi)) / (2 * math.Sqrt(2.0*_n_float64) * pi * (1 - pi)))
	return P_value, DecisionRule(P_value, LEVEL), nil

	/**
	* 2.3.5. Decision Rule (at the 1% Level)
	* If the computed P-value is < 0.01,
	* then conclude that the sequence is non-random.
	* Otherwise, conclude that the sequence is random.
	 */
}
