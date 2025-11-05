package main

func main() {
	// Code
}
func isValid(s string) bool {
	stack := []rune{}
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			if char == ')' && stack[len(stack)-1] != '(' {
				return false
			}
			if char == ']' && stack[len(stack)-1] != '[' {
				return false
			}
			if char == '}' && stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
