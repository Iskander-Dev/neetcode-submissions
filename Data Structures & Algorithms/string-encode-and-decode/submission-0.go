type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder

	for _, str := range strs {
		fmt.Fprintf(&sb, "%d#%s", len(str), str)
	}

	return sb.String()
}

func (s *Solution) Decode(strs string) []string {
	decodedStrs := []string{}
	pos := 0

	for pos < len(strs) {
		j := pos

		for strs[j] != '#' {
			j++
		}

		length, err := strconv.Atoi(strs[pos:j])
		if err != nil {
			return decodedStrs
		}

		decodedStrs = append(decodedStrs, strs[j+1:j+length+1])
		pos = j + length + 1
	}

	return decodedStrs
}
