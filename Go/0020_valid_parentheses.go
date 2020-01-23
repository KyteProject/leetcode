/*
https://leetcode.com/problems/valid-parentheses/

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
	Open brackets must be closed by the same type of brackets.
	Open brackets must be closed in the correct order.
	An empty string.
*/

package leetcode

// Stack solution O(n)
func isValid(s string) bool {
	stack := make([]rune, 0)
	m := map[rune]rune{'(': ')', '[': ']', '{': '}'}

	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, m[c])
		case ')', ']', '}':
			if len(stack) > 0 && stack[len(stack)-1] == c {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
