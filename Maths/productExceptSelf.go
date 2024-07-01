package main

import "fmt"

func ProductExceptSelf() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

func productExceptSelf(nums []int) []int {
	p := 1
	n := len(nums)
	for _, num := range nums {
		p *= num
	}
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			arr[i] = getProduct(nums, i)
		} else {
			arr[i] = p / nums[i]
		}
	}
	return arr
}

func getProduct(nums []int, i int) int {
	p := 1
	for j, num := range nums {
		if j != i {
			p *= num
		}
	}
	return p
}
