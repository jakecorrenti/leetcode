package medium

func twoSum(numbers []int, target int) []int {
    p1 := 0
    p2 := len(numbers) - 1

    for p1 != p2 {
        if numbers[p1] + numbers[p2] < target {
            p1++
        } else if numbers[p1] + numbers[p2] > target {
            p2--
        } else {
            numbers[0] = p1 + 1
            numbers[1] = p2 + 1
            return numbers[:2]
        }
    }

    return []int{}
}
