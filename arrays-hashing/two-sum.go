package main

func twoSum(nums []int, target int) []int {
    seen := make(map[int]int, len(nums))

    for idx, num := range nums {
        if seenIdx, ok := seen[target - num]; ok {
            return []int{idx, seenIdx}  
        }

        seen[num] = idx
    }

    return []int{}
}
