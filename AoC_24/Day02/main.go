package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filePath string) ([]float64, []float64, error) {
	raw, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	data := string(raw)
	arrayData := strings.Split(data, "\n")
	left := make([]float64, len(arrayData))
	right := make([]float64, len(arrayData))

	for i := 0; i < len(arrayData); i++ {
		split := strings.Split(arrayData[i], "   ")
		leftVal, _ := strconv.ParseFloat(split[0], 64)
		rightVal, _ := strconv.ParseFloat(split[1], 64)
		left[i] = leftVal
		right[i] = rightVal
	}

	return left, right, nil
}
func FirstPuzzle() {

}