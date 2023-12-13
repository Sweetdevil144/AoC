package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(repeatedSubstringPattern("abac"))
	//fmt.Println(repeatedSubstringPattern("aba"))
	//fmt.Println(repeatedSubstringPattern("ababab"))
	fmt.Println(repeatedSubstringPattern("babbabbabbabbab"))
}

func repeatedSubstringPattern(s string) bool {
	if len(s) == 1 {
		return false
	}
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i == 0 {
			s1 := s[0:i]
			str := strings.Repeat(s1, len(s)/i)
			if str == s {
				return true
			}
		}
	}
	return false
}
