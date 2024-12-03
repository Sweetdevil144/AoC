package Arrays

import (
	"fmt"
	"math"
)

func hourseRobber1() {
	fmt.Println(rob([]int{1, 2, 3, 1}))                         // 4
	fmt.Println(rob([]int{2, 1, 1, 2}))                         // 4
	fmt.Println(rob([]int{1, 2, 2, 1}))                         // 3
	fmt.Println(rob([]int{6, 3, 10, 8, 2, 10, 3, 5, 10, 5, 3})) // 39
	fmt.Println(rob([]int{2, 2, 4, 3, 2, 5}))                   // 11
	fmt.Println(rob([]int{8, 9, 9, 4, 10, 5, 6, 9, 7, 9}))      // 45
}

func rob(nums []int) int {
	var (
		sum1 = 0
		sum2 = 0
		sum  = 0
		prev = 0
	)

	for i, num := range nums {
		if i%2 == 0 {
			sum1 += num
		} else {
			sum2 += num
		}
	}

	for i := 0; i < len(nums); i += 2 {
		if i == len(nums)-1 {
			if prev != i-1 {
				sum += nums[i]
				break
			}
		}
		if nums[i] >= nums[i+1] {
			sum += nums[i]
		} else if nums[i] == nums[i+1] {
			if prev != i-1 {
				sum += nums[i]
			}
		} else {
			sum += nums[i+1]
			i++
		}
		prev = i
	}
	fmt.Println(sum1, "\t", sum2, "\t", sum)
	return int(math.Max(float64(sum1), math.Max(float64(sum2), float64(sum))))
}
