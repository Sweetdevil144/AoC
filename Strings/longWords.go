package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		n   int
		str string
	)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&str)
		if len(str) > 10 {
			str = getString(str)
		}
		fmt.Println(str)
	}

}

func getString(str string) string {
	length := len(str) - 2
	c1 := str[0]
	c2 := str[length+1]
	return string(c1) + strconv.Itoa(length) + string(c2)
}
