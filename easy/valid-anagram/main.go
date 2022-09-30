package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dict := make(map[rune]int)
	dict2 := make(map[rune]int)

	for _, c := range s {
		dict[c] = dict[c] + 1
	}

	for _, c := range t {
		dict2[c] = dict2[c] + 1
	}

	for k, v := range dict {
		if v != dict2[k] {
			return false
		}
	}

	return true
}
