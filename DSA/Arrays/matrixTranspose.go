package Arrays

import "fmt"

func Transpose() {
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func transpose(matrix [][]int) [][]int {
	var (
		r = len(matrix)
		c = len(matrix[0])
	)
	arr := make([][]int, c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			arr[j] = append(arr[j], matrix[i][j])
		}
	}
	return arr
}
