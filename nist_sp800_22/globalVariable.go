package nist_sp800_22

// This Variable is the unknown sequence whether random or not.
// The reason, why this variable is []uint8, is the minimum variable with regard to memory.
// I couldn't find the best and easiest variable which is almost same as std::bitset in C++.
var epsilon []uint8

var MAXLOG float64 = 7.09782712893383996732e2
var big float64 = 4.503599627370496e15
var biginv float64 = 2.22044604925031308085e-16
var MACHEP float64 = 1.38777878078144567553e-17
