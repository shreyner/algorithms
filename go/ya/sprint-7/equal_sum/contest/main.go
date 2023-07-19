package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum(numbers []int) (s int) {
	for _, number := range numbers {
		s += number
	}

	return
}

func hasEqualSum(numbers []int) bool {
	k := sum(numbers)
	if k%2 != 0 {
		return false
	}

	dp := make([]bool, k/2+1)
	dp[0] = true

	for i := 0; i < len(numbers); i++ {
		for j := k / 2; j >= numbers[i]; j-- {
			if j-numbers[i] >= 0 {
				dp[j] = dp[j-numbers[i]] || dp[j]
			}

			if j == k/2 && dp[j] {
				return true
			}
		}
	}

	return dp[len(dp)-1]
}

func main() {
	var n int
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())

		numbers[i] = v
	}

	res := hasEqualSum(numbers)

	if res {
		fmt.Fprintln(writer, "True")
	} else {
		fmt.Fprintln(writer, "False")
	}

	writer.Flush()
}
