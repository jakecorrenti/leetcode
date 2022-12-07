package easy

import (
	"strings"
)

func halvesAreAlike(s string) bool {
	s = strings.ToLower(s)
	vowels := "aeiou"
	h1 := s[:len(s)/2]
	h2 := s[len(s)/2:]

	h1Cnt := 0
	h2Cnt := 0
	for i := 0; i < len(h1); i++ {
		if strings.Contains(vowels, string(h1[i])) {
			h1Cnt += 1
		}
		if strings.Contains(vowels, string(h2[i])) {
			h2Cnt += 1
		}
	}
	return h1Cnt == h2Cnt
}
