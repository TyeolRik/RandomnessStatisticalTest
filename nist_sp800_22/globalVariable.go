package nist_sp800_22

// This Variable is the unknown sequence whether random or not.
// The reason, why this variable is []uint8, is the minimum variable with regard to memory.
// I couldn't find the best and easiest variable which is almost same as std::bitset in C++.
var epsilon []uint8

var MAXLOG float64 = 7.09782712893383996732e2
var big float64 = 4.503599627370496e15
var biginv float64 = 2.22044604925031308085e-16
var MACHEP float64 = 1.38777878078144567553e-17

var __ERROR_float64__ float64 = 7.123456789e-16

func inputEpsilon(_input []uint8) {
	epsilon = _input

	// Revert
	for i, j := 0, len(epsilon)-1; i < j; i, j = i+1, j-1 {
		epsilon[i], epsilon[j] = epsilon[j], epsilon[i]
	}
}

func inputEpsilonAsString(_input string) {
	epsilon = []uint8{}
	for _, value := range _input {
		var temp uint8 = uint8(value - '0')
		if temp < 0 || temp > 1 {
			panic("inputEpsilonAsString :: ERROR Input is wrong")
		}
		epsilon = append(epsilon, uint8(value-'0'))
	}

	// Revert
	for i, j := 0, len(epsilon)-1; i < j; i, j = i+1, j-1 {
		epsilon[i], epsilon[j] = epsilon[j], epsilon[i]
	}
}
