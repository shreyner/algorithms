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
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	writer := bufio.NewWriter(os.Stdout)

	scanner.Scan()
	values := strings.Split(scanner.Text(), " ")

	n, _ := strconv.Atoi(values[0])
	m, _ := strconv.Atoi(values[1])

	points := make([][]int, n+1)
	dp := make([][]int, n+1)

	points[n] = make([]int, m+1)
	dp[n] = make([]int, m+1)

	for i := 0; i < len(points)-1; i++ {
		scanner.Scan()
		values := strings.Split(scanner.Text(), "")

		points[i] = make([]int, m+1)
		dp[i] = make([]int, m+1)

		for j := 1; j < len(points[i]); j++ {
			flower, _ := strconv.Atoi(values[j-1])

			points[i][j] = flower
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := 1; j <= m; j++ {
			dp[i][j] = max(dp[i][j-1], dp[i+1][j]) + points[i][j]
		}
	}

	path := make([]string, 0, n+m)

	x, y := 0, m
	for i := 0; i < n; i++ {
		for j := m; j > 0; j-- {
			if x != i || y != j {
				continue
			}

			if dp[i][j-1] >= dp[i+1][j] && j-1 != 0 {
				path = append(path, "R")
				x, y = i, j-1
			} else if i+1 != n {
				path = append(path, "U")
				x, y = i+1, j
			}
		}
	}

	fmt.Fprintln(writer, dp[0][m])

	for i := len(path) - 1; i >= 0; i-- {
		fmt.Fprint(writer, path[i])
	}

	fmt.Fprint(writer, "\n")

	writer.Flush()
}
