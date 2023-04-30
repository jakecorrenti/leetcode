package main

import "sort"

func removeElement(nums []int, val int) int {
    sort.Slice(nums, func (a, b int) bool {
        if nums[a] == val { return false }
        if nums[b] == val { return true }
        return nums[a] > nums[b]
    })

    for idx, n := range nums {
        if n == val {
            return idx
        }
    }
    return len(nums)
}
