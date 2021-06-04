// From NIST SP800-22 Revision 1a.
// 2.9.1 Test Purpose
// The focus of this test is the number of bits between matching patterns (a measure that is related to the length of a compressed sequence).
// The purpose of the test is to detect whether or not the sequence can be significantly compressed without loss of information.
// A significantly compressible sequence is considered to be non-random.

package nist_sp800_22

import (
	"math"
)

func recommandedInputSize(n uint64) (L uint64, Q uint64) {
	if n >= 1059061760 {
		L = 16
		Q = 655360
	} else if n >= 496435200 {
		L = 15
		Q = 327680
	} else if n >= 231669760 {
		L = 14
		Q = 163840
	} else if n >= 107560960 {
		L = 13
		Q = 81920
	} else if n >= 49643520 {
		L = 12
		Q = 40960
	} else if n >= 22753280 {
		L = 11
		Q = 20480
	} else if n >= 10342400 {
		L = 10
		Q = 10240
	} else if n >= 4654080 {
		L = 9
		Q = 5120
	} else if n >= 2068480 {
		L = 8
		Q = 2560
	} else if n >= 904960 {
		L = 7
		Q = 1280
	} else if n >= 387840 {
		L = 6
		Q = 640
	} else {
		panic("length of test case is too small!")
	}
	return
}

func array2Binaryint(arr []uint8) uint64 {
	var numberOfDigit uint64 = 1
	var _index_T uint64 = 0

	// (1) Divide into L-bits
	for i := len(arr) - 1; i >= 0; i-- {
		_index_T += uint64(arr[i]) * numberOfDigit
		numberOfDigit *= 2
	}

	return _index_T
}

// Input Size Recommendation
// n >= (Q + K)L
// 6 <= L <= 16, Q = 10 * 2^{L}, K =floor(n/L)- Q ≈ 1000 * 2^{L}
// The values of L, Q and n should be chosen as follows
func Universal(L uint64, Q uint64, n uint64) (float64, bool, error) {
	// Pre-calculated Value from "Handbook of Applied Cryptography", Page 184. Table 5.3
	var expectedValue_mu [16]float64 = [16]float64{0.7326495, 1.5374383, 2.4016068, 3.3112247, 4.2534266, 5.2177052, 6.1962507, 7.1836656, 8.1764248, 9.1723243, 10.170032, 11.168765, 12.168070, 13.167693, 14.167488, 15.167379}
	var variance_sigma [16]float64 = [16]float64{0.690, 1.338, 1.901, 2.358, 2.705, 2.954, 3.125, 3.238, 3.311, 3.356, 3.384, 3.401, 3.410, 3.416, 3.419, 3.421}

	var K uint64 = (n / L) - Q
	//var _int_L = int(L)
	// var _float64_L float64 = float64(L)
	var _float64_Q float64 = float64(Q)

	var blocks [][]uint8 = make([][]uint8, 0, Q+K)
	var T []float64 = make([]float64, Q)

	// Divide into L-bits
	var blockNum uint64 = 0
	for {
		blocks = append(blocks, epsilon[blockNum*L:blockNum*L+L])
		blockNum++
		if blockNum >= Q+K {
			break
		}
	}

	var sum float64 = 0.0
	for blockNumber, eachBlocks := range blocks {
		var _blockNumber_float64 float64 = float64(blockNumber)

		// (2) the L-bit value is used as an index into the table
		var _index_T uint64 = array2Binaryint(eachBlocks)

		if _blockNumber_float64 < _float64_Q {
			// (2) The block number of the last occurrence of each L-bit block is noted in the table
			T[_index_T] = _blockNumber_float64 + 1.0
		} else {
			// (3) Examine each of the K blocks in the test segment and determine the number of blocks since the last occurrence of the same L-bit block (i.e., i – T[j]).
			sum += math.Log2(float64(blockNumber) + 1.0 - T[_index_T])
			T[_index_T] = float64(blockNumber) + 1.0
		}
	}

	// (4) Compute the test statistic
	var f_n float64 = sum / float64(K)

	// (5) Compute P-value

	// 5-1. Compute σ
	// var c float64 = 0.7 - 0.8/_float64_L + (4.0+32.0/_float64_L)*math.Pow(float64(K), -3.0/_float64_L)/15.0
	var P_value float64 = math.Erfc(math.Abs((f_n - expectedValue_mu[L-1]) / (math.Sqrt2 * variance_sigma[L-1])))
	// P_value := math.Erfc(math.Abs(son / mom))

	return P_value, DecisionRule(P_value, LEVEL), nil
}

func Universal_Recommended() (float64, bool, error) {
	var n uint64 = uint64(len(epsilon))
	L, Q := recommandedInputSize(n)
	return Universal(L, Q, n)
}
