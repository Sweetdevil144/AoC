package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	var res []int
	for i := 0; i < k; i++ {
		max := 0
		maxKey := 0
		for k, v := range m {
			if v > max {
				max = v
				maxKey = k
			}
		}
		res = append(res, maxKey)
		delete(m, maxKey)
	}
	return res
}
