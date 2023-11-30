package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		s1 string
		s2 string
		s3 string
		s4 string
	)
	fmt.Scan(&s1, &s2, &s3)
	s4 = sortString(s1 + s2)
	s3 = sortString(s3)
	if len(s4) < len(s3) || len(s4) > len(s3) {
		fmt.Println("NO")
	} else {
		if s4 == s3 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
