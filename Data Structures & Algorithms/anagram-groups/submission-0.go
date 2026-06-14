func groupAnagrams(strs []string) [][]string {
    seen := make(map[[26]int][]string)

	for _, str := range strs {
		var key [26]int

		for _, ch := range str {
			key[ch-'a']++
		}

		seen[key] = append(seen[key], str)
	}

	result := make([][]string, 0, len(seen))

	for _, group := range seen {
		result = append(result, group)
	}

	return result
}
