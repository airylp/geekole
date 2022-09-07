package main

func longestPalindrome(s string) string {
	palInit := 0
	PalEnd := 0
	l := 0
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if len(s)-i+1 > l {
				for j := len(s) - 1; j > i; j-- {
					lAux := j - i + 1
					if lAux > l {
						isPal := true
						for k := 0; k < (lAux)/2; k++ {
							if s[i+k] != s[j-k] {
								isPal = false
								break
							}
						}
						if isPal == true {
							l = lAux
							palInit = i
							PalEnd = j
						}
					} else {
						break
					}
				}
				if l == len(s) {
					return s
				}
			}
		}
		return s[palInit : PalEnd+1]
	} else {
		return s
	}
}

func main() {
	//fmt.Println("res: ", longestPalindrome("12345654321abaaba"))
}
