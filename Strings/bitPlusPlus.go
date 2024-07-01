package Strings

import "fmt"

func bitPlusPlus() {
	var (
		n int
		s string
		x int
	)
	x = 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		if s[1] == '+' {
			x++
		} else {
			x--
		}
	}
	fmt.Println(x)
}
