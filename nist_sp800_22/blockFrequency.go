/**
* From NIST SP800-22 Revision 1a.
* Test Purpose
* The focus of the test is the proportion of ones within M-bit blocks.
* The purpose of this test is to determine
* whether the frequency of ones in an M-bit block is approximately M/2,
* as would be expected under an assumption of randomness.
* For block size M=1, this test degenerates to test 1, the Frequency (Monobit) test.
 */

package nist_sp800_22

func piWithBaseI(M uint64, N uint64) []float64 {
	var sum uint64
	var _N = int(N)
	var _M = int(M)
	var ret = []float64{}
	for i := 1; i <= _N; i++ {
		sum = 0
		for j := 0; j < _M; j++ {
			// In Official document, j starts from 1 to 3 but, array index starts from 0 in computer (at least Golang and C++).
			// But, I don't know why, this document specified (i - 1) in the equation.
			sum = sum + uint64(epsilon[(i-1)*_M+j])
			// fmt.Print(epsilon[(i-1)*_M+j], " ")
		}
		// fmt.Printf("\n")
		ret = append(ret, float64(sum)/float64(M))
	}
	return ret
}

// Input Size Recommendation
// The block size M should be selected such that M >= 20, M > 0.01n and N < 100.
// n >= 100
func BlockFrequency(M uint64, n uint64) (float64, bool, error) {

	// (1) Partition the input sequence into N = floor(n / M) non-overlapping blocks
	var N uint64 = n / M
	/**
	* For example
	* if n = 10, M = 3 and ε= 0110011010, 3 blocks (N=3) would be created, consisting of 011, 001 and 101.
	* The final 0 would be discarded.
	 */

	// (2) Determine the proportion πi of ones in each M-bit block using the equation for 1 <= i <= N
	pi := piWithBaseI(M, N)

	// (3) Compute the X^2 statistic
	var tempSum float64 = 0
	for _, value := range pi {
		tempSum = tempSum + (value-0.5)*(value-0.5)
	}
	var X2_statistic float64 = 4 * float64(M) * tempSum

	// fmt.Println("X2_statistic", X2_statistic)
	//fmt.Println("n", n)
	//fmt.Println("M", M)
	//fmt.Println("N", N)

	// (4) Compute P-value
	var P_value float64 = igamc(float64(N)/2.0, X2_statistic/2.0)
	return P_value, DecisionRule(P_value, LEVEL), nil
}
