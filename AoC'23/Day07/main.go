package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("data.txt")
	//file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Fields(line)
		data = append(data, splitLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	cardPoints := map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 10,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
	}

	fmt.Println(partOne(data, cardPoints))
}

func partOne(data [][]string, points map[string]int) int64 {
	n := len(data)
	arr := make([]int, n)
	for i, element := range data {
		arr[i] = determineHand(element[0])
	}
	a1 := make([][]string, 0)
	a2 := make([][]string, 0)
	a3 := make([][]string, 0)
	a4 := make([][]string, 0)
	a5 := make([][]string, 0)
	a6 := make([][]string, 0)
	a7 := make([][]string, 0)
	for index, element := range arr {
		if element == 1 {
			a1 = append(a1, data[index])
		} else if element == 2 {
			a2 = append(a2, data[index])
		} else if element == 3 {
			a3 = append(a3, data[index])
		} else if element == 4 {
			a4 = append(a4, data[index])
		} else if element == 5 {
			a5 = append(a5, data[index])
		} else if element == 6 {
			a6 = append(a6, data[index])
		} else if element == 7 {
			a7 = append(a7, data[index])
		} else {
			fmt.Println("Error")
		}
	}
	//fmt.Println(a1, "\t", a2, "\t", a3, "\t", a4, "\t", a5, "\t", a6, "\t", a7)

	a1 = sortArray(a1, points)
	a2 = sortArray(a2, points)
	a3 = sortArray(a3, points)
	a4 = sortArray(a4, points)
	a5 = sortArray(a5, points)
	a6 = sortArray(a6, points)
	a7 = sortArray(a7, points)

	mainArr := make([][]string, 0)
	mainArr = append(mainArr, a7...)
	mainArr = append(mainArr, a6...)
	mainArr = append(mainArr, a5...)
	mainArr = append(mainArr, a4...)
	mainArr = append(mainArr, a3...)
	mainArr = append(mainArr, a2...)
	mainArr = append(mainArr, a1...)
	var sum int64 = 0
	for i := 0; i < len(mainArr); i++ {
		x, _ := strconv.Atoi(mainArr[i][1])
		sum += int64(x * (i + 1))
	}
	fmt.Println(mainArr)
	return sum
}

func sortArray(arr [][]string, points map[string]int) [][]string {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if findHigherPriority(points, arr[i][0], arr[j][0]) {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func findHigherPriority(points map[string]int, str1 string, str2 string) bool {
	for i := 0; i < 5; i++ {
		if points[string(str1[i])] > points[string(str2[i])] {
			return true
		} else if points[string(str1[i])] < points[string(str2[i])] {
			return false
		}
	}
	return false // If all cards are equal, return false
}

func determineHand(cards string) int {
	cardCounts := make(map[rune]int)
	for _, card := range cards {
		cardCounts[card]++
	}

	maxCount := 0
	for _, count := range cardCounts {
		if count > maxCount {
			maxCount = count
		}
	}

	switch {
	case len(cardCounts) == 1:
		// "Five of a Kind"
		return 1
	case maxCount == 4:
		//"Four of a Kind"
		return 2
	case maxCount == 3 && len(cardCounts) == 2:
		//"Full House"
		return 3
	case maxCount == 3:
		//"Three of a Kind"
		return 4
	case maxCount == 2 && len(cardCounts) == 3:
		//"Two Pairs"
		return 5
	case maxCount == 2:
		//"One Pair"
		return 6
	default:
		//"High Card"
		return 7
	}
}
