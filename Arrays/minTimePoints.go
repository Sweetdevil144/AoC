package main

import (
	"fmt"
	"math"
)

func MinTimeToVisitAllPoints() {
	fmt.Println(minTimeToVisitAllPoints([][]int{{1, 1}, {3, 4}, {-1, 0}}))
}

func minTimeToVisitAllPoints(points [][]int) int {
	sum := 0
	for i := 1; i < len(points); i++ {
		dx := points[i][0] - points[i-1][0]
		dy := points[i][1] - points[i-1][1]
		sum += int(math.Max(math.Abs(float64(dx)), math.Abs(float64(dy))))
	}
	return sum
}
