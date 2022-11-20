package main

func containsDuplicate(nums []int) bool {
	occur := make(map[int]bool, 0)
	for _, n := range nums {
		if _, ok := occur[n]; ok {
			return true
		}
		occur[n] = true
	}
	return false
}
