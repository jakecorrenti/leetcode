package medium

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}
	occur := make([][]int, len(nums)+1)
	hash := map[int]int{}
	out := []int{}

	// create a hash that contains the occurences of each integer
	for _, n := range nums {
		hash[n] += 1
	}

	// use the bucket sort algorithm that maps the occurences to the index and the value at that index is a list of integers with that occurence
	for k, v := range hash {
		occur[v] = append(occur[v], k)
	}

	// starting at the end of the array(since the index is mapped to the # of occurences), find the first k elements
	for i := len(occur) - 1; i >= 0; i-- {
		if len(occur[i]) > 0 && len(out) < k {
			out = append(out, occur[i]...)
		}
	}

	return out
}
