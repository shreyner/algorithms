package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	days := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		days[i] = v
	}

	sum := 0

	for i := 1; i < n; i++ {
		if days[i-1] < days[i] {
			sum += days[i] - days[i-1]
		}
	}

	fmt.Fprintln(writer, sum)
}
