package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		n, k, l, c, d, p, nl, np int
	)
	fmt.Scan(&n, &k, &l, &c, &d, &p, &nl, &np)

	toastsFromDrink := (k * l) / nl
	toastsFromLime := c * d
	toastsFromSalt := p / np

	minToasts := min(toastsFromDrink, toastsFromLime, toastsFromSalt)

	fmt.Println(minToasts / n)
}

func min(a, b, c int) int {
	return int(math.Min(float64(a), math.Min(float64(b), float64(c))))
}
