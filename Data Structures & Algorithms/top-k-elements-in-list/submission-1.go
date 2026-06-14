func topKFrequent(nums []int, k int) []int {
    count := map[int]int{}

	for _, n := range nums {
		count[n]++
	}

	buckets := make([][]int, len(nums)+1)

	for key, freq := range count {
		buckets[freq] = append(buckets[freq], key)
	}

	result := []int{}

	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		result = append(result, buckets[i]...)
	}

	return result
}
