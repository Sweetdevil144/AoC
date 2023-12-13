package main

import "fmt"

func main() {
	fmt.Println(judgeCircle("UPUP"))
	fmt.Println(judgeCircle("UDD"))
}

func judgeCircle(moves string) bool {
	if len(moves)%2 == 0 {
		return false
	}
	var (
		cU = 0
		cD = 0
		cL = 0
		cR = 0
	)
	for _, move := range moves {
		if move == 'U' {
			cU += 1
		} else if move == 'D' {
			cD += 1
		} else if move == 'L' {
			cL += 1
		} else {
			cR += 1
		}
	}
	return cU-cD == 0 && cL-cR == 0
}
