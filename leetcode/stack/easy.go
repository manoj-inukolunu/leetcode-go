package main

import (
	"fmt"
)

type stack []rune

func (s *stack) push(num rune) {
	*s = append(*s, num)
}

func (s stack) size() int {
	return len(s)
}

func (s *stack) pop() rune {
	ret := (*s)[s.size()-1]
	*s = (*s)[0 : s.size()-1]
	return ret
}

func (s *stack) peek() rune {
	ret := (*s)[s.size()-1]
	return ret

}

func isValid(s string) bool {
	st := make(stack, 0)
	for i := 0; i < len(s); i++ {
		if st.size() == 0 {
			st.push(rune(s[i]))
		} else if s[i] == ')' && rune(st.peek()) == '(' {
			st.pop()
		} else if s[i] == ']' && rune(st.peek()) == '[' {
			st.pop()
		} else if s[i] == '}' && rune(st.peek()) == '{' {
			st.pop()
		} else {
			st.push(rune(s[i]))
		}
	}
	return st.size() == 0
}

func _main() {
	fmt.Println(isValid("(]"))
}
