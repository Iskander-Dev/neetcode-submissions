func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}

	freq := map[byte]int{}

    for i := range s {
		freq[s[i]]++
		freq[t[i]]--
    }

	for _, v := range freq {
		if v != 0 {
			return false
		}
	}

	return true
}
