package day05

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day05() {
	firstSum := 0
	secondSum := 0
	rules, array := readData("AoC_24/Day05/input.txt")
	// fmt.Println("Rules :\t", rules)
	update := make([]bool, len(array))
	for i := range update {
		update[i] = true
	}
	for _, value := range rules {
		i1, i2 := value[0], value[1]
		for i, data := range array {
			if contains(data, i1) && contains(data, i2) {
				if update[i] && !liesAfter(data, i1, i2) {
					update[i] = false
				}
			}
		}
	}
	for i, value := range update {
		if value {
			middleIndex := len(array[i]) / 2
			firstSum += array[i][middleIndex]
		} else {
			fmt.Println("Rules :\t", rules)
			fmt.Println("Before Rearrangement :\t", array[i])
			array[i] = fixOrder(array[i], rules)
			fmt.Println("After Rearrangement :\t", array[i])
		}
	}
	for i, value := range update {
		if !value {
			fmt.Println(array[i])
			middleIndex := len(array[i]) / 2
			secondSum += array[i][middleIndex]
		}
	}
	fmt.Println(firstSum)
	fmt.Println(secondSum)
}

func fixOrder(array []int, rules [][]int) []int {
	for {
		changed := false
		for _, rule := range rules {
			i1, i2 := rule[0], rule[1]
			if contains(array, i1) && contains(array, i2) {
				if !liesAfter(array, i1, i2) {
					array = rearrageArray(array, i1, i2)
					changed = true
				}
			}
		}
		if !changed {
			break
		}
	}
	return array
}

func rearrageArray(data []int, i1, i2 int) []int {
	var i1Indices, i2Indices []int
	for i, v := range data {
		if v == i1 {
			i1Indices = append(i1Indices, i)
		}
		if v == i2 {
			i2Indices = append(i2Indices, i)
		}
	}
	for _, i1Index := range i1Indices {
		for _, i2Index := range i2Indices {
			if i1Index > i2Index {
				data[i1Index], data[i2Index] = data[i2Index], data[i1Index]
			}
		}
	}
	return data
}

func liesAfter(data []int, i1, i2 int) bool {
	index1, index2 := -1, -1

	for i, v := range data {
		if v == i1 && index1 == -1 {
			index1 = i
		}
		if v == i2 && index2 == -1 {
			index2 = i
		}
		if index1 != -1 && index2 != -1 {
			break
		}
	}

	return index1 != -1 && index2 != -1 && index1 < index2
}

func contains(arr []int, elem int) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}

func readData(path string) ([][]int, [][]int) {
	case1 := true
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error Reading File: %v", err)
	}
	defer file.Close()

	var arrayData [][]int
	var singleData [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && case1 {
			case1 = false
			continue
		}
		if case1 {
			parts := strings.Split(line, "|")
			var pair []int
			for _, part := range parts {
				num, err := strconv.Atoi(strings.TrimSpace(part))
				if err == nil {
					pair = append(pair, num)
				}
			}
			if len(pair) == 2 {
				arrayData = append(arrayData, pair)
			} else {
				log.Fatal("Len != 2")
			}
		} else {
			parts := strings.Split(line, ",")
			var lineData []int
			for _, value := range parts {
				num, err := strconv.Atoi(strings.TrimSpace(value))
				if err == nil {
					lineData = append(lineData, num)
				} else {
					fmt.Println(num, "\n", value)
					log.Fatal("Error occurred while parsing S -> Z")
				}
			}
			singleData = append(singleData, lineData)
		}
	}
	return arrayData, singleData
}
