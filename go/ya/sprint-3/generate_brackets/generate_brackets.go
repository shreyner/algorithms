package main

import (
	"fmt"
)

func GenerateBrackets(n, open, close int, sb []rune) {
	if n == open && n == close {
		fmt.Println(string(sb))

		return
	}

	if open < n {
		GenerateBrackets(n, open+1, close, append(sb, '('))
	}

	if close < open {
		GenerateBrackets(n, open, close+1, append(sb, ')'))
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	GenerateBrackets(n, 0, 0, []rune{})
}
