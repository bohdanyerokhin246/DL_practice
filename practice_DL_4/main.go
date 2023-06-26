package main

import "fmt"

func main() {
	numberOfBacteria()    //TASk 1
	sumOfNumbers()        //TASK 2
	percentageOfNumbers() //TASK 3
}

func numberOfBacteria() {
	var numberOfBacteria int = 1

	for i := 1; i < 19; i++ {
		numberOfBacteria *= 2
	}
	fmt.Println("///TASK 1///\nNumber of bacteria =", +numberOfBacteria)
}

func sumOfNumbers() {
	var sumOfNumbers int = 0

	for i := 1; i < 1200000; i++ {
		if i%2 == 0 || i%3 == 0 {
			sumOfNumbers += i
		}
	}
	fmt.Println("///TASK 2///\nSum of numbers =", +sumOfNumbers)
}

func percentageOfNumbers() {
	var amountOfNumbers float32 = 0

	for i := 1; i < 1200000; i++ {
		if i%2 == 0 || i%5 == 0 {
			amountOfNumbers += 1

		}
	}
	var percentageOfNumbers float32 = (amountOfNumbers / 1200000) * 100
	fmt.Println("///TASK 3///\nPercentage of numbers =", +percentageOfNumbers, "%")
}
