package main

import "fmt"

func Solution(X int, Y int, D int) int {

	distance := Y - X
	jumps := distance / D
	remainder := distance % D
	if remainder > 0 {
		jumps = jumps + 1
	}

	return jumps
}

func main() {
	fmt.Println(Solution(10, 85, 30))
}
