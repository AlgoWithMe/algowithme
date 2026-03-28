package validparentheses

func ValidParentheses(input string) bool {
	stack := []rune{}
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range input {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != mapping[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
