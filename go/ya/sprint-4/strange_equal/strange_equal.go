package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hashGen() func(key rune) int {
	mapKeys := make(map[rune]int)
	lastIndex := 0

	return func(key rune) int {
		keyIndex, ok := mapKeys[key]

		if ok {
			return keyIndex
		}

		lastIndex++
		mapKeys[key] = lastIndex
		return lastIndex
	}
}

func StrangeEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	aHash := hashGen()
	bHash := hashGen()

	for i := 0; i < len(a); i++ {
		if aHash(a[i]) != bHash(b[i]) {
			return false
		}
	}

	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	aRowStr, _ := reader.ReadString('\n')
	bRowStr, _ := reader.ReadString('\n')

	writer := bufio.NewWriter(os.Stdout)

	aStr := []rune(strings.TrimSpace(aRowStr))
	bStr := []rune(strings.TrimSpace(bRowStr))

	res := StrangeEqual(aStr, bStr)

	if res {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}

	writer.Flush()
}
