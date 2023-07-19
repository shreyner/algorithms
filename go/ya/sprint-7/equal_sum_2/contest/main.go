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

	sum := 0
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())

		numbers[i] = v
		sum += v
	}

	if sum%2 != 0 {
		fmt.Fprintln(writer, "False")
		return
	}

	prevDp := make([]bool, sum/2+1)
	currentDp := make([]bool, sum/2+1)

	for i := 0; i < len(numbers); i++ {
		for j := 1; j < len(currentDp); j++ {
			currentDp[j] = prevDp[j]
			if j == numbers[i] {
				currentDp[j] = true
			}

			if j > numbers[i] && prevDp[j-numbers[i]] {
				currentDp[j] = true
			}
		}

		currentDp, prevDp = prevDp, currentDp
	}

	if prevDp[len(prevDp)-1] {
		fmt.Fprintln(writer, "True")
	} else {
		fmt.Fprintln(writer, "False")
	}
}
