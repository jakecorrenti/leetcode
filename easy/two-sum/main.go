package easy

func twoSum(nums []int, target int) []int {
	seen := map[int]int{}
	for idx, n := range nums {
		if _, ok := seen[target-n]; ok {
			return []int{seen[target-n], idx}
		}
		seen[n] = idx
	}
	return []int{}
}
