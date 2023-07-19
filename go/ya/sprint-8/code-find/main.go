package main

import "fmt"

func find(text, pattern string, start int) int {
	if len(text) < len(pattern) {
		return -1
	}

	for pos := start; pos <= len(text)-len(pattern); pos++ {
		match := true
		for offset := 0; offset < len(pattern); offset++ {
			if text[pos+offset] != pattern[offset] {
				match = false
				break
			}
		}
		if match {
			return pos
		}
	}

	return -1
}

func findAll(text, pattern string) (occurrences []int) {
	start := 0

	for {
		pos := find(text, pattern, start)

		if pos == -1 {
			break
		}

		occurrences = append(occurrences, pos)
		start = pos + len(pattern)
	}

	return
}

func main() {
	text := "Hello world, world, world"
	pattern := "world"

	fmt.Println(findAll(text, pattern))
}
