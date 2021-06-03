// From NIST SP800-22 Revision 1a.
// 2.10.1 Test Purpose
// The focus of this test is the length of a linear feedback shift register (LFSR).
// The purpose of this test is to determine whether or not the sequence is complex enough to be considered random.
// Random sequences are characterized by longer LFSRs.
// An LFSR that is too short implies non-randomness.

package nist_sp800_22

import (
	"math"
)

// From NIST SP800-22, linearComplexity.c
// Handbook of Applied Cryptography, Page 201.
// 6.30 Algorithm Berlekamp-Massey algorithm
func BerlekampMasseyAlgorithmFromNIST(s []uint8) uint64 {
	// 1. Initialization
	var n int = len(s)
	// fmt.Println("Length : ", n)
	var C []uint64 = make([]uint64, n)
	var B []uint64 = make([]uint64, n)
	var T []uint64 = make([]uint64, n)
	var P []uint64 = make([]uint64, n)

	L := 0
	m := -1
	var d uint64 = 0
	C[0] = 1
	B[0] = 1.0

	/* DETERMINE LINEAR COMPLEXITY */
	N_ := 0
	M := n // block size
	for {
		if N_ >= M {
			break
		}
		d = uint64(s[N_])
		for i := 1; i <= L; i++ {
			d += uint64(C[i]) * uint64(s[N_-i])
		}
		d = d % 2
		if d == 1 {
			for i := 0; i < M; i++ {
				T[i] = C[i]
				P[i] = 0
			}
			for j := 0; j < M; j++ {
				if B[j] == 1.0 {
					P[j+N_-m] = 1
				}
			}

			for i := 0; i < M; i++ {
				C[i] = (C[i] + P[i]) % 2
			}

			if L <= N_/2 {
				L = N_ + 1 - L
				m = N_
				for i := 0; i < M; i++ {
					B[i] = T[i]
				}
			}
		}
		N_++
		// fmt.Println("N:", N_, "d:", d)
	}
	return uint64(L)
}

// TyeolRik's Implementation with Go-lang
// Handbook of Applied Cryptography, Page 201.
// 6.30 Algorithm Berlekamp-Massey algorithm
func BerlekampMasseyAlgorithm(s []uint8) uint64 {
	// 1. Initialization
	var n float64 = float64(len(s))
	var L float64 = 0.0
	var m float64 = -1.0
	var N float64 = 0.0
	var C []float64 = make([]float64, len(s))
	var T []float64 = make([]float64, len(s))
	var B []float64 = make([]float64, len(s))
	C[0] = 1.0 // Other C[1] ... C[len(s)] would be all 0
	B[0] = 1.0 // Other B[1] ... B[len(s)] would be all 0

	var d int64 // discrepancy

	// 2. While (N < n) do the follwing:
	for {
		if N >= n {
			break
		}
		// 2.1 Compute the next discrepancy d.
		d = int64(s[int(N)])
		for i := 1; i <= int(L); i++ {
			d += int64(C[i]) * int64(s[int(N)-i])
		}
		d = d % 2

		// 2.2 If d = 1 then do the following
		if d == 1 {
			for i := 0; i < len(s); i++ {
				T[i] = C[i]
				C[i] = C[i] + B[i]*math.Pow(float64(i), (N-m))
			}
			if int64(L) <= int64(N/2) {
				L = N + 1 - L
				m = N
				for i := range B {
					B[i] = T[i]
				}
			}
		}

		// 2.3
		N++
		// fmt.Println("N:", N, "d:", d)
	}

	return uint64(L)
}

func LinearComplexity(M uint64, n uint64) (float64, bool, error) {
	// var K uint64 // The number of degrees of freedom

	// (1) Partition the n-bit sequence into N independent blocks of M bits, where n = MN.
	var N uint64 = n / M
	var blocks = make([][]uint8, 0, N)
	var i uint64
	for i = 0; i < N; i++ {
		if i*M+M > n {
			// Discard or Error
			break
		}
		blocks = append(blocks, epsilon[i*M:i*M+M])
	}

	// (2) Using the Berlekamp-Massey algorithm, determine the linear complexity L[i] of each of the N blocks (i = 0,…,N-1).
	var L []uint64 = make([]uint64, N)
	for i := range L {
		L[i] = BerlekampMasseyAlgorithmFromNIST(blocks[i])
	}
	// fmt.Println(L)

	// (3) Under an assumption of randomness, calculate the theoretical mean μ:
	var mu float64 = float64(M)/2.0 + (9.0+math.Pow(-1.0, float64(M+1)))/36.0 - (float64(M)/3.0+0.2222222222)/math.Pow(2.0, float64(M))
	//fmt.Println("MU : ", mu)

	// (4) For each substring, calculate a value of T[i]
	var T []float64 = make([]float64, N)
	var v []float64 = make([]float64, 7)
	for i := range T {
		T[i] = math.Pow(-1.0, float64(M))*(float64(L[i])-mu) + 2.0/9.0

		// (5) Record the Ti values in v0,…, v6 as follows:
		if T[i] <= -2.5 {
			v[0]++
		} else if T[i] <= -1.5 {
			v[1]++
		} else if T[i] <= -0.5 {
			v[2]++
		} else if T[i] <= 0.5 {
			v[3]++
		} else if T[i] <= 1.5 {
			v[4]++
		} else if T[i] <= 2.5 {
			v[5]++
		} else {
			v[6]++
		}
	}
	// fmt.Println(T)
	//fmt.Println("v", v)

	// (6) Compute χ^2

	// the probabilities computed by the equations in Section 3.10. (Page 79)
	// var _PI []float64 = make([]float64, 7)
	var _PI []float64 = []float64{0.010417, 0.03125, 0.125, 0.5, 0.25, 0.0625, 0.020833}
	var K int = 6

	/* Computing of PI is un-implemented.
	var center int = 7 / 2
	_PI[7/2] = 0.5
	for i := 1; i <= K/2; i++ {
		// Section 3.10 Equation (9)
		_PI[center-i] = 1.0 / (3.0 * math.Pow(2.0, float64(2*i+1)-1.0))
		// Section 3.10 Equation (8)
		_PI[center+i] = 1.0 / (3.0 * math.Pow(2.0, float64(2*i)-2.0))
	}
	*/
	//fmt.Println(_PI)

	var chi_square float64 = 0.0
	var N_pi float64
	for i := 0; i <= K; i++ {
		N_pi = float64(N) * _PI[i]
		chi_square += (v[i] - N_pi) * (v[i] - N_pi) / N_pi

	}
	// fmt.Println("x^2(obs):", chi_square)

	var P_value float64 = igamc(float64(K)/2.0, chi_square/2.0)

	return P_value, DecisionRule(P_value, LEVEL), nil
}
