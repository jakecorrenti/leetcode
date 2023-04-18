package main

import (
    "sort"
)

func groupAnagrams(strs []string) [][]string {
    // sorted string : [array of original strings]
    seen := make(map[string][]string, len(strs))

    for _, str := range strs {
        stringByteSlice := []byte(str)
        sort.Slice(stringByteSlice, func (a, b int) bool {
            return stringByteSlice[a] < stringByteSlice[b]
        })
        if _, ok := seen[string(stringByteSlice)]; ok {
            seen[string(stringByteSlice)] = append(seen[string(stringByteSlice)], str)
        } else {
            seen[string(stringByteSlice)] = []string{str}
        }
    }

    output := [][]string{}
    for _, v := range seen {
        output = append(output, v)
    }

    return output
}

func isAnagram(first string, second string) bool {

    if len(first) != len(second) {
        return false
    }

    firstStringRuneSeen := make([]int, 26)
    secondStringRuneSeen := make([]int, 26)

    for idx, r := range first {
        firstStringRuneSeen[byte(r) - 'a']++
        secondStringRuneSeen[second[idx] - 'a']++
    }

    for i := 0; i < 26; i++ {
        if firstStringRuneSeen[i] != secondStringRuneSeen[i] {
            return false
        }
    }

    return true 
}
