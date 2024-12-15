package day06

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func FirstPuzzle() {
	path := "AoC_24/Day06/example.txt"
	data := readFile(path)
	for _, value := range data {
		fmt.Println(value)
	}
	i, j := findGuard(data)
	fmt.Println(i,j)
}

func findGuard(data [][]string) (i, j int) {
	for i := range data {
		for j := range data[i] {
			if data[i][j] == "^" {
				return i, j
			}
		}
	}
	return -1, -1
}

func readFile(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error Reading File: %v", err)
	}
	defer file.Close()

	var arrayData [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var pair []string
		parts := strings.Split(line, "")
		pair = append(pair, parts...)
		if len(pair) > 0 {
			arrayData = append(arrayData, pair)
		}
	}
	return arrayData
}
