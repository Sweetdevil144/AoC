package main

import "fmt"

func main() {
	time := []int{49, 87, 78, 95}
	distance := []int{356, 1378, 1502, 1882}
	p := 1
	for i := 0; i < 4; i++ {
		p *= partOne(time[i], distance[i])
	}
	fmt.Println(p)
	fmt.Println(partOne(49877895, 356137815021882))
}

func partOne(time int, distance int) int {
	c := 0
	for i := 0; i <= time; i++ {
		x := i * (time - i)
		if x > distance {
			c++
		}
	}
	return c
}
