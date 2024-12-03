package Strings

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagramsHelper() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(anagramsOptimised([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

// Brute Force -> O(n^2 * log(n))
func groupAnagrams(strings []string) [][]string {
	arr := make([][]string, len(strings))
	x := 0
	for i := 0; i < len(strings); {
		str := strings[i]
		for j := i + 1; j < len(strings); {
			if sortStr(str) == sortStr(strings[j]) {
				arr[x] = append(arr[x], strings[j])
				strings = append(strings[:j], strings[j+1:]...)
			} else {
				j++
			}
		}
		arr[x] = append(arr[x], str)
		strings = append(strings[:i], strings[i+1:]...)
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

// Optimized -> O(n^2)
func anagramsOptimised(strs []string) [][]string {
	stringMap := make(map[string][]string)

	for _, s := range strs {
		c := strings.Split(s, "")
		sort.Strings(c)
		hashString := strings.Join(c, "")

		stringMap[hashString] = append(stringMap[hashString], s)
	}

	var result [][]string
	for _, v := range stringMap {
		result = append(result, v)
	}

	return result
}
