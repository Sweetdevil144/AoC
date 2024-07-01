package Strings

import (
	"fmt"
	"strconv"
)

func boyOrGirl() {
	var str string
	fmt.Scan(&str)

	// Create a map with string keys and boolean values
	set := make(map[string]bool)

	for i := 0; i < len(str); i++ {
		set[strconv.Itoa(int(str[i]))] = true
	}
	if len(set)%2 == 1 {
		fmt.Println("IGNORE HIM!")
	} else {
		fmt.Println("CHAT WITH HER!")
	}
}
