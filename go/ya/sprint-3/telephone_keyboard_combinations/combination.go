package main

import (
	"fmt"
)

var mapKeyboard = map[string][]rune{
	"2": {'a', 'b', 'c'},
	"3": {'d', 'e', 'f'},
	"4": {'g', 'h', 'i'},
	"5": {'j', 'k', 'l'},
	"6": {'m', 'n', 'o'},
	"7": {'p', 'q', 'r', 's'},
	"8": {'t', 'u', 'v'},
	"9": {'w', 'x', 'y', 'z'},
}

func Combinations(n int, num string, prefix []rune) {
	if len(num) == n {
		fmt.Print(string(append(prefix, ' ')))
		return
	}

	number := string(num[n])
	charNumber := mapKeyboard[number]
	for i := 0; i < len(charNumber); i++ {
		Combinations(n+1, num, append(prefix, charNumber[i]))
	}
}

func main() {
	var num string
	fmt.Scan(&num)

	Combinations(0, num, []rune{})
}
