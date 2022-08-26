package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := make(map[int]int, len(s))

	l := 0
	previosStart := 0
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			c := int(s[i])

			//fmt.Println(string(s[i]))
			if firsOccurence, exist := m[c]; exist {
				//fmt.Println("exist", i, "l", l, "len(m)")
				if len(m) > l {
					l = len(m)
					//fmt.Println("l", l)
				}
				//fmt.Println("firsOccurence:", firsOccurence)
				//fmt.Println("previosStart:", previosStart)
				for j := previosStart; j <= firsOccurence; j++ {
					cAux := int(s[j])
					//fmt.Println("deleting:", j)
					delete(m, cAux)
				}
				previosStart = firsOccurence + 1
				m[c] = i

			} else {
				m[c] = i
			}
		}
	}
	if len(m) > l {
		l = len(m)
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

}
