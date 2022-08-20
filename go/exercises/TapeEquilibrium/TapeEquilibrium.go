package main

import "fmt"

func Solution(A []int) int {
	sum := 0
	for i := 0; i < len(A); i++ {
		sum = sum + A[i]
	}

	aux := 0
	res := sum
	for i := 0; i < len(A)-1; i++ {
		aux = aux + A[i]
		cal := sum - aux
		difference := abs(cal - aux)
		if i == 0 || difference < res {
			res = difference
		}
	}

	return res
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func main() {
	fmt.Println(Solution([]int{-1000, 1000}))
}
