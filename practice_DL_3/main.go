package main

import "fmt"

func main() {
	//TASK 1
	var height, weight uint
	fmt.Print("/// TASK 1 ///\nInput height and weight. For example: 180 70\n")
	fmt.Scanln(&height, &weight)
	normBodyWeight(height, weight)

	//TASK 2
	var number uint
	fmt.Print("\n/// TASK 2 ///\nInput number for description\n")
	fmt.Scanln(&number)
	numberDescription(number)

	//TASK 3
	var day, month uint
	fmt.Print("\n/// TASK 3 ///\nInput day and month. For example: 25 6\n")
	fmt.Scanln(&day, &month)
	fmt.Print("Ви ")
	zodiacSign(day, month)
}

func normBodyWeight(height, weight uint) {

	if (height - weight) < 90 {
		fmt.Println("You need to lose weight")

	} else if (height-weight) >= 90 && (height-weight) <= 110 {
		fmt.Println("You have a normal weight")

	} else if (height - weight) > 110 {
		fmt.Println("You need to gain weight")

	}
}

func numberDescription(number uint) {
	var oddOrEven string
	var numberOfDigits string

	//Check number odd or even
	if number%2 == 0 {
		oddOrEven = "Number is even and "

	} else if number%2 == 1 {
		oddOrEven = "Number is odd and "
	}

	//Find number of digits
	if number > 1 && number < 10 {
		numberOfDigits = "number is one-digit"

	} else if number >= 10 && number < 100 {
		numberOfDigits = "number is two-digit"

	} else if number >= 100 && number < 1000 {
		numberOfDigits = "number is three-digit"

	}

	fmt.Println(oddOrEven + numberOfDigits)

}

func zodiacSign(day, month uint) {

	if month == 1 {
		if day < 20 {
			fmt.Println("Козеріг")
		} else {
			fmt.Println("Водолій")
		}

	} else if month == 2 {
		if day < 19 {
			fmt.Println("Водолій")
		} else {
			fmt.Println("Риби")
		}

	} else if month == 3 {
		if day < 20 {
			fmt.Println("Риби")
		} else {
			fmt.Println("Овен")
		}

	} else if month == 4 {
		if day < 20 {
			fmt.Println("Овен")
		} else {
			fmt.Println("Телець")
		}

	} else if month == 5 {
		if day < 21 {
			fmt.Println("Телець")
		} else {
			fmt.Println("Близнюки")
		}

	} else if month == 6 {
		if day < 21 {
			fmt.Println("Близнюки")
		} else {
			fmt.Println("Рак")
		}

	} else if month == 7 {
		if day < 22 {
			fmt.Println("Рак")
		} else {
			fmt.Println("Лев")
		}

	} else if month == 8 {
		if day < 23 {
			fmt.Println("Лев")
		} else {
			fmt.Println("Діва")
		}

	} else if month == 9 {
		if day < 23 {
			fmt.Println("Діва")
		} else {
			fmt.Println("Терези")
		}

	} else if month == 10 {
		if day < 23 {
			fmt.Println("Терези")
		} else {
			fmt.Println("Скорпіон")
		}

	} else if month == 11 {
		if day < 22 {
			fmt.Println("Скорпіон")
		} else {
			fmt.Println("Стрілець")
		}

	} else if month == 12 {
		if day < 21 {
			fmt.Println("Стрілець")
		} else {
			fmt.Println("Козеріг")
		}
	}
}
