package implementation

func twoSum(numbers []int, target int) []int {
	var (
		i   = 0
		j   = len(numbers) - 1
		sum = 0
	)
	for i < j {
		sum = numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else {
			if sum > target {
				j--
			} else {
				i++
			}
		}
	}
	return []int{-1, -1}
}
