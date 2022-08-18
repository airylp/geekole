package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	pos1 := 0
	pos2 := 0

	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2

	val1 := 0
	val2 := 0
	if l%2 == 1 {
		for i := 0; i <= int(l/2); i++ {

			if pos1 < l1 {
				if pos2 < l2 {
					if nums1[pos1] < nums2[pos2] {
						val1 = nums1[pos1]
						pos1++
					} else {
						val1 = nums2[pos2]
						pos2++
					}
				} else {
					val1 = nums1[pos1]
					pos1++
				}
			} else {
				if pos2 < l2 {
					val1 = nums2[pos2]
					pos2++
				}
			}
		}
		return float64(val1)
	} else {
		for i := 0; i <= int(l/2); i++ {
			val1 = val2
			if pos1 < l1 {
				if pos2 < l2 {
					if nums1[pos1] < nums2[pos2] {
						val2 = nums1[pos1]
						pos1++
					} else {
						val2 = nums2[pos2]
						pos2++
					}
				} else {
					val2 = nums1[pos1]
					pos1++
				}
			} else {
				if pos2 < l2 {
					val2 = nums2[pos2]
					pos2++
				}
			}
		}
		return float64(val2+val1) / 2.0
	}

}

func main() {
	fmt.Println("the median: ", findMedianSortedArrays([]int{1}, []int{100}))
}
