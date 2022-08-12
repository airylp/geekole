package main

import "fmt"

func Solution(A []int, K int) []int {
	if len(A) > 0 { //Correctness is important, prevent empty array (focus on correctness)
		for i := 0; i < K; i++ {
			aux := A[len(A)-1]
			for j := len(A) - 1; j > 0; j-- {
				A[j] = A[j-1]
			}
			A[0] = aux
		}
	}
	return A
}

func main() {
	A := []int{}
	fmt.Println("A", A)
	fmt.Println("A", Solution(A, 10))

	A2 := []int{9, 0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("A2", A2)
	fmt.Println("A2", Solution(A2, 10))
}
