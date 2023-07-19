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
	defer writer.Flush()

	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	values := strings.Split(scanner.Text(), " ")

	aStr := make([]int, a)
	for i := 0; i < a; i++ {
		v, _ := strconv.Atoi(values[i])
		aStr[i] = v
	}

	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	values = strings.Split(scanner.Text(), " ")

	bStr := make([]int, b)
	for i := 0; i < b; i++ {
		v, _ := strconv.Atoi(values[i])
		bStr[i] = v
	}

	dp := make([][]int, b+1)
	dp[0] = make([]int, a+1)

	for i := 1; i < b; i++ {
		dp[i] = make([]int, a)
		for j := 1; j < a; j++ {
			if bStr[i-1] == aStr[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	fmt.Fprintln(writer, aStr, bStr)

	for i := 0; i < len(dp); i++ {
		fmt.Fprintln(writer, dp[i])
	}

	maxR := dp[b-1][a-1]
	fmt.Fprintln(writer, maxR)

	aAnswer := make([]int, maxR)
	bAnswer := make([]int, maxR)

	x, y := a, b

	for x != 0 && b != 0 {
		if aStr[x-1] == bStr[b-1] {
			aAnswer[dp[y][x]-1] = x
			bAnswer[dp[y][x]-1] = b

			y -= 1
			x -= 1
		} else {
			if dp[y-1][x] > dp[y][x-1] {
				y -= 1
			} else {
				x -= 1
			}

		}
	}

	fmt.Fprintln(writer, aAnswer, bAnswer)
}
