func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}
	freq := make([][]int, len(nums)+1)
	output := []int{}
	
	for _, num := range nums {
		count[num]++
	}

	for x, c := range count {
		freq[c] = append(freq[c], x)
	}

	for i := len(freq) - 1; i > 0; i-- {
		for _, num := range freq[i] {
			output = append(output, num)
			if len(output) == k {
				return output
			}
		}
	}
	return output
}