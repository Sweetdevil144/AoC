package Arrays

import (
	"fmt"
	"sort"
)

func LongestConsecutive() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	max := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] == nums[i-1]+1 {
			count++
		} else {
			count = 1
		}
		if count > max {
			max = count
		}
	}
	return max
}
