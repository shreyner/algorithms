package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	var n int
	fmt.Scan(&n)

	var m int
	fmt.Scan(&m)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fields := make([][]int, n)

	for i := n - 1; i >= 0; i-- {
		fields[i] = make([]int, m)

		scanner.Scan()
		values := strings.Split(scanner.Text(), "")

		for j := 0; j < m; j++ {
			v, _ := strconv.Atoi(values[j])
			fields[i][j] = v
		}
	}

	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}

	for j := 1; j <= m; j++ {
		for i := 1; i <= n; i++ {
			dp[i][j] = max(dp[i][j-1], dp[i-1][j]) + fields[i-1][j-1]
		}
	}

	fmt.Fprintln(writer, dp[n][m])
}
