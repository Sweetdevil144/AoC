package Arrays

import (
	"fmt"
	"math"
)

func ImageSmoother() {
	fmt.Println(imageSmoother([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
	fmt.Println(imageSmoother([][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}))
}

func imageSmoother(img [][]int) [][]int {
	if len(img) == 0 {
		return img
	}
	rows := len(img)
	cols := len(img[0])
	res := make([][]int, rows)
	for i := 0; i < rows; i++ {
		res[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res[i][j] = getAverage(img, i, j)
		}
	}
	return res
}

func getAverage(arr [][]int, i int, j int) int {
	sum := 0
	var (
		a     = 0
		b     = 0
		c     = 0
		d     = 0
		e     = 0
		f     = 0
		g     = 0
		h     = 0
		k     = 0
		count = 1
	)
	if i-1 >= 0 && j-1 >= 0 {
		a = arr[i-1][j-1]
		count++
	}
	if i-1 >= 0 {
		b = arr[i-1][j]
		count++
	}
	if i-1 >= 0 && j+1 < len(arr[0]) {
		c = arr[i-1][j+1]
		count++
	}
	if j-1 >= 0 {
		d = arr[i][j-1]
		count++
	}
	if j+1 < len(arr[0]) {
		e = arr[i][j+1]
		count++
	}
	if i+1 < len(arr) && j-1 >= 0 {
		f = arr[i+1][j-1]
		count++
	}
	if i+1 < len(arr) {
		g = arr[i+1][j]
		count++
	}
	if i+1 < len(arr) && j+1 < len(arr[0]) {
		h = arr[i+1][j+1]
		count++
	}
	k = arr[i][j]
	sum = (a + b + c + d + e + f + g + h + k) / count

	return int(math.Floor(float64(sum)))
}
