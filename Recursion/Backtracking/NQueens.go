package main

import "fmt"

func main() {
	fmt.Println(NQueens(4))
}

func NQueens(n int) [][]string {
	s := make([][]string, 0)
	ans := make([][]string, 0)
	for i := 0; i < n; i++ {
		s[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		x, _ := solveQueens(n, s, 0)
		ans = append(s, transformX(x))
	}
	return ans
}

func transformX(x [][]string) []string {
	res := make([]string, 0)
	for i := 0; i < len(x); i++ {
		str := ""
		for j := 0; j < len(x[i]); j++ {
			str += x[i][j]
		}
		res = append(res, str)
	}
	return res
}

func solveQueens(n int, str [][]string, i int) ([][]string, bool) {
	comp := false
	if i == n {
		return str, true
	}
	for j := 0; j < len(str[i]); j++ {
		if str[i][j] != "." {
			str[i][j] = "Q"
			fmt.Println("Str of", i, j, "is", str)
			arr := changeRCD(str, i, j)
			arr, comp = solveQueens(n, arr, i+1)
			if comp {
				str = arr
			} else {
				str[i][j] = ""
			}
		}
	}
	return str, comp
}

func changeRCD(str [][]string, i int, j int) [][]string {
	// Create a deep copy of str
	copyStr := make([][]string, len(str))
	for i := range str {
		copyStr[i] = make([]string, len(str[i]))
		copy(copyStr[i], str[i])
	}

	n := len(copyStr)
	// Change all elements in the same row and column
	for k := 0; k < n; k++ {
		if copyStr[i][k] != "Q" {
			copyStr[i][k] = "."
		}
		if copyStr[k][j] != "Q" {
			copyStr[k][j] = "."
		}
	}

	// Change all elements in the same diagonals
	for k := 0; k < n; k++ {
		// Top-left to bottom-right diagonal
		if i+k < n && j+k < n && copyStr[i+k][j+k] != "Q" {
			copyStr[i+k][j+k] = "."
		}
		// Bottom-left to top-right diagonal
		if i-k >= 0 && j-k >= 0 && copyStr[i-k][j-k] != "Q" {
			copyStr[i-k][j-k] = "."
		}
		// Top-right to bottom-left diagonal
		if i+k < n && j-k >= 0 && copyStr[i+k][j-k] != "Q" {
			copyStr[i+k][j-k] = "."
		}
		// Bottom-right to top-left diagonal
		if i-k >= 0 && j+k < n && copyStr[i-k][j+k] != "Q" {
			copyStr[i-k][j+k] = "."
		}
	}
	return copyStr
}
