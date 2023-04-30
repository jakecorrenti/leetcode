package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int)  {
    if n < 1 {
        return
    }

    for i := m; i < (m + n); i++ {
        nums1[i] = nums2[i - m]
    }

    sort.Ints(nums1)
}
