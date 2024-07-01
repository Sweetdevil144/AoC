package Strings

import (
	"strconv"
)


func largestOddNumber(num string) string {
	s := ""
	for i := len(num) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(num[i]))
		if digit%2 == 1 {
			s = num[0 : i+1]
			break
		}
	}
	return s
}
