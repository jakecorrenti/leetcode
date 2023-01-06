package easy

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}

	stack := []rune{}

	for _, r := range s {
		if len(stack) >= 1 {
			top := stack[len(stack)-1]

			if (r == ')' && top == '(') ||
				(r == ']' && top == '[') ||
				(r == '}' && top == '{') {
				stack = stack[:len(stack)-1]
				continue
			}
		}

		stack = append(stack, r)
	}

	return len(stack) < 1
}
