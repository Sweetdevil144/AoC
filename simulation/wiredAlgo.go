package simulation

import "fmt"
func wierdAlgo() {
	var n int
	fmt.Scan(&n)
	fmt.Print(n)
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		fmt.Print(" ", n)
	}
	fmt.Println()
}