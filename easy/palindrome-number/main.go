package easy

import (
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	sBytes := []byte(s)
	sBytesCopy := []byte{}
	for i := len(sBytes) - 1; i >= 0; i-- {
		sBytesCopy = append(sBytesCopy, sBytes[i])
	}
	return string(sBytesCopy) == string(sBytes)
}
