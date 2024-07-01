package main

import "fmt"

func CountMatchingSubarrays() {
	fmt.Println(countMatchingSubarrays([]int{1, 2, 3, 4, 5, 6}, []int{1, 1}))
	fmt.Println(countMatchingSubarrays([]int{1, 4, 4, 1, 3, 5, 5, 3}, []int{-1, 0, 1}))
}

func countMatchingSubarrays(nums []int, pattern []int) int {
	count := 0
	l := len(pattern)
	for i := 0; i < len(nums)-l; i++ {
		if compare(nums[i:i+l+1], pattern) {
			count++
		}
	}
	return count
}

func compare(a []int, p []int) bool {
	for i := 0; i < len(p); i++ {
		if p[i] == 1 && a[i+1] <= a[i] {
			return false
		} else if p[i] == 0 && a[i+1] != a[i] {
			return false
		} else if p[i] == -1 && a[i+1] >= a[i] {
			return false
		}
	}
	return true
}
