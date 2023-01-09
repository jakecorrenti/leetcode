package medium

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	// reverse polish notation is just postfix notation
	// ex. 2 3 + == 2 + 3

	stack := []int{}
	for _, tok := range tokens {
		dig, err := strconv.Atoi(tok)
		if err == nil {
			stack = append(stack, dig)
			continue
		}

		op2 := stack[len(stack)-1]
		op1 := stack[len(stack)-2]
		stack = stack[:len(stack)-2]

		switch tok {
		case "+":
			stack = append(stack, op1+op2)
		case "-":
			stack = append(stack, op1-op2)
		case "*":
			stack = append(stack, op1*op2)
		case "/":
			stack = append(stack, op1/op2)
		}
	}

	return stack[0]
}
