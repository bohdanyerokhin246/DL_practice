package main

import "fmt"

func main() {
	swapMaxAndMinNumber()     //TASK 1
	amountOfSameNumber()      //TASK 2
	amountOfDifferentNumber() //TASk 3
}

func swapMaxAndMinNumber() {
	var array = [...]int{4, 3, 1, 5, 3, 14, 3, 4, 7, 10, 11, 12, 4, 3, 3, 56, 3, 12, 122, 3, 2, 3}
	var Maximum = array[0]
	var Minimum = array[0]

	//Find Max number
	for i := 0; i < len(array); i++ {
		if array[i] > Maximum {
			Maximum = array[i]
		}
	}

	//Find min number
	for i := 0; i < len(array); i++ {
		if array[i] < Minimum {
			Minimum = array[i]
		}
	}

	//Swap Maximum and Minimum number
	for i := 0; i < len(array); i++ {
		if array[i] == Maximum {
			array[i] = Minimum
		} else if array[i] == Minimum {
			array[i] = Maximum
		}
	}
	fmt.Println("///TASK 1///\nArray after change:", array)
}

func amountOfSameNumber() {
	var array = [...]int{4, 3, 1, 5, 3, 14, 3, 4, 7, 10, 11, 12, 4, 3, 3, 56, 3, 12, 122, 3, 2, 3}
	var amountOfDifferentNumber, counter int = 0, 0

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[i] == array[j] {
				counter++
			}
		}
		if counter > amountOfDifferentNumber {
			amountOfDifferentNumber = counter
		}
		counter = 0
	}

	fmt.Println("///TASK 2///\nAmount of same number =", amountOfDifferentNumber)
}

func amountOfDifferentNumber() {
	var array = [...]int{4, 3, 1, 5, 3, 14, 3, 4, 7, 10, 11, 12, 4, 3, 3, 56, 3, 12, 122, 3, 2, 3}
	var amountOfDifferentNumber int = 0
	for i := 0; i < len(array); i++ {

		var isNew bool = true
		for j := 0; j < i; j++ {
			if array[j] == array[i] {
				isNew = false
				continue
			}
		}
		if isNew == true {
			amountOfDifferentNumber++
		}
	}
	fmt.Println("///TASK 3///\nAmount of different number =", amountOfDifferentNumber)
}
