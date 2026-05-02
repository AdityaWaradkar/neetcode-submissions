func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    counts := make(map[byte]int)

    for i := 0; i < len(s); i++ {
        counts[s[i]]++ 
    }

    for i := 0; i < len(t); i++ {
        if counts[t[i]] == 0 {
            return false 
        }
        counts[t[i]]-- 
	}
    return true
}