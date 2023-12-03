package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Int int
type Char string

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data [][]string
	for scanner.Scan() {
		line := scanner.Text()
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		data = append(data, row)
	}

	sum := 0
	for i, line := range data {
		for j, char := range line {
			s := Char(char)
			if s.isDigit() {
				if isAdjacentToSymbol(i, j, data) {
					number, _ := strconv.Atoi(string(s))
					sum += number
				}
			}
		}
	}
	fmt.Println(sum)
}

func isAdjacentToSymbol(i, j int, data [][]string) bool {
	directions := []struct{ dx, dy int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}
	for _, dir := range directions {
		ni, nj := i+dir.dx, j+dir.dy
		if ni >= 0 && ni < len(data) && nj >= 0 && nj < len(data[0]) {
			if Char(data[ni][nj]).isSymbol() {
				return true
			}
		}
	}
	return false
}

func (n Int) isDigit() bool {
	return n <= 9 && n >= 0
}

func (c Char) isDigit() bool {
	return c >= "0" && c <= "9"
}

func (c Char) isSymbol() bool {
	return c != "." && !c.isDigit()
}
