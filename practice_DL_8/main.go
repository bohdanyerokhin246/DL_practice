package main

import "fmt"

func main() {
	//TASK 1
	numberXForFirstTask, numberYForFirstTask := 1, 2
	fmt.Println("///TASK 1///\nNumber X before swapping =", numberXForFirstTask)
	fmt.Println("Number Y before swapping =", numberYForFirstTask)
	swapNumbers(&numberXForFirstTask, &numberYForFirstTask)
	fmt.Println("Number X after swapping =", numberXForFirstTask)
	fmt.Println("Number Y after swapping =", numberYForFirstTask)

	//TASK 2
	numberForSecondTask := 2
	fmt.Println("///TASK 2///\nNumber before change =", numberForSecondTask)
	changeValue(&numberForSecondTask)
	fmt.Println("Number after change =", numberForSecondTask)

	//TASK 3
	numberForThirdTask := 2
	fmt.Println("///TASK 3///\nNumber before multiply =", numberForThirdTask)
	multiplyByThree(&numberForThirdTask)
	fmt.Println("Number after multiply =", numberForThirdTask)
}

func swapNumbers(firstNumber, secondNumber *int) {
	buff := *firstNumber
	*firstNumber = *secondNumber
	*secondNumber = buff
}
func changeValue(number *int) {
	*number = *number * *number * *number
}

func multiplyByThree(number *int) {
	*number = *number * 3
}
