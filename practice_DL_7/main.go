package main

import "fmt"

func main() {
	sumRange(10, 20)          //TASK 1
	areaOfMiddleCircle(5, 10) //TASK 2
	amountOfDigits(321)       //TASK 3
	amountOfDigits(1243)      //TASK 3
	amountOfDigits(12543)     //TASK 3
}

func sumRange(firstNumber, secondNumber int) {
	var sum int
	if firstNumber > secondNumber {
		fmt.Println("///TASK 1///\n0", sum)
	} else {
		for i := firstNumber; i <= secondNumber; i++ {
			sum += i
		}
		fmt.Println("///TASK 1///\nSum of numbers =", sum)
	}
}

func areaOfMiddleCircle(radiusOfSmall, radiusOfBig float64) {
	var Pi float64 = 3.14
	var areaOfCircle float64

	if radiusOfSmall < 0 || radiusOfBig < 0 || radiusOfSmall > radiusOfBig {
		fmt.Println("Value of radius incorrect")
	} else {
		areaOfCircle = (Pi * (radiusOfBig * radiusOfBig)) - (Pi * (radiusOfSmall * radiusOfSmall))
		fmt.Println("///TASK 2///\nArea of middle circle =", areaOfCircle)
	}
}

func amountOfDigits(number int) {
	var amount = 0
	if number < 0 {
		fmt.Println("///TASK 3///\nNumber incorrect")
	}
	for number > 0 {
		number /= 10
		amount++
	}

	fmt.Println("///TASK 3///\nAmount of digits in number =", amount)
}
