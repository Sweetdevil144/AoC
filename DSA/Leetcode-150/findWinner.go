package leetcode150

import "fmt"

func FindTheWinner(n int, k int) int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	x := k

	for len(arr) > 1 {
		fmt.Println(arr, "\t", x)
		arr = append(arr[0:x-1], arr[x:]...)
		x += k - 1
		l := len(arr)
		if x > l {
			x = x % l
		} else if x == 0 {
			x += 1
		}
	}
	fmt.Println(arr, "\t", x)

	return arr[0]
}
