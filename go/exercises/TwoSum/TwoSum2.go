package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums)-1)
	for i := 0; i < len(nums); i++ {

		if val, ok := m[target-nums[i]]; ok {
			return []int{i, val}
		}
		m[nums[i]] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
