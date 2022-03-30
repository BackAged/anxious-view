// Problem: https://leetcode.com/problems/valid-parentheses/

package main

// top returns the top element of the stack
func top(stack []rune) rune {
	if len(stack) == 0 {
		panic("invalid operation")
	}

	return stack[len(stack)-1]
}

// pop removes the top element from the stack
func pop(stack []rune) []rune {
	if len(stack) == 0 {
		panic("invalid operation")
	}

	return stack[:len(stack)-1]
}

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))

	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, r := range s {
		if r == '(' || r == '{' || r == '[' {
			stack = append(stack, r)

			continue
		}

		if len(stack) == 0 {
			return false
		}

		if top(stack) != bracketMap[r] {
			return false
		}

		stack = pop(stack)
	}

	if len(stack) == 0 {
		return true
	}

	return false
}
