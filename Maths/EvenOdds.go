package Maths

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func EvenOdds() {
	scanner := bufio.NewScanner(os.Stdin)
	str := ""
	if scanner.Scan() {
		str = scanner.Text()
	}
	arr := strings.Split(str, " ")
	n, _ := strconv.Atoi(arr[0])
	k, _ := strconv.Atoi(arr[1])
	if n%2 == 0 {
		if k <= n/2 {
			fmt.Println(k*2 - 1)
		} else {
			fmt.Println((k - n/2) * 2)
		}
	} else {
		if k <= n/2+1 {
			fmt.Println(k*2 - 1)
		} else {
			k = k - (n/2 + 1)
			fmt.Println(k * 2)
		}
	}

}
