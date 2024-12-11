package day04

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FirstPuzzle() {
	file := "AoC_24/Day04/example.txt"
	array, err := readFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	occurrences := findOccurrences(array, "XMAS")
	fmt.Println("Total occurrences of XMAS:", occurrences)
}

func SecondPuzzle() {
	file := "AoC_24/Day04/input.txt"
	array, err := readFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	occurrences := findXMASPattern(array)
	fmt.Println("Total occurrences of X-MAS pattern:", occurrences)
}

func findXMASPattern(array [][]string) int {
	count := 0
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if isXMASAt(array, i, j, 1) || isXMASAt(array, i, j, 2) || isXMASAt(array, i, j, 3) || isXMASAt(array, i, j, 4) {
				count++
			}			
		}
	}
	return count
}

func isXMASAt(array [][]string, x, y int, caseType int) bool {
	switch caseType {
	case 1:
		// Case 1: M M
		//          A
		//         S S
		return isValid(array, x-1, y-1) && array[x-1][y-1] == "M" &&
			isValid(array, x-1, y+1) && array[x-1][y+1] == "M" &&
			isValid(array, x, y) && array[x][y] == "A" &&
			isValid(array, x+1, y-1) && array[x+1][y-1] == "S" &&
			isValid(array, x+1, y+1) && array[x+1][y+1] == "S"
	case 2:
		// Case 2: M S
		//          A
		//         M S
		return isValid(array, x-1, y-1) && array[x-1][y-1] == "M" &&
			isValid(array, x-1, y+1) && array[x-1][y+1] == "S" &&
			isValid(array, x, y) && array[x][y] == "A" &&
			isValid(array, x+1, y-1) && array[x+1][y-1] == "M" &&
			isValid(array, x+1, y+1) && array[x+1][y+1] == "S"
	case 3:
		// Case 3: S M
		//          A
		//         S M
		return isValid(array, x-1, y-1) && array[x-1][y-1] == "S" &&
			isValid(array, x-1, y+1) && array[x-1][y+1] == "M" &&
			isValid(array, x, y) && array[x][y] == "A" &&
			isValid(array, x+1, y-1) && array[x+1][y-1] == "S" &&
			isValid(array, x+1, y+1) && array[x+1][y+1] == "M"
	case 4:
		// Case 4: S S
		//          A
		//         M M
		return isValid(array, x-1, y-1) && array[x-1][y-1] == "S" &&
			isValid(array, x-1, y+1) && array[x-1][y+1] == "S" &&
			isValid(array, x, y) && array[x][y] == "A" &&
			isValid(array, x+1, y-1) && array[x+1][y-1] == "M" &&
			isValid(array, x+1, y+1) && array[x+1][y+1] == "M"
	}
	return false
}
func isValid(array [][]string, x, y int) bool {
    return x >= 0 && y >= 0 && x < len(array) && y < len(array[0])
}


func readFile(input string) ([][]string, error) {
	var result [][]string
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		result = append(result, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func findOccurrences(array [][]string, word string) int {
	count := 0
	wordLen := len(word)
	directions := [][]int{{0, 1}, {1, 0}, {1, 1}, {1, -1}, {0, -1}, {-1, 0}, {-1, -1}, {-1, 1}}
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			for _, dir := range directions {
				if isWordAt(array, word, i, j, dir, wordLen) {
					count++
				}
			}
		}
	}
	return count
}

func isWordAt(array [][]string, word string, x, y int, dir []int, wordLen int) bool {
	for k := 0; k < wordLen; k++ {
		nx, ny := x+k*dir[0], y+k*dir[1]
		if nx < 0 || ny < 0 || nx >= len(array) || ny >= len(array[0]) || array[nx][ny] != string(word[k]) {
			return false
		}
	}
	return true
}
