package main

import "fmt"

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
}

func findSpecialInteger(arr []int) int {
	n := len(arr)
	req := n / 4
	for i := 0; i < n-req; i++ {
		if arr[i] == arr[i+req] {
			return arr[i]
		}
	}
	return 0
}
