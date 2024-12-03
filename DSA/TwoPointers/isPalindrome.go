package TwoPointers

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	var str = ""
	for _, char := range s {
		cond := (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')
		if cond {
			str += string(unicode.ToLower(char))
		}
	}
	var (
		start = 0
		end   = len(str) - 1
	)
	for start <= end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	fmt.Println(str)
	return true
}
