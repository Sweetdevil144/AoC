package Arrays

import "fmt"

func FindAmazing() {
	var n int
	fmt.Scan(&n)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}
	fmt.Println(findAmazing(x))
}

func findAmazing(x []int) int32 {
	c := 0
	for i := 1; i < len(x); i++ {
		if x[i] > findMax(x[:i]) || x[i] < findMin(x[:i]) {
			c++
		}
	}
	return int32(c)
}

func findMax(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func findMin(nums []int) int {
	min := nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}
