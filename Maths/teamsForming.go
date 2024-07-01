package Math

import (
	"fmt"
	"sort"
)

func TeamsForming() {
	var n int
	fmt.Scan(&n)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	if n == 1 {
		fmt.Println(0)
	} else {
		fmt.Println(teamsForming(arr))
	}
}

func teamsForming(arr []int) int {
	var res int
	sort.Ints(arr)
	for i := 0; i < len(arr); i += 2 {
		res += arr[i+1] - arr[i]
	}
	return res
}
