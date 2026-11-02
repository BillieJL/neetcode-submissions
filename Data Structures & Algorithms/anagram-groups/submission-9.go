func groupAnagrams(strs []string) [][]string {
	word := map[[26]int][]string{}

	for _, s := range strs {
		count := [26]int{}
		for _, c := range s {
			count[c-'a']++
		}
		word[count] = append(word[count], s)
	}
	
	result := [][]string{}
	for _, group := range word {
		result = append(result, group)
	}
	return result
}
