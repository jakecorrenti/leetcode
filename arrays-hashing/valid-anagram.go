package main

// my solution
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sChars := make(map[byte]int, len(s))
    tChars := make(map[byte]int, len(t))

    for i, val := range s {
        sChars[byte(val)]++
        tChars[t[i]]++
    }

    for k, v := range sChars {
        if v != tChars[k] {
            return false
        }
    }

    return true
}

// better solution (takes up less space)
/*
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sArr := make([]int, 26)
    tArr := make([]int, 26)

    for i, c := range s {
        sArr[byte(c) - byte('a')] += 1
        tArr[byte(t[i]) - byte('a')] += 1
    }

    for i, c := range sArr {
        if c != tArr[i] {
            return false 
        }
    }

    return true
}
*/
