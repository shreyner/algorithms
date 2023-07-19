package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countWays(n, k int) int {
	mod := 1000000007

	if n < 1 {
		return 0
	}

	dp := make([]int, n+1)

	dp[1] = 1

	for i := 2; i < len(dp); i++ {
		if k >= i {
			for j := 1; j < i; j++ {
				dp[i] += dp[j]
			}
		} else {
			for j := i - k; j < i; j++ {
				dp[i] += dp[j]
			}
		}

		dp[i] = dp[i] % mod
	}

	return dp[len(dp)-1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	values := strings.Split(scanner.Text(), " ")

	n, _ := strconv.Atoi(values[0])
	k, _ := strconv.Atoi(values[1])

	writer := bufio.NewWriter(os.Stdout)

	res := countWays(n, k)

	fmt.Fprintln(writer, res)
	writer.Flush()
}
