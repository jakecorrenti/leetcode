package easy

func countBits(n int) []int {
	ret := make([]int, n+1)

	for i := 1; i <= n; i++ {
		count := 0
		tmp := i

		for tmp != 0 {
			if tmp&1 != 0 {
				count += 1
			}
			tmp >>= 1
		}
		ret[i] = count
	}
	return ret
}
