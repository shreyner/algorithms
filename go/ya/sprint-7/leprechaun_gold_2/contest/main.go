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
	values := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(values[0])
	capacity, _ := strconv.Atoi(values[1])

	scanner.Scan()
	values = strings.Split(scanner.Text(), " ")

	goldBars := make([]int, n)
	for i := 0; i < n; i++ {
		v, _ := strconv.Atoi(values[i])
		goldBars[i] = v
	}

	prevDp := make([]int, capacity+1)
	currentDp := make([]int, capacity+1)

	for i := 0; i < n; i++ {
		for j := 1; j < len(currentDp); j++ {
			if j-goldBars[i] < 0 {
				currentDp[j] = prevDp[j]
			} else {
				currentDp[j] = max(prevDp[j], goldBars[i]+(prevDp[j-goldBars[i]]))
			}
		}

		prevDp, currentDp = currentDp, prevDp
	}

	fmt.Fprintln(writer, prevDp[len(prevDp)-1])
}
