func characterReplacement(s string, k int) int {
	maxSub, l, n := 0, 0, len(s)
	seen := make(map[byte]int)

	for r := range n {
		seen[s[r]]++

		for (r-l+1)-maxVal(seen) > k {
			seen[s[l]]--
			l++
		}

		maxSub = max(maxSub, r-l+1)
	}
	return maxSub
}

func maxVal(m map[byte]int) int {
	bestVal := 0

	for _, v := range m {
		if v > bestVal {
			bestVal = v
		}
	}
	return bestVal
}