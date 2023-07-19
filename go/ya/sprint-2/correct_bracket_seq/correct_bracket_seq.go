package main

import (
	"bufio"
	"fmt"
	"os"
)

var mapBracket = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
}

func IsCorrectBracketSeq(str []rune) bool {
	if len(str)%2 == 1 {
		return false
	}

	stack := make([]rune, 0)

	for _, c := range str {
		_, ok := mapBracket[c]

		if ok {
			stack = append(stack, c)
			continue
		}

		if len(stack) == 0 || mapBracket[stack[len(stack)-1]] != c {
			return false
		}

		stack = stack[:len(stack)-1]

	}

	if len(stack) != 0 {
		return false
	}

	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)

	scanner.Scan()
	line := scanner.Text()

	res := IsCorrectBracketSeq([]rune(line))

	if res {
		fmt.Fprintln(writer, "True")
	} else {
		fmt.Fprintln(writer, "False")
	}

	writer.Flush()
}
