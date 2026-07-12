func twoSum(numbers []int, target int) []int {
	result := make([]int, 0, 2)
	n := len(numbers)

	left, right := 0, n-1

	for left < right {
		if numbers[left]+numbers[right] == target {
			result = append(result, left+1)
			result = append(result, right+1)
			break
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}

	return result
}
