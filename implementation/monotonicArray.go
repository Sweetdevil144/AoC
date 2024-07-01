package implementation


func isMonotonic(arr []int) bool {
	if len(arr) == 1 {
		return true
	}
	if isIncreasing(arr) || isDecreasing(arr) {
		return true
	} else {
		return false
	}
}

func isIncreasing(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func isDecreasing(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			return false
		}
	}
	return true
}
