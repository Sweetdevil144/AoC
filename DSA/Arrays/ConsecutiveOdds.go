package Arrays

func threeConsecutiveOdds(arr []int) bool {
	x := 0
	for _, element := range arr {
		if x == 3 {
			return true
		}
		if element%2 == 1 {
			x += 1
		} else {
			x = 0
		}
	}
	return x==3
}
