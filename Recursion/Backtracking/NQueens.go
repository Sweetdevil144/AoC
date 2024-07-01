package main

import "fmt"

func NQueens() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	s := make([][]string, n)
	for i := 0; i < n; i++ {
		s[i] = make([]string, n)
	}
	ans := make([][]string, 0)
	solveQueens(n, s, 0, &ans)
	return ans
}

func solveQueens(n int, str [][]string, i int, ans *[][]string) {
	if i == n {
		newSolution := transformX(str)
		*ans = append(*ans, newSolution)
		return
	}

	for j := 0; j < n; j++ {
		if isSafe(str, i, j) {
			str[i][j] = "Q"
			solveQueens(n, str, i+1, ans)
			str[i][j] = ""
		}
	}
}

func isSafe(str [][]string, row, col int) bool {
	for i := 0; i < row; i++ {
		if str[i][col] == "Q" {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if str[i][j] == "Q" {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j < len(str); i, j = i-1, j+1 {
		if str[i][j] == "Q" {
			return false
		}
	}
	return true
}

func transformX(x [][]string) []string {
	res := make([]string, len(x))
	for i := 0; i < len(x); i++ {
		str := ""
		for j := 0; j < len(x[i]); j++ {
			if x[i][j] == "Q" {
				str += "Q"
			} else {
				str += "."
			}
		}
		res[i] = str
	}
	return res
}
