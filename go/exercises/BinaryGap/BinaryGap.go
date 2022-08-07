package main

import (
	"fmt"
)

func convertNumberToBinary(number int) *[]int {
	binaryArray := []int{}
	if number != 0 {
		for number != 0 {
			residuo := number % 2
			binaryArray = append(binaryArray, residuo)
			number = number / 2
		}
	} else {
		binaryArray = append(binaryArray, 0)
	}

	return &binaryArray
}

func Solution(number int) int {
	binaryArray := convertNumberToBinary(number)
	fmt.Println("binaryArray: ", binaryArray)
	counter := 0
	gap := 0
	init := false
	countGap := 0

	for counter < len(*binaryArray) {

		if (*binaryArray)[counter] == 1 {
			init = true
			fmt.Println("countGap:", countGap)
			if countGap > 1 && countGap > gap {
				gap = countGap
			}
			countGap = 0
		}
		if (*binaryArray)[counter] == 0 {
			if init {
				countGap++
			}
		}
		counter++
	}
	return gap
}

func main() {
	fmt.Println("Solution: ", Solution(5))
}
