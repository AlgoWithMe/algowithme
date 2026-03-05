package validparentheses

func ValidParentheses(input string) bool {
	stack := []rune{}
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, c := range input {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else if open, ok := pairs[c]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
