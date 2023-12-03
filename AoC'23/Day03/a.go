package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := readAndPadFile("data.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	sum := 0
	for i := 1; i < len(data)-1; i++ {
		for j := 1; j < len(data[i])-1; j++ {
			if isDigit(data[i][j]) {
				symbolCount := symbolInNeighbourhood(data, i, j)
				if symbolCount > 0 {
					numberStr := ""
					k := j
					for k < len(data[i]) && isDigit(data[i][k]) {
						numberStr += string(data[i][k])
						k++
					}
					number, _ := strconv.Atoi(numberStr)
					sum += number
					j = k - 1
				}
			}
		}
	}
	fmt.Println(sum)
}

func readAndPadFile(filePath string) ([][]rune, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data [][]rune
	data = append(data, make([]rune, 0)) // First row of padding

	for scanner.Scan() {
		line := []rune{'.'} // Start with padding
		line = append(line, []rune(scanner.Text())...)
		line = append(line, '.') // End with padding
		data = append(data, line)
	}
	padLength := len(data[1])
	data = append(data, make([]rune, padLength)) // Last row of padding

	for i := range data {
		if len(data[i]) < padLength {
			data[i] = append(data[i], make([]rune, padLength-len(data[i]))...)
		}
	}

	return data, nil
}

func symbolInNeighbourhood(data [][]rune, i, j int) int {
	symbols := 0
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			ni, nj := i+di, j+dj
			if ni >= 0 && ni < len(data) && nj >= 0 && nj < len(data[ni]) {
				if di == 0 && dj == 0 {
					continue
				}
				if data[ni][nj] != '.' && !isDigit(data[ni][nj]) {
					symbols++
				}
			}
		}
	}
	return symbols
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
