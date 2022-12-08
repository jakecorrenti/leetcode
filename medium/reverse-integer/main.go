package medium

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	isNeg := false
	str := strconv.Itoa(x)
	if x < 0 {
		str = str[1:]
		isNeg = true
	}

	rev := ""
	if isNeg {
		rev += "-"
	}
	for i := len(str) - 1; i >= 0; i-- {
		s := string(str[i])
		if (rev == "-" || rev == "") && s == "0" {
			continue
		}
		rev += string(str[i])
	}

	i, _ := strconv.Atoi(rev)
	if i >= math.MaxInt32 || i <= math.MinInt32 {
		return 0
	}
	return i
}
