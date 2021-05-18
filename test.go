package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	b := make([]byte, 8)
	n, err := rand.Read(b)
	if n != 8 {
		panic(n)
	} else if err != nil {
		panic(err)
	}

	ret := make([]uint8, 64)
	retIndex := 0
	for _, value := range b {
		convert_uint8_to_bitString := fmt.Sprintf("%08b", value)
		for i := 0; i < 8; i++ {
			ret[retIndex] = convert_uint8_to_bitString[i] - 48
			retIndex = retIndex + 1
		}
	}

	fmt.Print(ret)
}
