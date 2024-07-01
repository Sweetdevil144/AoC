package main

import "fmt"

func horseshoe() {
	nums := make([]int, 4)
	for i := 0; i < 4; i++ {
		fmt.Scan(&nums[i])
	}

	nums = removeDuplicates(nums)
	fmt.Println(4 - len(nums))
}

func removeDuplicates(nums []int) []int {
	occurred := make(map[int]bool)
	result := []int{}

	for _, val := range nums {
		if _, ok := occurred[val]; !ok {
			occurred[val] = true
			result = append(result, val)
		}
	}

	return result
}
