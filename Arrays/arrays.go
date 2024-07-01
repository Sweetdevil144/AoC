package main

import "fmt"

func arrays() {
	var (
		n int
	)
	// Taking user input length of Array
	fmt.Scan(&n)
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	a1 := arr[1 : n-1]
	// Slicing an array creates a slice. Slices can be manipulated as we want.
	for index, element := range arr {
		// First variable denotes index during for loop
		// Second variable denotes element at that particular index while running the loop
		fmt.Println(index, " -> ", element)
	}
	fmt.Println(arr)
	fmt.Println(a1)
}
