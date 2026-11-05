func twoSum(numbers []int, target int) []int {
	seen := map[int]int{}
	var now int
	for key, n := range numbers {
		now = target - n
		if seen[now] > 0 {
			return []int{seen[now], key+1}
		}
		seen[n] = key+1
	}
	return []int{}
}