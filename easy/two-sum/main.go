package main

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int, len(nums))

	for i, num := range nums {
		if val, ok := dict[num]; ok {
			return []int{val, i}
		}
		dict[target-num] = i
	}
	return []int{}
}
