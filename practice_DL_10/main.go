package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var inputNumber int = 12345 //Task 1 variables
	fmt.Println("/// Task 1 ///\nInput number  =", inputNumber, "\nOutput number =", reverseNumber(inputNumber))

	var AC, BC, cornerACB float64 = 22, 21, math.Pi / 3 //Task 2 variables
	fmt.Println("/// Task 2 ///\nResult  =", arrivalOfCosines(AC, BC, cornerACB))

	input := "Я люблю програмування на мові C++" //Task 3 variables
	fmt.Println("/// Task 3 ///")
	actionsWithString(input)

}

func reverseNumber(inputNumber int) float64 {
	var outputNumber float64 = 0
	for inputNumber > 0 {
		buff := math.Mod(float64(inputNumber), 10)
		outputNumber *= 10
		outputNumber += buff
		inputNumber /= 10
		math.Remainder(float64(inputNumber), 10)
	}
	return outputNumber
}

func arrivalOfCosines(AC, BC, cornerACB float64) float64 {
	var result float64
	result = math.Pow(AC, 2) + math.Pow(BC, 2) - 2*AC*BC*math.Cos(cornerACB)
	return result
}

func actionsWithString(input string) {

	var ukrAlphabet = []string{"А", "Б", "В", "Г", "Ґ", "Д", "Е", "Є", "Ж", "З", "И", "І", "Ї",
		"Й", "К", "Л", "М", "Н", "О", "П", "Р", "С", "Т", "У", "Ф", "Х",
		"Ц", "Ч", "Ш", "Щ", "Ь", "Ю", "Я"}

	var engAlphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	var amountOfRepetitions int = 10
	subStr := "Go"
	fmt.Println(input)

	input = strings.Replace(input, "C++", subStr, 1)
	fmt.Println(input)

	input = strings.ToUpper(input)
	fmt.Println(input)

	fmt.Println(strings.Repeat(input+"\n", amountOfRepetitions))

	fmt.Println("\n***UKRAINIAN ALPHABET***")
	for i := 0; i < len(ukrAlphabet); i++ {
		fmt.Println("Letter", ukrAlphabet[i], "occurs", strings.Count(input, ukrAlphabet[i])*amountOfRepetitions, "times")

	}
	fmt.Println("\n***ENGLISH ALPHABET***")

	for i := 0; i < len(engAlphabet); i++ {
		fmt.Println("Letter", engAlphabet[i], "occurs", strings.Count(input, engAlphabet[i])*amountOfRepetitions, "times")
	}
}
