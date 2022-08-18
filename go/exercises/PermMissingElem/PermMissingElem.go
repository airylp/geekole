package main

import "fmt"

func solution(A []int) int {
	l := len(A) + 1
	sum := (l + 1) * l / 2

	fmt.Println(sum)
	aux := 0
	for i := 0; i < len(A); i++ {
		aux = aux + A[i]
	}

	return sum - aux
}

func main() {
	fmt.Println(solution([]int{5, 2, 3, 4}))
}
