package main

import "fmt"

func CalculateResult() {
	var testCases, number int
	fmt.Scan(&testCases)
	for testCases > 0 {
		fmt.Scan(&number)
		l, r := calculateResult(number)
		fmt.Println(l, r)
		testCases--
	}
}

func calculateResult(num int) (int, int) {
	if num == 1 {
		return 0, 1
	} else if num == 2 {
		return -1, 2
	}

	factors := getFactors(num)
	for _, factor := range factors {
		otherFactor := num / factor
		lowerBound := (2*factor - otherFactor + 1) / 2
		upperBound := (2*factor + otherFactor - 1) / 2

		if lowerBound < upperBound && (lowerBound+upperBound)*(upperBound-lowerBound+1)/2 == num {
			return lowerBound, upperBound
		}
	}
	return -1, -1
}

func getFactors(num int) []int {
	var factors []int
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
			if otherFactor := num / i; otherFactor != i {
				factors = append(factors, otherFactor)
			}
		}
	}
	return factors
}
