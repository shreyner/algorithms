package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverse(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}

	return arr
}

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
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	values := strings.Split(scanner.Text(), " ")

	dp := make([][]int, n+1)

	firstStr := make([]int, n+1)
	for i := 0; i < n; i++ {
		v, _ := strconv.Atoi(values[i])
		firstStr[i+1] = v
	}

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	values = strings.Split(scanner.Text(), " ")

	secondStr := make([]int, m+1)
	for i := 0; i < m; i++ {
		v, _ := strconv.Atoi(values[i])
		secondStr[i+1] = v
	}

	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if firstStr[i] == secondStr[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	firstAnswer := make([]int, 0, len(firstStr))
	secondAnswer := make([]int, 0, len(secondStr))

	x, y := n, m

	for x != 0 && y != 0 {
		if firstStr[x] == secondStr[y] {
			firstAnswer = append(firstAnswer, x)
			secondAnswer = append(secondAnswer, y)

			x -= 1
			y -= 1
		} else {
			if dp[x][y-1] > dp[x-1][y] {
				y -= 1
			} else {
				x -= 1
			}
		}
	}

	fmt.Fprintln(writer, dp[len(dp)-1][len(dp[len(dp)-1])-1])

	if len(firstAnswer) > 0 {
		for i := len(firstAnswer) - 1; i >= 0; i-- {
			fmt.Fprintf(writer, "%v ", firstAnswer[i])
		}
		fmt.Fprint(writer, "\n")
	}

	if len(secondAnswer) > 0 {
		for i := len(secondAnswer) - 1; i >= 0; i-- {
			fmt.Fprintf(writer, "%v ", secondAnswer[i])
		}
		fmt.Fprint(writer, "\n")
	}

	for i := 0; i < len(dp); i++ {
		fmt.Fprintln(writer, dp[i])
	}
}
