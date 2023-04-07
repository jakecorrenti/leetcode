package main

func containsDuplicate(nums []int) bool {
    seen := make(map[int]bool, len(nums))

    for _, n := range nums {
        if _, ok := seen[n]; ok {
            return true
        }

        seen[n] = true
    }

    return false
}
