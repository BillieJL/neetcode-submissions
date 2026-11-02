func twoSum(nums []int, target int) []int {
	prev := map[int]int{}

	for i, num := range nums {
		x := target - num
		if j, found := prev[x]; found {
			return []int{j,i}
		}
		prev[num] = i
	}
	return []int{}
}