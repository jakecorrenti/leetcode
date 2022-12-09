package medium

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	pre := 1
	post := 1

	output[0] = 1
	// do the first run of the data calculating the prefix values
	for i := 1; i < len(nums); i++ {
		pre *= nums[i-1]
		output[i] = pre
	}

	// do the second run of the data from the other direction calculating
	// the postfix values and mulitplying them by the prefix values of that index
	for i := len(nums) - 2; i >= 0; i-- {
		post *= nums[i+1]
		output[i] *= post
	}

	return output
}
