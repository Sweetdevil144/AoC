package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type GameData struct {
	Red   int `json:"red"`
	Blue  int `json:"blue"`
	Green int `json:"green"`
}

var (
	value     bool
	max_red   int
	max_blue  int
	max_green int
)

func main() {
	var sum = 0

	// Define a variable of type map[string][]GameData
	var data map[string][]GameData
	// Open the JSON file
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new JSON decoder
	decoder := json.NewDecoder(file)

	// Decode the JSON data from the file into the map
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= 100; i++ {
		str := strconv.Itoa(i)
		sum += partTwo(data, str)
		if getResultOne(data, str) {
			sum += i
		}
	}
	fmt.Println(sum)
}

func partTwo(data map[string][]GameData, str string) int {
	max_red = 0
	max_blue = 0
	max_green = 0
	// Max maximum of red, green and blue in each row of data and return max_red*max_blue*max_green of each row.
	for j := 0; j < len(data[str]); j++ {
		if data[str][j].Red > max_red {
			max_red = data[str][j].Red
		}
		if data[str][j].Blue > max_blue {
			max_blue = data[str][j].Blue
		}
		if data[str][j].Green > max_green {
			max_green = data[str][j].Green
		}
	}
	return max_red * max_blue * max_green
}

func getResultOne(data map[string][]GameData, s string) bool {
	for j := 0; j < len(data[s]); j++ {
		value = getValueOne(data[s][j].Red, data[s][j].Blue, data[s][j].Green)
		if !value {
			return false
		}
	}
	return true
}

func getValueOne(red, blue, green int) bool {
	if red <= 12 && blue <= 14 && green <= 13 {
		return true
	} else {
		return false
	}
}
