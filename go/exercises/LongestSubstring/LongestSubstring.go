package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := make(map[string]int, len(s))

	l := 0
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			m[string(s[i])] = i
			for j := i + 1; j < len(s); j++ {
				//fmt.Println(string(s[i]), string(s[j]))
				if _, exist := m[string(s[j])]; exist {
					//fmt.Println(string(s[j]))
					if len(m) > l {
						l = len(m)
					}
					for k := range m {
						delete(m, k)
					}
					break
				} else {
					m[string(s[j])] = j
				}
			}
		}
	}
	if len(m) > l {
		l = len(m)
	}

	return l
}

func main() {
	fmt.Println("datos: ", lengthOfLongestSubstring("jbpnbwwd"))
}
