package main

import (
	"fmt"
	"strings"
)

func MaxScore() {
	fmt.Println(maxScore("011101"))
	fmt.Println(maxScore("00111"))
}

func maxScore(s string) int {
	zeros := 0
	ones := strings.Count(s, "1")
	maxScore := 0

	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			zeros++
		} else {
			ones--
		}
		score := zeros + ones
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}
