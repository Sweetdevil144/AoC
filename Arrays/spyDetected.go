package main

import "fmt"

func spyDetected() {
	var (
		n int
		k int
	)
	fmt.Scan(&n)
	for n > 0 {
		fmt.Scan(&k)
		var arr = make([]int, k)
		for i := 0; i < k; i++ {
			fmt.Scan(&arr[i])
		}
		loopResult(arr)
		n--
	}
}

func loopResult(arr []int) {
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] != arr[i-1] && arr[i] != arr[i+1] {
			fmt.Println(i + 1)
			return
		}
	}
	if arr[0] != arr[1] {
		fmt.Println(1)
		return
	}
	fmt.Println(len(arr))
}
