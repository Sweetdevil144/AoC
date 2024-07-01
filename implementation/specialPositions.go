package main


func numSpecial(mat [][]int) int {
	c := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 && checkRow(i, mat) && checkCol(j, mat) {
				c++
			}
		}
	}
	return c
}

func checkRow(row int, arr [][]int) bool {
	c := 0
	for i := 0; i < len(arr[row]); i++ {
		if arr[row][i] == 1 {
			c++
		}
	}
	return c == 1
}

func checkCol(col int, arr [][]int) bool {
	c := 0
	for i := 0; i < len(arr); i++ {
		if arr[i][col] == 1 {
			c++
		}
	}
	return c == 1
}
