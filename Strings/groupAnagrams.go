package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{"", "", ""}))
}

func groupAnagrams(strs []string) [][]string {
	arr := make([][]string, len(strs))
	x := 0
	for i := 0; i < len(strs); {
		str := strs[i]
		for j := i + 1; j < len(strs); {
			if sortStr(str) == sortStr(strs[j]) {
				arr[x] = append(arr[x], strs[j])
				strs = append(strs[:j], strs[j+1:]...)
			} else {
				j++
			}
		}
		arr[x] = append(arr[x], str)
		strs = append(strs[:i], strs[i+1:]...)
		//fmt.Println("Current array = ", strs, " \t x =", x, " and \t arr =", arr)
		x++
	}
	arr = filterEmpty(arr)
	return arr
}

func filterEmpty(arr [][]string) [][]string {
	result := make([][]string, 0)
	for _, v := range arr {
		if len(v) != 0 {
			result = append(result, v)
		}
	}
	return result
}

func sortStr(s string) string {
	if len(s) == 0 {
		return ""
	}
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
