package main

import (
	"fmt"

	. "github.com/tyeolrik/RandomnessStatisticalTest/nist_sp800_22"
)

func main() {

	readERR := Prepare_CONSTANT_E_asEpsilon()
	if readERR != nil {
		panic("FAILED to load natural E")
	}
	SetEpsilon(GetEpsilon()[0:100000])

	P_value, _, _ := Rank(100000)
	fmt.Printf("P-value : %f\n", P_value)
}
