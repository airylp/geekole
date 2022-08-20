package main

import "fmt"

func Solution(A []int) int {
	m := make(map[int]int, len(A))

	for i := 0; i < len(A); i++ {
		if _, exist := m[A[i]]; exist {
			m[A[i]] = m[A[i]] + 1
		} else {
			m[A[i]] = 1
		}
	}
	/*for i := 0; i < len(A); i++ {
		if m[A[i]]%2 == 1 {
			return A[i]
		}
	}*/
	for number, value := range m {
		if value%2 == 1 {
			return number
		}
	}
	return 0
}

func main() {
	A := []int{5, 5, 6, 7, 6, 6, 6}
	fmt.Println(":", Solution(A))
}
