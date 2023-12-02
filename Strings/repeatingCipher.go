package main

import "fmt"

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n)
	fmt.Scan(&s)
	fmt.Println(decipher(s, n))
}

func decipher(s string, n int) string {
	var res string
	for i := 0; i < n; i = 2*i + 1 {
		res += string(s[i])
	}
	return res
}
