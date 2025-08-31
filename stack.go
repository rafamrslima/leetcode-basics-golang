func validParentheses(s string) bool {
	dict := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := make([]rune, 0, len(s))

	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != dict[ch] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
