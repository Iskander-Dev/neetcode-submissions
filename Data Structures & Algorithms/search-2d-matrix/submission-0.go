func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := (right + left) / 2
		row := mid / n
		col := mid % n
		midValue := matrix[row][col]

		if midValue == target {
			return true
		}

		if midValue < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
