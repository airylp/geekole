package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	i := 0

	for i < len(nums) {
		j := i + 1
		for j < len(nums) {

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
			j++
		}
		i++
	}
	return []int{}
}

func main() {
	fmt.Println(TwoSum([]int{-3, 4, 3, 90}, 0))
}
