package day02

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"math"
)

func FirstPuzzle() {
	res := 0
	raw, err := os.ReadFile("AoC_24/Day02/input.txt")
	if err != nil {
		log.Fatal("Error Reading File")
	}

	data := string(raw)
	arrayData := strings.Split(data, "\n")
	for _, line := range arrayData {
		if line == "" {
			continue
		}
		dat := strings.Split(line, " ")
		var array = []float64{}

		for _, i := range dat {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			array = append(array, float64(j))
		}
		fmt.Println(array)
		if isSafe, _ := isSafeArray(array); isSafe {
			res++
		} else {
			// Check by removing each level
			if checkSafetyByRemovingOne(array) {
				res++
			}
		}
	}

	fmt.Println("Total Safe Reports:", res)
}

func isSafeArray(array []float64) (bool, int) {
	var dec bool = false
	if len(array) > 1 && array[0] > array[1] {
		dec = true
	}
	for i := 0; i < len(array)-1; i++ {
		if dec {
			if array[i] < array[i+1] {
				return false, i
			}
		} else {
			if array[i] > array[i+1] {
				return false, i
			}
		}
		diff := math.Abs(array[i] - array[i+1])
		if !(diff >= 1 && diff <= 3) {
			return false, i
		}
	}
	return true, -1
}

func checkSafetyByRemovingOne(array []float64) bool {
	for i := range array {
		newArray := make([]float64, 0, len(array)-1)
		newArray = append(newArray, array[:i]...)
		newArray = append(newArray, array[i+1:]...)
		if isSafe, _ := isSafeArray(newArray); isSafe {
			return true
		}
	}
	return false
}
