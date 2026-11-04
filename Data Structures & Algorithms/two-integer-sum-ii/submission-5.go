func twoSum(numbers []int, target int) []int {
	seen := map[int]int{}
	for idx, num := range numbers {
		res := target - num
		if seen[res] > 0 {
			return []int{seen[res], (idx+1)}
		}
		seen[num] = idx + 1
	}
	return []int{}
}