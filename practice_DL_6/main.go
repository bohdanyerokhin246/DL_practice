package main

import "fmt"

func main() {
	var array = []int{48, 96, 86, 68, 57, 82, 63, 70}
	findEvenAndSort(array)   //TASK 1
	increaseToMaximum(array) //TASK 2
	squareAndSum(array)      //TASK 2
}

func findEvenAndSort(array []int) {
	var arraySlice = make([]int, 0)
	var temp = 0

	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			arraySlice = append(arraySlice, array[i])
		}
	}

	for i := 0; i < len(arraySlice); i++ {
		for j := 0; j < len(arraySlice)-1; j++ {
			if arraySlice[j] > arraySlice[j+1] {
				temp = arraySlice[j]
				arraySlice[j] = arraySlice[j+1]
				arraySlice[j+1] = temp
			}
		}
	}

	fmt.Println("///TASK 1///\nSorted slice =", arraySlice)
}

func increaseToMaximum(array []int) {
	var arraySlice = make([]int, 0)
	var sizeOfSlice int = 5
	var temp = 0

	for i := 0; i < sizeOfSlice; i++ {
		if temp < array[i] {
			temp = array[i]
		}
	}

	for i := 0; i < sizeOfSlice; i++ {
		arraySlice = append(arraySlice, array[i]+temp)
	}

	fmt.Println("///TASK 2///\nIncreased slice =", arraySlice)
}

func squareAndSum(array []int) {
	var sizeOfSlice int = 4
	var sliceSum int = 0
	arraySlice := array[len(array)-sizeOfSlice:]

	for i := 0; i < len(arraySlice); i++ {
		if arraySlice[i] < arraySlice[len(arraySlice)-1] {
			arraySlice[i] *= arraySlice[i]
		}
		sliceSum += arraySlice[i]
	}

	fmt.Println("///TASK 3///\nSum of squared =", sliceSum)
}
