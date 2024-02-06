package main

import (
	"fmt"
	"strconv"
	"strings"
)

type bigInt struct {
	value  []uint64
	hexStr string
}

func (v *bigInt) setHex(hexStr string) {
	var bitArraySize int

	if len(hexStr)%16 != 0 {
		bitArraySize = (len(hexStr) / 16) + 1
		v.value = make([]uint64, bitArraySize)

	} else {
		bitArraySize = len(hexStr) / 16
		v.value = make([]uint64, bitArraySize)
	}

	//Calculate size for align array
	growthDifference := (bitArraySize * 16) - len(hexStr)

	//Align array
	growthString := strings.Repeat("0", growthDifference) + hexStr

	//Convert hex string to uint array
	for i := len(growthString) / 16; i > 0; i-- {
		subStr := growthString[i*16-16 : i*16]
		block, _ := strconv.ParseUint(subStr, 16, 64)
		v.value[i-1] = block
	}
}

func (v *bigInt) getHex() string {

	var result string
	//Convert uint array to hex string
	for _, block := range v.value {
		result += strconv.FormatUint(block, 16)
	}

	return result
}

func XOR(x []uint64, y []uint64) []uint64 {
	var result []uint64
	var tmpSlice []uint64

	if len(x) > len(y) {
		tmpSlice = make([]uint64, len(x))
		for i := len(y); i > 0; i-- {
			if i >= 1 {
				tmpSlice[i] = y[i-1]
			} else {
				tmpSlice[i] = 0
			}
		}
		result = make([]uint64, len(tmpSlice))
		for i := len(tmpSlice) - 1; i >= 0; i-- {

			result[i] = x[i] ^ tmpSlice[i]
		}

	} else if len(x) < len(y) {
		tmpSlice = make([]uint64, len(y))
		for i := len(x); i > 0; i-- {
			if i >= 1 {
				tmpSlice[i] = x[i-1]
			} else {
				tmpSlice[i] = 0
			}
		}
		result = make([]uint64, len(tmpSlice))

		for i := len(tmpSlice) - 1; i >= 0; i-- {
			result[i] = y[i] ^ tmpSlice[i]
		}
	} else {
		result = make([]uint64, len(x))
		for i := len(x) - 1; i >= 0; i-- {
			result[i] = x[i] ^ y[i]
		}
	}

	return result
}

func (v *bigInt) INV() []uint64 {
	var result []uint64
	result = make([]uint64, len(v.value))

	for i := len(v.value) - 1; i >= 0; i-- {
		result[i] = ^v.value[i]
	}

	return result
}

func OR(x []uint64, y []uint64) []uint64 {
	var result []uint64
	var tmpSlice []uint64

	if len(x) > len(y) {
		tmpSlice = make([]uint64, len(x))
		for i := len(y); i > 0; i-- {
			if i >= 1 {
				tmpSlice[i] = y[i-1]
			} else {
				tmpSlice[i] = 0
			}
		}
		result = make([]uint64, len(tmpSlice))
		for i := len(tmpSlice) - 1; i >= 0; i-- {

			result[i] = x[i] | tmpSlice[i]
		}

	} else if len(x) < len(y) {
		tmpSlice = make([]uint64, len(y))
		for i := len(x); i > 0; i-- {
			if i >= 1 {
				tmpSlice[i] = x[i-1]
			} else {
				tmpSlice[i] = 0
			}
		}
		result = make([]uint64, len(tmpSlice))

		for i := len(tmpSlice) - 1; i >= 0; i-- {
			result[i] = y[i] | tmpSlice[i]
		}
	} else {
		result = make([]uint64, len(x))
		for i := len(x) - 1; i >= 0; i-- {
			result[i] = x[i] | y[i]
		}
	}

	return result
}

func AND(x []uint64, y []uint64) []uint64 {
	var result []uint64
	var tmpSlice []uint64

	if len(x) > len(y) {
		tmpSlice = make([]uint64, len(x))
		for i := len(y); i > 0; i-- {
			if i >= 1 {
				tmpSlice[i] = y[i-1]
			} else {
				tmpSlice[i] = 0
			}
		}
		result = make([]uint64, len(tmpSlice))
		for i := len(tmpSlice) - 1; i >= 0; i-- {

			result[i] = x[i] & tmpSlice[i]
		}

	} else if len(x) < len(y) {
		tmpSlice = make([]uint64, len(y))
		for i := len(x); i > 0; i-- {
			if i >= 1 {
				tmpSlice[i] = x[i-1]
			} else {
				tmpSlice[i] = 0
			}
		}
		result = make([]uint64, len(tmpSlice))

		for i := len(tmpSlice) - 1; i >= 0; i-- {
			result[i] = y[i] & tmpSlice[i]
		}
	} else {
		result = make([]uint64, len(x))
		for i := len(x) - 1; i >= 0; i-- {
			result[i] = x[i] & y[i]
		}
	}

	return result
}

func (v *bigInt) shiftR(x int) []uint64 {

	var result []uint64

	result = make([]uint64, len(v.value))

	for i := len(v.value) - 1; i >= 0; i-- {

		result[i] = v.value[i] >> x
	}
	return result
}

func (v *bigInt) shiftL(x int) []uint64 {

	var result []uint64

	result = make([]uint64, len(v.value))

	for i := len(v.value) - 1; i >= 0; i-- {

		result[i] = v.value[i] << x

	}
	return result
}

func main() {

	var numberA bigInt
	var numberB bigInt
	var numberC bigInt

	numberA.setHex("51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4")
	numberB.setHex("fdb396dd73f6331c403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c")

	numberC.value = numberA.INV()
	fmt.Println("INV:\n", numberC.getHex())

	numberC.value = XOR(numberA.value, numberB.value)
	fmt.Println("XOR:\n", numberC.getHex())

	numberC.value = OR(numberA.value, numberB.value)
	fmt.Println("OR:\n", numberC.getHex())

	numberC.value = AND(numberA.value, numberB.value)
	fmt.Println("AND:\n", numberC.getHex())

	numberC.value = numberA.shiftR(2)
	fmt.Println("Shift right:\n", numberC.getHex())

	numberC.value = numberA.shiftL(2)
	fmt.Println("Shift left:\n", numberC.getHex())

}
