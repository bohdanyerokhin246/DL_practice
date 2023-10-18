package main

import (
	"fmt"
	"math/big"
	"strconv"
)

type myBigInt struct {
	value *big.Int
	bin   []int
}

func (v *myBigInt) setHex(hexStr string) {

	hexInt, success := new(big.Int).SetString(hexStr, 16)
	if !success {
		fmt.Errorf("Invalid hexadecimal string: %s", hexStr)
	}

	binaryStr := hexInt.Text(2)

	bitArray := make([]int, len(binaryStr))
	for i, char := range binaryStr {
		if char == '0' {
			bitArray[i] = 0
		} else if char == '1' {
			bitArray[i] = 1
		} else {
			fmt.Errorf("Invalid character in binary string: %c", char)
		}
	}
	v.bin = bitArray

}

func (v *myBigInt) getHex() {

	var newSlice = v.bin

	res := len(newSlice) % 4

	if res != 0 {
		if res == 1 {
			v.bin = append(v.bin, 000)
			newSlice = make([]int, len(v.bin)+3)
			newSlice[0] = 0
			newSlice[1] = 0
			newSlice[2] = 0
			copy(newSlice[3:], v.bin)
		} else if res == 2 {
			newSlice = make([]int, len(v.bin)+2)
			newSlice[0] = 0
			newSlice[1] = 0
			copy(newSlice[2:], v.bin)
		} else if res == 3 {
			newSlice = make([]int, len(v.bin)+1)
			newSlice[0] = 0
			copy(newSlice[1:], v.bin)
		}
	}

	var temp string
	tempStr, hexStr := "", ""

	for i := len(newSlice) - 1; i > 0; {
		temp = strconv.Itoa(newSlice[i-3]) + strconv.Itoa(newSlice[i-2]) + strconv.Itoa(newSlice[i-1]) + strconv.Itoa(newSlice[i])
		i -= 4

		switch temp {
		case "0000":
			tempStr += "0"
		case "0001":
			tempStr += "1"
		case "0010":
			tempStr += "2"
		case "0011":
			tempStr += "3"
		case "0100":
			tempStr += "4"
		case "0101":
			tempStr += "5"
		case "0110":
			tempStr += "6"
		case "0111":
			tempStr += "7"
		case "1000":
			tempStr += "8"
		case "1001":
			tempStr += "9"
		case "1010":
			tempStr += "a"
		case "1011":
			tempStr += "b"
		case "1100":
			tempStr += "c"
		case "1101":
			tempStr += "d"

		case "1110":
			tempStr += "e"

		case "1111":
			tempStr += "f"
		}
	}

	for i := len(tempStr) - 1; i >= 0; i-- {
		str := tempStr[i]
		hexStr += string(str)
	}
	fmt.Println(hexStr)
}

func (v *myBigInt) INV() {

	bitArr := make([]int, len(v.bin))
	for i, _ := range v.bin {
		if v.bin[i] == 0 {
			bitArr[i] = 1
		} else if v.bin[i] == 1 {
			bitArr[i] = 0
		}
	}

	v.bin = bitArr

}

func (v *myBigInt) XOR(numberA []int, numberB []int) {

	newSlice1 := numberA
	newSlice2 := numberB
	var result []int

	if len(newSlice1) > len(newSlice2) {
		newSlice := make([]int, len(newSlice2))
		for i := 0; i < len(newSlice1)-len(newSlice2); i++ {
			newSlice = append(newSlice, 0)
			copy(newSlice[len(newSlice2)-len(numberB):], numberB)
			newSlice2 = newSlice
		}
	} else if len(newSlice1) < len(newSlice2) {
		newSlice := make([]int, len(newSlice1))
		for i := 0; i < len(newSlice2)-len(newSlice1); i++ {
			newSlice = append(newSlice, 0)
			copy(newSlice[len(numberB)-len(numberA):], numberA)
			newSlice1 = newSlice
		}
	}

	result = make([]int, len(newSlice1))

	for i := 0; i < len(result); i++ {
		if (newSlice1[i] == 1 || newSlice2[i] == 1) && newSlice1[i] != newSlice2[i] {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}

	v.bin = result

}

func (v *myBigInt) AND(numberA []int, numberB []int) {

	newSlice1 := numberA
	newSlice2 := numberB
	var result []int

	if len(newSlice1) > len(newSlice2) {
		newSlice := make([]int, len(newSlice2))
		for i := 0; i < len(newSlice1)-len(newSlice2); i++ {
			newSlice = append(newSlice, 0)
			copy(newSlice[len(newSlice2)-len(numberB):], numberB)
			newSlice2 = newSlice
		}
	} else if len(newSlice1) < len(newSlice2) {
		newSlice := make([]int, len(newSlice1))
		for i := 0; i < len(newSlice2)-len(newSlice1); i++ {
			newSlice = append(newSlice, 0)
			copy(newSlice[len(numberB)-len(numberA):], numberA)
			newSlice1 = newSlice
		}
	}

	result = make([]int, len(newSlice1))

	for i := 0; i < len(result); i++ {
		if newSlice1[i] == 1 && newSlice2[i] == 1 {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}

	v.bin = result
}

func (v *myBigInt) OR(numberA []int, numberB []int) {

	newSlice1 := numberA
	newSlice2 := numberB
	var result []int

	if len(newSlice1) > len(newSlice2) {
		newSlice := make([]int, len(newSlice2))
		for i := 0; i < len(newSlice1)-len(newSlice2); i++ {
			newSlice = append(newSlice, 0)
			copy(newSlice[len(newSlice2)-len(numberB):], numberB)
			newSlice2 = newSlice
		}
	} else if len(newSlice1) < len(newSlice2) {
		newSlice := make([]int, len(newSlice1))
		for i := 0; i < len(newSlice2)-len(newSlice1); i++ {
			newSlice = append(newSlice, 0)
			copy(newSlice[len(numberB)-len(numberA):], numberA)
			newSlice1 = newSlice
		}
	}

	result = make([]int, len(newSlice1))

	for i := 0; i < len(result); i++ {
		if newSlice1[i] == 1 || newSlice2[i] == 1 {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}

	v.bin = result
}

func (v *myBigInt) lShift(positions int) {
	length := len(v.bin)

	shiftedSlice := make([]int, length)

	for i, element := range v.bin {

		newIndex := (i + positions) % length
		shiftedSlice[newIndex] = element
	}
	v.bin = shiftedSlice
}

func (v *myBigInt) rShift(positions int) {
	length := len(v.bin)

	shiftedSlice := make([]int, length)

	for i, element := range v.bin {
		newIndex := (i - positions + length) % length
		shiftedSlice[newIndex] = element
	}
	v.bin = shiftedSlice
}

func main() {
	var numberA myBigInt
	var numberB myBigInt
	var numberC myBigInt

	numberA.setHex("51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4")

	numberB.setHex("403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c")

	numberA.INV()
	fmt.Println("INV A")
	numberA.getHex()

	numberB.INV()
	fmt.Println("INV B")
	numberB.getHex()

	numberC.XOR(numberA.bin, numberB.bin)
	fmt.Println("XOR")
	numberC.getHex()

	numberC.AND(numberA.bin, numberB.bin)
	fmt.Println("AND")
	numberC.getHex()

	numberC.OR(numberA.bin, numberB.bin)
	fmt.Println("OR")
	numberC.getHex()

	numberA.lShift(5)
	fmt.Println("L shift")
	numberA.getHex()

	numberA.rShift(5)
	fmt.Println("R shift")
	numberA.getHex()

}
