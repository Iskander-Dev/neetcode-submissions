func lengthOfLongestSubstring(s string) int {
	maxSubStr := 0
	l, r := 0, 0
	seen := [127]bool{}

	for r < len(s) {
		if seen[s[r]-' '] {
			maxSubStr = max(maxSubStr, r-l)
			for seen[s[r]-' '] {
				seen[s[l]-' '] = false
				l++
			}
		}
		seen[s[r]-' '] = true
		r++
	}
	maxSubStr = max(maxSubStr, r-l)
	return maxSubStr
}
