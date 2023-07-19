package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func find(text, pattern []int, start int) int {
	for pos := start; pos <= len(text)-len(pattern); pos++ {
		match := true
		c := pattern[0] - text[pos]

		for offset := 1; offset < len(pattern); offset++ {
			if text[pos+offset]+c != pattern[offset] {
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	weathers := make([]int, n)
	scanner.Scan()
	for i, v := range strings.Split(scanner.Text(), " ") {
		temp, _ := strconv.Atoi(v)
		weathers[i] = temp
	}

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	pattern := make([]int, m)
	scanner.Scan()
	for i, v := range strings.Split(scanner.Text(), " ") {
		temp, _ := strconv.Atoi(v)
		pattern[i] = temp
	}

	if len(weathers) < len(pattern) {
		return
	}

	start := 0
	for {
		pos := find(weathers, pattern, start)

		if pos == -1 {
			break
		}

		fmt.Fprintf(writer, "%v ", pos+1)
		start = pos + 1
	}
}
