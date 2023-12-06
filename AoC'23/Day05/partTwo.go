package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

type DataTwo struct {
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
	content, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}

	var data DataTwo
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	minLocation := make(chan int)
	minValue := int(^uint(0) >> 1) // Max int value

	for i := 0; i < len(data.Seeds); i += 2 {
		start, length := data.Seeds[i], data.Seeds[i]+data.Seeds[i+1]
		wg.Add(1)
		go func(start, length int) {
			defer wg.Done()
			localMin := processRange(start, length, data)
			minLocation <- localMin
		}(start, length)
	}

	// Collect results
	go func() {
		wg.Wait()
		close(minLocation)
	}()

	for loc := range minLocation {
		if loc < minValue {
			minValue = loc
		}
	}

	fmt.Println("Minimum location number:", minValue)
}

func processRange(start, length int, data DataTwo) int {
	resultArray := make([]int, 0)
	for seed := start; seed < start+length; seed++ {
		resultArray = append(resultArray, processSeed(seed, data))
	}
	return getMin(resultArray)
}

func processSeed(seed int, data DataTwo) int {
	var (
		num1 = mapValue(seed, data.SeedToSoilMap)
		num2 = mapValue(num1, data.SoilToFertilizerMap)
		num3 = mapValue(num2, data.FertilizerToWaterMap)
		num4 = mapValue(num3, data.WaterToLightMap)
		num5 = mapValue(num4, data.LightToTemperatureMap)
		num6 = mapValue(num5, data.TemperatureToHumidityMap)
		num7 = mapValue(num6, data.HumidityToLocationMap)
	)
	return num7
}

func mapValue(n int, mapping [][]int) int {
	for _, a := range mapping {
		if inRange(n, a[1], a[1]+a[2]) {
			return a[0] + n - a[1]
		}
	}
	return n
}

func inRange(n1, m1, m3 int) bool {
	return n1 >= m1 && n1 < m3
}

func getMin(arr []int) int {
	m := arr[0]
	for _, v := range arr {
		if v < m {
			m = v
		}
	}
	return m
}
