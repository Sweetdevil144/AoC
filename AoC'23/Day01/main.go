package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	a    int
	b    int
	sum1 int
	sum2 int
)

func main() {
	// Open the file.
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sum1 += partOne(scanner.Text())
		sum2 += partTwo(scanner.Text())
	}
	fmt.Println("Sum-1 is : ", sum1)
	fmt.Println("Sum-2 is : ", sum2)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func partOne(line string) int {
	for _, i2 := range line {
		// Find range [49,57]
		if i2 <= 57 && i2 >= 49 {
			a = int(i2 - 48)
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		ch := line[i]
		if ch <= 57 && ch >= 49 {
			b = int(ch - 48)
			break
		}
	}
	return a*10 + b
}

func partTwo(line string) int {
	// Declare a keyword map
	list := make(map[string]int)
	list["one"] = 1
	list["two"] = 2
	list["three"] = 3
	list["four"] = 4
	list["five"] = 5
	list["six"] = 6
	list["seven"] = 7
	list["eight"] = 8
	list["nine"] = 9
	str := replaceKeywords(line, list)
	return partOne(str)
}

func replaceKeywords(s string, list map[string]int) string {
	for key, value := range list {
		if strings.Contains(s, key) {
			s = strings.Replace(s, key, key+strconv.Itoa(value)+key, -1)
		}
	}
	return s
}
