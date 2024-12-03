package Arrays

import "fmt"

func TotalMoney() {
	fmt.Println(totalMoney(4))
	fmt.Println(totalMoney(10))
	fmt.Println(totalMoney(20))
}

func totalMoney(n int) int {
	changedValue := 0
	sum := 0
	for i := 1; i <= n; i++ {
		if i%7 == 0 {
			sum += 7 + changedValue
			changedValue++
		} else {
			sum += i%7 + changedValue
		}
	}
	return sum
}
