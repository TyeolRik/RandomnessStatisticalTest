package main

import (
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
	// Revert
	for i, j := 0, len(epsilon)-1; i < j; i, j = i+1, j-1 {
		epsilon[i], epsilon[j] = epsilon[j], epsilon[i]
	}
}

// Discrete Fourier Transform
// 3.6 Discrete Fourier Transform (Specral) Test, Page 68.

// https://gist.github.com/r9y9/8095894
func DFT_naive(input []float64) ([]float64, []float64) {
	real := make([]float64, len(input))
	imag := make([]float64, len(input))
	arg := 2.0 * math.Pi / float64(len(input))
	for k := 0; k < len(input); k++ {
		r, i := 0.0, 0.0
		for n := 0; n < len(input); n++ {
			r += input[n] * math.Cos(arg*float64(n)*float64(k))
			i += input[n] * math.Sin(arg*float64(n)*float64(k))
		}
		real[k], imag[k] = r, i
	}
	return real, imag
}

// Wiki definition // https://en.wikipedia.org/wiki/Discrete_Fourier_transform#Definition
func DFT(X []float64) ([]float64, []float64) {
	var N float64 = float64(len(X))
	var real []float64 = make([]float64, len(X))
	var imag []float64 = make([]float64, len(X))

	var _2_pi_divide_N float64 = 2 * math.Pi / N
	for k := range X {
		var r float64 = 0.0
		var i float64 = 0.0
		for n := 0; n < len(X); n++ {
			r = r + X[n]*math.Cos(_2_pi_divide_N*float64(k)*float64(n))
			i = i + X[n]*math.Sin(_2_pi_divide_N*float64(k)*float64(n))
		}
		real[k] = r
		imag[k] = i
	}
	return real, imag
}

// https://numpy.org/doc/stable/reference/generated/numpy.absolute.html
// For complex input, a + ib, the absolute value is Sqrt(a^2 + b^2)
func Modulus(real, imag []float64) []float64 {
	modulus := make([]float64, len(real)/2)
	for i := 0; i < len(real)/2; i++ {
		modulus[i] = math.Sqrt(real[i]*real[i] + imag[i]*imag[i])
	}
	return modulus
}

// Definition of Peak
// An observation in an ordered series is said to be a “peak” if its value is greater than the value of its two neighbouring observations.
// Reference : [OECD - Statistical Terms](https://stats.oecd.org/glossary/detail.asp?ID=3780)
func theNumberOfPeaksLessThanT(input []float64, T float64) int {
	if len(input) < 3 {
		panic("input length of array is too small")
	}
	var count int = 0
	var numberOfPeak int = 0
	for i := 1; i < len(input)-1; i++ {
		if (input[i] > input[i-1]) && (input[i] > input[i+1]) {
			numberOfPeak++
			if input[i] < T {
				count++
			}
		}
	}
	// fmt.Println("The Number of Peaks", numberOfPeak)
	return count
}

func main() {
	inputEpsilonAsString("1100100100001111110110101010001000100001011010001100001000110100110001001100011001100010100010111000")
	// inputEpsilonAsString("1001010011")
	var n int = len(epsilon)
	var X []float64 = make([]float64, 0, n)
	for _, value := range epsilon {
		X = append(X, 2*float64(value)-1)
	}

	// (2) Apply a Discrete Fourier transform (DFT) on X to produce: S = DFT(X).
	// real, imag := DFT_naive(X)
	real, imag := DFT(X)

	// (3) Calculate M = modulus(S´) ≡ |S'|,
	// where S´ is the substring consisting of the first n/2 elements in S,
	// and the modulus function produces a sequence of peak heights.
	M := Modulus(real, imag)

	// (4) Compute T
	T := math.Sqrt(2.995732274 * float64(n)) // math.Log(1.0/0.05) = 2.995732273553991
	//fmt.Println("T", T)

	// (5) Compute N0
	N0 := 0.95 * float64(n) / 2.0
	//fmt.Println("N0", N0)

	// (6) Compute N1
	var N1 int

	// I don't know why NIST Source code doesn't use the definition of Peaks.
	// peak : https://stats.oecd.org/glossary/detail.asp?ID=3780
	// N1 = theNumberOfPeaksLessThanT(M, T)
	count := 0
	for _, value := range M {
		if value < T {
			count++
		}
	}
	N1 = count
	//fmt.Println("N1", count)

	// (7) Compute d
	d := (float64(N1) - N0) / math.Sqrt(float64(n)*0.95*0.05/4)
	//fmt.Println("d", d)

	// (8) Compute P-Value
	P_value := math.Erfc(math.Abs(d) / math.Sqrt2)
	//fmt.Println("P_value", P_value)
}
