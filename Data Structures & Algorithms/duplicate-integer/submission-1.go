func hasDuplicate(nums []int) bool {
    seen := map[int]int{}

	for _, num := range nums {
		if seen[num] == 0 {
			seen[num] = 1
		} else {
			return true
		}
	}
	return false
}
