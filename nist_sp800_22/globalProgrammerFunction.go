package nist_sp800_22

func Uint_To_BitsArray(input uint64) (bitArray []uint8) {

	/*
		if input == 0 {
			panic("input should be over 0 or use another function :: Uint_To_BitsArray_size_N")
		}
	*/
	var quotient, remainer uint64
	for {
		quotient = input / 2
		remainer = input - quotient*2
		bitArray = append(bitArray, uint8(remainer))
		input = quotient
		if quotient == 0 {
			break
		}
	}

	// Revert
	for i, j := 0, len(bitArray)-1; i < j; i, j = i+1, j-1 {
		bitArray[i], bitArray[j] = bitArray[j], bitArray[i]
	}
	return
}

func Uint_To_BitsArray_size_N(input uint64, N uint64) (bitArray []uint8) {
	bitArray = make([]uint8, N)
	if input == 0 {
		return
	}

	var index uint64 = N - 1
	var quotient, remainer uint64
	for {
		quotient = input / 2
		remainer = input - quotient*2
		bitArray[index] = uint8(remainer)
		index--
		input = quotient
		if quotient == 0 {
			return
		}
	}
}

func isEqualBetweenBitsArray(a []uint8, b []uint8) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
