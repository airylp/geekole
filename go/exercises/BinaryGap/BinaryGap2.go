package main

import (
	"fmt"
)

func Solution(number int) int {
	gap := 0
	init := false
	countGap := 0
	if number != 0 {
		for number != 0 {
			residuo := number % 2
			if residuo == 1 {
				init = true
				if countGap > gap {
					gap = countGap
				}
				countGap = 0
			}
			if residuo == 0 {
				if init {
					countGap++
				}
			}
			number = number / 2
		}
	}
	return gap
}

func main() {
	fmt.Println("Solution: ", Solution(5))
}
