func hasDuplicate(nums []int) bool {
    x := map[int]bool{}

	for i:=0; i<len(nums); i++ {
		x[nums[i]] = true;
	}
	return len(x) != len(nums)
}
