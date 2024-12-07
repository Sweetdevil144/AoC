package day03

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parseFile(input string) []string {
	raw, err := os.ReadFile(input)
	if err != nil {
		log.Fatal("Error Reading File")
	}

	data := string(raw)
	arrayData := strings.Split(data, "\n")
	return arrayData
}

func readInteger(line string, pattern byte) (int, int, bool) {
	integer := 0
	for key, value := range line {
		if value == rune(pattern) {
			if key == 0 { // Edge case: Empty integer before delimiter
				return 0, 0, false
			}
			return integer, key + 1, true
		}
		if isDigit(value) {
			integer = integer*10 + int(value-'0')
		} else {
			// Stop parsing on invalid character
			return 0, 0, false
		}
	}
	return 0, 0, false
}

func isDigit(value rune) bool {
	return value >= '0' && value <= '9'
}

func SecondPuzzle() {
	match := "mul("
	enableMul := "do()"
	disableMul := "don't()"
	enabled := true
	res := 0
	arrayData := parseFile("AoC_24/Day03/input.txt")
	for _, line := range arrayData {
		for i := 0; i < len(line); {
			if i+4 > len(line) {
				break
				// Prevent index out of range
			}
			if i + 7 < len(line) && line[i:i+7] == disableMul {
				fmt.Println("Disabled now at i = ",i)
				enabled=false
			}
			if line[i:i+4] == enableMul {
				fmt.Println("Enabled now at i = ",i)
				enabled=true
			}
			if line[i:i+4] == match && enabled {
				i += 4
				// Move past "mul("
				first, updatedIndex, ok := readInteger(line[i:], ',')
				if !ok || updatedIndex >= len(line) {
					i += 4
					continue
				}
				i += updatedIndex
				second, nextIndex, ok := readInteger(line[i:], ')')
				fmt.Printf("%d * %d \n", first, second)
				if !ok || nextIndex >= len(line) {
					i += 4
					continue
				}
				i += nextIndex
				res += first * second
			} else {
				i++
			}
		}
	}
	fmt.Println("Result:", res)
}


func FirstPuzzle() {
	match := "mul("
	res := 0
	arrayData := parseFile("AoC_24/Day03/example.txt")
	for _, line := range arrayData {
		for i := 0; i < len(line); {
			if i+4 > len(line) {
				break // Prevent index out of range
			}
			if line[i:i+4] == match {
				i += 4 // Move past "mul("
				first, updatedIndex, ok := readInteger(line[i:], ',')
				if !ok || updatedIndex >= len(line) {
					i += 4
					continue
				}
				i += updatedIndex
				second, nextIndex, ok := readInteger(line[i:], ')')
				fmt.Printf("%d * %d \n", first, second)
				if !ok || nextIndex >= len(line) {
					i += 4
					continue
				}
				i += nextIndex
				res += first * second
			} else {
				i++
			}
		}
	}
	fmt.Println("Result:", res)
}