package main

import "fmt"

func diffOneMinusZero() {
	fmt.Println(onesMinusZeros([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}))
	fmt.Println(onesMinusZeros([][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}}))
	fmt.Println(onesMinusZeros([][]int{{1, 1, 1}, {1, 1, 1}}))
}

func onesMinusZeros(grid [][]int) [][]int {
	rowSums := make([]int, len(grid))
	colSums := make([]int, len(grid[0]))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				rowSums[i]++
				colSums[j]++
			} else {
				rowSums[i]--
				colSums[j]--
			}
		}
	}

	diff := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		diff[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			diff[i][j] = rowSums[i] + colSums[j]
		}
	}
	return diff
}
