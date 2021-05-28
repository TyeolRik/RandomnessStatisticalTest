package main

import (
	"fmt"
	"math"
)

var epsilon []uint8

func inputEpsilonAsString(_input string) {
	epsilon = []uint8{}
	for _, value := range _input {
		switch value {
		case '0':
			epsilon = append(epsilon, 0)
		case '1':
			epsilon = append(epsilon, 1)
		default:
			panic("inputEpsilonAsString :: ERROR Input is wrong")
		}
	}
	/*
		// Revert
		for i, j := 0, len(epsilon)-1; i < j; i, j = i+1, j-1 {
			epsilon[i], epsilon[j] = epsilon[j], epsilon[i]
		}
	*/
}

func main() {

	var B []uint8 = []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1}

	// Original Parameter
	var m int = len(B)
	//var n int = 1000000

	var M uint64 = 1032 // The length in bits of the substring of ε to be tested.
	//var N uint64 = (uint64(n) / M) // The number of independent blocks. N has been fixed at 8 in the test code.

	// (1) Partition the sequence into N independent blocks of length M.
	//var blocks [][]uint8 = make([][]uint8, N)
	var v []float64 = make([]float64, 6) // the number of occurrences of B in each block by incrementing an array v[i]
	//var partitionStart uint64 = 0
	//var partitionEnd uint64 = M
	/*
		for j := range blocks {
			blocks[j] = epsilon[partitionStart:partitionEnd]
			partitionStart = partitionEnd
			partitionEnd = partitionEnd + M
		}
		fmt.Println("N", N)

			var hit uint64 = 0
			// (2) Search for matches
				for j := range blocks {
					hit = 0
					for bitPosition := 0; bitPosition <= int(M)-m; bitPosition++ {
						for i := range B {
							if blocks[j][bitPosition+i] != B[i] {
								goto UN_HIT
							}
						}
						// Hit
						hit++
					UN_HIT:
					}
					// Misprint
					// In this part, There is example error. Page 40.
					if hit > 0 {
						if hit > 5 {
							v[5]++
						} else {
							v[hit]++
						}
					}
				}
	*/
	// v = []float64{0, 1, 1, 1, 1, 1}
	v = []float64{329, 164, 150, 111, 78, 136}

	// (3) Compute values for λand η
	// that will be used to compute the theoretical probabilities π_i corresponding to the classes of v0:
	var _float64_m float64 = float64(m)
	var lambda float64 = float64(float64(M)-_float64_m+1.0) / math.Pow(2.0, _float64_m)
	var eta float64 = lambda / 2.0
	fmt.Println("lambda\t", lambda)
	fmt.Println("eta\t", eta)

	// Page 40.
	// (4) Compute χ^2 as specified in Section 3.8 (Page. 74)
	var pi []float64 = []float64{0.364091, 0.185659, 0.139381, 0.100571, 0.070432, 0.139865} // On page 74
	// var pi []float64 = []float64{0.324652, 0.182617, 0.142670, 0.106645, 0.077147, 0.166269}
	//var pi2 []float64 = []float64{0.364091, 0.185659, 0.139381, 0.100571, 0.070432, 0.139865}
	var p float64 = math.Exp(-1 * eta)
	fmt.Println("P(U=0)\t", p)

	fmt.Println(pi)
	// Compute Probabilities
	/*
		sum := 0.0
		sum2 := 0.0
		K := 5
		for i := 0; i < K; i++ {
			pi[i] = Pr(i, eta)
			pi2[i] = Pr_ver2(i, eta)
			sum += pi[i]
			sum2 += pi2[i]
		}
		pi[K] = 1 - sum
		pi2[K] = 1 - sum2
		fmt.Printf("%.6f", pi)
		fmt.Printf("\n")
		fmt.Printf("%.6f", pi2)
		fmt.Printf("\n")
	*/

	var chi_square float64 = 0.0
	// var chi_square2 float64 = 0
	fmt.Println(v)
	for i := range v {
		var temp float64 = 5.0 * pi[i]
		//var temp2 float64 = 5.0 * pi2[i]
		chi_square += (v[i] - temp) * (v[i] - temp) / temp
		//chi_square2 += (v[i] - temp2) * (v[i] - temp2) / temp2
	}
	fmt.Println("chi_square\t", chi_square)
	//fmt.Println("chi_square2\t", chi_square2)
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func Sum(x uint64) float64 {
	if x == 0 {
		return 0.0
	}
	var ret float64 = 0.0
	var _x float64 = float64(x)
	var i float64
	for i = 0.0; i < _x; i = i + 1.0 {
		ret = ret + i
	}
	return ret
}

// https://encyclopediaofmath.org/wiki/Confluent_hypergeometric_function
func KummerFunction(a float64, lambda float64, z float64) float64 {
	var result float64 = 1.0
	var current float64 = 1.0
	var index uint64 = 1
	for {
		current *= (a + Sum(index-1)) / (lambda + Sum(index-1)) * z / float64(Factorial(index))
		result += current
		// fmt.Println(index, current)
		index++
		if current < 0.000000001 {
			break
		}
	}
	return result
}

// Reference : Page 74. function P(U = u)
func Pr_ver2(u int, eta float64) float64 {
	if u == 0 {
		return math.Exp(-1 * eta)
	} else {
		return eta * math.Exp(-2*eta) / math.Exp2(float64(u)) * KummerFunction(float64(u+1), 2, eta)
	}
}

// Reference : https://github.com/terrillmoore/NIST-Statistical-Test-Suite/blob/master/sts/src/overlappingTemplateMatchings.c#L95-L110
func Pr(u int, eta float64) float64 {
	var l int
	var sum, p float64

	if u == 0 {
		p = math.Exp(-1 * eta)
	} else {
		sum = 0.0
		for l = 1; l <= u; l++ {
			lgam_u, _ := math.Lgamma(float64(u))
			lgam_l, _ := math.Lgamma(float64(l))
			lgam_u_l_plus1, _ := math.Lgamma(float64(u - l + 1))
			sum += math.Exp(-1*eta - float64(u)*math.Log(2) + float64(l)*math.Log(eta) - lgam_u - lgam_l - lgam_u_l_plus1)
		}
		p = sum
	}
	return p
}
