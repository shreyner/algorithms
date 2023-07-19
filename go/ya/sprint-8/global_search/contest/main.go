package main

import (
	"bufio"
	"fmt"
	"os"
)

const BufferSize = 100000000

func prefixFun(text, pattern string) (result []int) {
	dp := make([]int, len(pattern))
	dpPrev := 0
	s := pattern + "#" + text

	for i := 1; i < len(s); i++ {
		k := dpPrev

		for k > 0 && s[k] != s[i] {
			k = dp[k-1]
		}

		if s[i] == s[k] {
			k += 1
		}

		if i < len(pattern) {
			dp[i] = k
		}

		dpPrev = k

		if k == len(pattern) {
			result = append(result, i-2*len(pattern))
		}
	}

	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	scanner.Buffer(make([]byte, BufferSize), BufferSize)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	text := scanner.Text()

	scanner.Scan()
	pattern := scanner.Text()

	scanner.Scan()
	replace := scanner.Text()

	pref := prefixFun(text, pattern)

	start := 0
	for _, pos := range pref {
		fmt.Fprint(writer, text[start:pos])
		fmt.Fprint(writer, replace)

		start = pos + len(pattern)
	}

	if start < len(text) {
		fmt.Fprint(writer, text[start:])
	}
}
