package main

import (
	"bufio"
	"fmt"
	"os"
)

const BufferSize = 100000000

func prefixFun(s string) []int {
	dp := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		k := dp[i-1]

		for k > 0 && s[k] != s[i] {
			k = dp[k-1]
		}

		if s[i] == s[k] {
			k += 1
		}

		dp[i] = k
	}

	return dp
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

	res := prefixFun(text)

	for _, v := range res {
		fmt.Fprintf(writer, "%v ", v)
	}
}
