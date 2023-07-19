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

	scanner.Scan()
	values = strings.Split(scanner.Text(), " ")

	goldBars := make([]int, n+1)

	for i := 0; i < n; i++ {
		costGoldBar, _ := strconv.Atoi(values[i])
		goldBars[i+1] = costGoldBar
	}

	prevDp := make([]int, m+1)
	currentDp := make([]int, m+1)

	for i := 1; i < len(goldBars); i++ {
		for j := 0; j < len(currentDp); j++ {
			if j-goldBars[i] >= 0 {
				currentDp[j] = max(prevDp[j], goldBars[i]+prevDp[j-goldBars[i]])
			} else {
				currentDp[j] = prevDp[j]
			}
		}

		prevDp, currentDp = currentDp, prevDp
	}

	fmt.Fprintln(writer, prevDp[len(prevDp)-1])

	writer.Flush()
}
