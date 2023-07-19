package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 1000000007

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	dp := make([]int, n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		if k > i {
			for j := 0; j < i; j++ {
				dp[i] += dp[j]
			}
		} else {
			for j := i - k; j < i; j++ {
				dp[i] += dp[j]
			}
		}

		dp[i] = dp[i] % mod
	}

	fmt.Fprintln(writer, dp[n-1])
}
