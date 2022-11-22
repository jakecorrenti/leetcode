package medium

import "sort"

// runtime: 28ms, faster than 85.06%
// memory usage: 7.9 MB, less than 75.88%

func groupAnagrams(strs []string) [][]string {

	groups := [][]string{}
	idxs := map[string][]int{}

	// iterate through all of the strings
	for idx, s := range strs {
		// make a copy of the string, because sorting will modify the actual string
		tmpStr := s
		bArr := []byte(tmpStr)
		// sort the characters in the byte array so that they are in ascending order (a -> z)
		sort.Slice(bArr, func(a, b int) bool {
			return bArr[a] < bArr[b]
		})
		tmpStr = string(bArr)

		// check if the map contains the sorted string already
		if _, ok := idxs[tmpStr]; ok {
			// if the map contains the sorted string, add the index it was found at to the list
			idxs[tmpStr] = append(idxs[tmpStr], idx)
			continue
		}
		// if the sorted string wasn't previously seen, add it to the map
		idxs[tmpStr] = []int{idx}
	}

	// iterate through all of the sorted strings that were seen
	for _, v := range idxs {
		vals := []string{}
		// iterate through each of the lists of indexes for the strings, and add the strings at those indexes to their own groups
		for _, idx := range v {
			vals = append(vals, strs[idx])
		}
		groups = append(groups, vals)
	}

	return groups
}
