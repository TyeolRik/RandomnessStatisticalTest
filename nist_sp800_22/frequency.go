/**
* From NIST SP800-22 Revision 1a.
* Test Purpose
* The focus of the test is the proportion of zeroes and ones for the entire sequence.
* The purpose of this test is to determine
* whether the number of ones and zeros in a sequence are approximately the
* same as would be expected for a truly random sequence.
* The test assesses the closeness of the fraction of ones to ½,
* that is, the number of ones and zeroes in a sequence should be about the same.
* All subsequent tests depend on the passing of this test.
 */

package nist_sp800_22

import (
	"errors"
	"math"
)

// Param n is The length of the bit string.
func Frequency(n uint64) (float64, bool, error) {

	// Step 1. Conversion to ±1
	var S_n int64 = 0
	for _, v := range epsilon {
		if v == 0 {
			S_n = S_n - 1
		} else if v == 1 {
			S_n = S_n + 1
		} else {
			return __ERROR_float64__, false, errors.New("one of input bits is neither 0 nor 1")
		}
	}

	// Step 2. Compute the test statistic S_obs
	var S_obs float64 = (math.Abs(float64(S_n)) / math.Sqrt(float64(len(epsilon))))

	// Step 3. Compute P-value
	var P_value float64 = math.Erfc(S_obs / math.Sqrt(2))

	return P_value, DecisionRule(P_value, LEVEL), nil

	/**
	* 2.1.5 Decision Rule (at the 1% Level)
	* If the computed P-value is < 0.01,
	* then conclude that the sequence is non-random.
	* Otherwise, conclude that the sequence is random.
	 */
}
