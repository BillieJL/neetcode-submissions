func isAnagram(s string, t string) bool {
    // If lengths are different, they cannot be anagrams
    if len(s) != len(t) {
        return false
    }

    var freq [26]int

    // Increment for s, decrement for t in a single loop
    for i := 0; i < len(s); i++ {
        freq[s[i]-'a']++
        freq[t[i]-'a']--
    }

    // If any value is not zero, the letter counts didn't match
    for _, count := range freq {
        if count != 0 {
            return false
        }
    }

    return true
}
