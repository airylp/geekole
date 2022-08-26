package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := make(map[int]int)

	l := 0
	previosStart := 0
	if len(s) > 0 {
		count := 0
		for count < len(s) {
			c := int(s[count])
			if firsOccurence, exist := m[c]; exist {
				if firsOccurence >= previosStart {
					previosStart = firsOccurence + 1
				}
			}
			m[c] = count

			if count+1-previosStart > l {
				l = count + 1 - previosStart
			}

			count++
		}

	}

	return l
}

func main() {
	fmt.Println("datos: ", lengthOfLongestSubstring("jbpnbwwd"))  //4
	fmt.Println("datos: ", lengthOfLongestSubstring("abcabcbb"))  //3
	fmt.Println("datos: ", lengthOfLongestSubstring("bb"))        //1
	fmt.Println("datos: ", lengthOfLongestSubstring("pwwkew"))    //3
	fmt.Println("datos: ", lengthOfLongestSubstring("anviaj"))    //5
	fmt.Println("datos: ", lengthOfLongestSubstring("umvejcuuk")) //6
	fmt.Println("datos: ", lengthOfLongestSubstring("tmmzuxt"))   //5

}
