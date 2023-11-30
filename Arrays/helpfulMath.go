package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
	}
	if len(input) == 1 {
		fmt.Println(input)
		return
	}
	x := len(input)/2 + 1
	nums := make([]int, x)
	//fmt.Println(nums)
	var c = 0
	for i := 0; i < len(input); i += 2 {
		number := input[i]
		nums[c] = int(number - 48)
		c++
	}
	var str = ""
	sort.Ints(nums)
	//fmt.Println(nums)
	for i := 0; i < c-1; i++ {
		str = str + strconv.Itoa(nums[i]) + "+"
	}
	str = str + strconv.Itoa(nums[c-1])
	fmt.Println(str)
}
