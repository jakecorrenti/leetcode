package easy

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sCount := map[byte]int{}
	tCount := map[byte]int{}

	for i := 0; i < len(s); i++ {
		sCount[s[i]] += 1
		tCount[t[i]] += 1
	}

	for k, v := range sCount {
		if v != tCount[k] {
			return false
		}
	}
	return true
}
