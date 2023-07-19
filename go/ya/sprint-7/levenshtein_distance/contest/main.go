package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func min(a, b int) int {
	if a < b {
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
	firstStr := strings.Split(scanner.Text(), "")
	scanner.Scan()
	secondStr := strings.Split(scanner.Text(), "")

	n, m := len(firstStr), len(secondStr)

	if n > m {
		n, m = m, n
		firstStr, secondStr = secondStr, firstStr
	}

	prevDp := make([]int, m+1)
	currentDp := make([]int, m+1)
	for i := 1; i < m+1; i++ {
		prevDp[i] = i
	}

	for i := 1; i <= n; i++ {
		currentDp[0] = i

		for j := 1; j <= m; j++ {
			change := 1
			if firstStr[i-1] == secondStr[j-1] {
				change = 0
			}

			currentDp[j] = min(prevDp[j]+1, min(currentDp[j-1]+1, prevDp[j-1]+change))
		}

		prevDp, currentDp = currentDp, prevDp
	}

	fmt.Fprintln(writer, prevDp[m])
}
