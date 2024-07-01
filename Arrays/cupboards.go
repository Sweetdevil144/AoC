package Arrays

import (
	"fmt"
)

func copboards() {
	var n int
	fmt.Scan(&n)

	// Counters for 0s and 1s in both columns
	countZeroInZero, countOneInZero := 0, 0
	countZeroInOne, countOneInOne := 0, 0

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		if a == 0 {
			countZeroInZero++
		} else {
			countOneInZero++
		}
		if b == 0 {
			countZeroInOne++
		} else {
			countOneInOne++
		}
	}

	minMoves := min(countZeroInZero, countOneInZero) + min(countZeroInOne, countOneInOne)
	fmt.Println(minMoves)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
