package easy

func singleNumber(nums []int) int {
	m := map[int]int{}
	for _, n := range nums {
		m[n] += 1
	}
	for k, v := range m {
		if v < 2 {
			return k
		}
	}
	return -1
}
