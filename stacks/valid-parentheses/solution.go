package stack

import lls "github.com/emirpasic/gods/stacks/linkedliststack"

// IsValid returns true if the input string is a valid sequence of parentheses.
var mapping = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

func IsValid(s string) bool {
	stack := lls.New()
	for _, v := range s {
		switch v {
		case '{', '[', '(':
			stack.Push(v)

		case ')', '}', ']':
			if val, ok := stack.Pop(); !ok || val != mapping[v] {
				return false
			}
		}
	}

	return stack.Empty()
}
