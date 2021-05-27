package nist_sp800_22

import (
	"io/ioutil"
)

const __FILE_CONSTANT_E_LOCATION_ string = "../assets/data.e"
const __FILE_CONSTANT_PI_LOCATION_ string = "../assets/data.pi"

const LEVEL float64 = 0.01 // for Decision Rule

// This Variable is the unknown sequence whether random or not.
// The reason, why this variable is []uint8, is the minimum variable with regard to memory.
// I couldn't find the best and easiest variable which is almost same as std::bitset in C++.
var epsilon []uint8

var MAXLOG float64 = 7.09782712893383996732e2
var big float64 = 4.503599627370496e15
var biginv float64 = 2.22044604925031308085e-16
var MACHEP float64 = 1.38777878078144567553e-17

var __ERROR_float64__ float64 = 7.123456789e-16

var CONSTANT_E []uint8
var CONSTANT_PI []uint8

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

func prepare_CONSTANT_E_asEpsilon() error {
	dat, err := ioutil.ReadFile(__FILE_CONSTANT_E_LOCATION_)
	var constant_E_binary []uint8
	if err != nil {
		panic(err)
	}
	constant_E_binary = make([]uint8, 0, len(dat))
	for _, value := range dat {
		switch value {
		case 48:
			constant_E_binary = append(constant_E_binary, 0)
		case 49:
			constant_E_binary = append(constant_E_binary, 1)
		}
	}
	epsilon = constant_E_binary
	return nil
}

func prepare_CONSTANT_PI_asEpsilon() error {
	dat, err := ioutil.ReadFile(__FILE_CONSTANT_PI_LOCATION_)
	var constant_PI_binary []uint8
	if err != nil {
		panic(err)
	}
	constant_PI_binary = make([]uint8, 0, len(dat))
	for _, value := range dat {
		switch value {
		case 48:
			constant_PI_binary = append(constant_PI_binary, 0)
		case 49:
			constant_PI_binary = append(constant_PI_binary, 1)
		}
	}
	epsilon = constant_PI_binary
	return nil
}
