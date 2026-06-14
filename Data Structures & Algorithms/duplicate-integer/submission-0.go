func hasDuplicate(nums []int) bool {
    seen := map[int]struct{}{}

    for _, n := range nums {
        if _, ok := seen[n]; ok {
            return true
        }
        seen[n] = struct{}{}
    }

    return false
}
