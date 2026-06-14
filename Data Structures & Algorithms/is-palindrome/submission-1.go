func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1

	for left <= right {
		for left <= right && !unicode.IsDigit(rune(s[left])) && !unicode.IsLetter(rune(s[left])) {
			left++
		}
		for left <= right && !unicode.IsDigit(rune(s[right])) && !unicode.IsLetter(rune(s[right])) {
			right--
		}
		if left > right {
			break
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
