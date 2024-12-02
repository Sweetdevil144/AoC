package day01

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
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
	left, right, err := parseInput("AoC_24/Day01/input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	slices.Sort(left)
	slices.Sort(right)

	diff := 0.0
	for i := 0; i < len(left); i++ {
		diff += math.Abs(right[i] - left[i])
	}
	fmt.Println(diff)
}

func SecondPuzzle() float64 {
	var sum float64 = 0
	left, right, err := parseInput("AoC_24/Day01/input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	leftMap := make(map[float64]float64)
	rightMap := make(map[float64]float64)

	for i := range left {
		leftMap[left[i]]++
		rightMap[right[i]]++
	}

	for key := range leftMap {
		if _, ok := rightMap[key]; ok {
			sum+=key * rightMap[key] * leftMap[key]
		}
	}
	fmt.Printf("%.0f\n", sum)
	return sum
}
