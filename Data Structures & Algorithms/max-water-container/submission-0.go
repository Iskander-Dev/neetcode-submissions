func maxArea(heights []int) int {
	maxArea := 0
	l, r := 0, len(heights)-1

	for l < r {
		minHeight := min(heights[l], heights[r])
		width := r - l

		maxArea = max(maxArea, minHeight*width)

		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea
}
