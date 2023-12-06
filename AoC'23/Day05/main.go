package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Data struct {
	Seeds                    []int   `json:"seeds"`
	SeedToSoilMap            [][]int `json:"seed-to-soil map"`
	SoilToFertilizerMap      [][]int `json:"soil-to-fertilizer map"`
	FertilizerToWaterMap     [][]int `json:"fertilizer-to-water map"`
	WaterToLightMap          [][]int `json:"water-to-light map"`
	LightToTemperatureMap    [][]int `json:"light-to-temperature map"`
	TemperatureToHumidityMap [][]int `json:"temperature-to-humidity map"`
	HumidityToLocationMap    [][]int `json:"humidity-to-location map"`
}

func main() {
	content, err := os.ReadFile("example.json")
	//content, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}

	var data Data
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Seeds:", data.Seeds)
	//fmt.Println("Seed-to-soil map:", data.SeedToSoilMap)
	//fmt.Println("Soil-to-fertilizer map:", data.SoilToFertilizerMap)
	//fmt.Println("Fertilizer-to-water map:", data.FertilizerToWaterMap)
	//fmt.Println("Water-to-light map:", data.WaterToLightMap)
	//fmt.Println("Light-to-temperature map:", data.LightToTemperatureMap)
	//fmt.Println("Temperature-to-humidity map:", data.TemperatureToHumidityMap)
	//fmt.Println("Humidity-to-location map:", data.HumidityToLocationMap)
	seedsArray := make([]int, 0)
	for i := 0; i < len(data.Seeds)-1; i += 2 {
		for j := 0; j < data.Seeds[i+1]; j++ {
			seedsArray = append(seedsArray, data.Seeds[i]+j)
		}
	}
	data.Seeds = seedsArray
	fmt.Println(partOne(data))

}

func partOne(data Data) int {
	resultArray := make([]int, 0)
	for _, seed := range data.Seeds {

		var (
			num1 = 0
			num2 = 0
			num3 = 0
			num4 = 0
			num5 = 0
			num6 = 0
			num7 = 0
		)

		// 1st Mapping
		for _, a := range data.SeedToSoilMap {
			if isInRange(seed, a[1], a[1]+a[2]) {
				num1 = a[0] + seed - a[1]
			}
		}
		if num1 == 0 {
			num1 = seed
		}

		// 2nd Mapping
		for _, a := range data.SoilToFertilizerMap {
			if isInRange(num1, a[1], a[1]+a[2]) {
				num2 = a[0] + num1 - a[1]
			}
		}
		if num2 == 0 {
			num2 = num1
		}

		// 3rd Mapping
		for _, a := range data.FertilizerToWaterMap {
			if isInRange(num2, a[1], a[1]+a[2]) {
				num3 = a[0] + num2 - a[1]
			}
		}
		if num3 == 0 {
			num3 = num2
		}

		// 4th Mapping
		for _, a := range data.WaterToLightMap {
			if isInRange(num3, a[1], a[1]+a[2]) {
				num4 = a[0] + num3 - a[1]
			}
		}
		if num4 == 0 {
			num4 = num3
		}

		// 5th Mapping
		for _, a := range data.LightToTemperatureMap {
			if isInRange(num4, a[1], a[1]+a[2]) {
				num5 = a[0] + num4 - a[1]
			}
		}
		if num5 == 0 {
			num5 = num4
		}

		// 6th Mapping
		for _, a := range data.TemperatureToHumidityMap {
			if isInRange(num5, a[1], a[1]+a[2]) {
				num6 = a[0] + num5 - a[1]
			}
		}
		if num6 == 0 {
			num6 = num5
		}

		// 7th Mapping
		for _, a := range data.HumidityToLocationMap {
			if isInRange(num6, a[1], a[1]+a[2]) {
				num7 = a[0] + num6 - a[1]
			}
		}
		if num7 == 0 {
			num7 = num6
		}
		resultArray = append(resultArray, num7)
	}
	return findMin(resultArray)
}

func isInRange(n1, m1, m3 int) bool {
	return n1 >= m1 && n1 < m3
}

func findMin(arr []int) int {
	m := arr[0]
	for _, v := range arr {
		if v < m {
			m = v
		}
	}
	return m
}
