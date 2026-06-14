func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
		return 0
	}

	hashSet := map[int]struct{}{}

	for _, num := range nums {
		hashSet[num] = struct{}{}
	}

	maxCounter := 0
	counter := 1

	for _, num := range nums {
		_, exist := hashSet[num-1]

		if !exist {
			for true {
				if _, next := hashSet[num+counter]; next {
					counter++
				} else {
					if maxCounter < counter {
						maxCounter = counter
					}
					counter = 1
					break
				}
			}
		}
	}

	return maxCounter
}
