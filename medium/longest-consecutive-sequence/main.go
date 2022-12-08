package medium

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	// NOTE: this is not a very good solution. it needs to be O(n) and its slow. very slow.
	if len(nums) <= 1 {
		return len(nums)
	}
	sort.Ints(nums)
	prevCnt := 1
	cnt := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] != 1 {
			// if they are the same number then that doesn't affect anything
			if nums[i+1]-nums[i] == 0 {
				continue
			}

			// if they aren't the same number and they aren't consecutive, reset the count and cache the previous
			if cnt > prevCnt {
				prevCnt = cnt
			}
			cnt = 1
			continue
		}
		cnt += 1
	}

	if prevCnt > cnt {
		return prevCnt
	}
	return cnt
}
