package main

import (
	"fmt"
	"math"
)

func main() {
	var number1, number2, number3, number4 float32 = 3, 5, 7, 9 //Task 1 variables
	var cathetus1, cathetus2 float64 = 16, 9                    //Task 2 variables
	var inputNumber int = 485                                   //Task 3 variables
	var speedTube1, speedTube2 float32 = 7, 9                   //Task 4 variables

	sumOfFourNumber(number1, number2, number3, number4) //Task 1
	findAreaAndHypotenuse(cathetus1, cathetus2)         //Task 2
	reverseNumber(inputNumber)                          //Task 3
	poolFillingTime(speedTube1, speedTube2)             //Task 4
}

// Find the sum of four numbers.
func sumOfFourNumber(number1, number2, number3, number4 float32) {

	fmt.Println("/// Task 1 ///")
	fmt.Println("***INPUT***")
	fmt.Println("Number1 =", number1)
	fmt.Println("Number2 =", number2)
	fmt.Println("Number3 =", number3)
	fmt.Println("Number4 =", number4)

	fmt.Println("***OUTPUT***")
	fmt.Println("Sum of input number = ", number1+number2+number3+number4)
}

// Find the area and hypotenuse of a right triangle by cathetus
func findAreaAndHypotenuse(cathetus1, cathetus2 float64) {

	fmt.Println("/// Task 2 ///")
	fmt.Println("***INPUT***")
	fmt.Println("Сathetus1 =", cathetus1)
	fmt.Println("Сathetus2 =", cathetus2)

	var hypotenuse float64 = math.Sqrt(math.Pow(cathetus1, 2) + math.Pow(cathetus2, 2))
	var area float64 = (cathetus1 * cathetus2) / 2

	fmt.Println("***OUTPUT***")
	fmt.Println("Hypotenuse = ", hypotenuse)
	fmt.Println("Area = ", area)
}

// A three-digit number is given.
// Print the number that we get if we read the input number from right to left
func reverseNumber(inputNumber int) {

	fmt.Println("/// Task 3 ///")
	fmt.Println("***INPUT***")
	fmt.Println("Input number =", inputNumber)

	var outputNumber int = 0

	if inputNumber > 99 && inputNumber < 1000 {
		for inputNumber > 0 {
			buff := inputNumber % 10
			outputNumber *= 10
			outputNumber += buff
			inputNumber /= 10
		}
		fmt.Println("***OUTPUT***")
		fmt.Println("Output number =", outputNumber)
	} else {
		fmt.Println("Please enter a three-digit number")
	}
}

// One pipe will fill the pool in n days.
// Another pipe will fill the pool for k.
// Develop a program that determines how long it will take to fill a pool
// if two pipes fill the pool together
func poolFillingTime(speedTube1, speedTube2 float32) {

	fmt.Println("/// Task 4 ///")
	fmt.Println("***INPUT***")
	fmt.Println("Filling speed of the 1 tube", speedTube1)
	fmt.Println("Filling speed of the 2 tube", speedTube2)

	var poolFillingTime float32 = 1 / ((1 / speedTube1) + (1 / speedTube2))

	fmt.Println("***OUTPUT***")
	fmt.Println("Time = ", poolFillingTime, "days")
}
