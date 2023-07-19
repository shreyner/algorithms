package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SleightOfHand(arr []int, k int) int {
	numbers := make([]int, 10)
	fingerCount := k * 2

	for i := 0; i < len(arr); i++ {
		if arr[i] == -1 {
			continue
		}

		numbers[arr[i]] += 1
	}

	points := 0
	for _, numberCount := range numbers {
		if numberCount == 0 || numberCount > fingerCount {
			continue
		}

		points += 1
	}

	return points

}

func main() {
	var k int
	fmt.Scan(&k)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	writer := bufio.NewWriter(os.Stdout)

	numbers := make([]int, 0)
	for scanner.Scan() {
		lineNumbers := scanner.Text()
		stringNumbers := strings.Split(lineNumbers, "")

		for i := 0; i < len(stringNumbers); i++ {
			if stringNumbers[i] == "." {
				numbers = append(numbers, -1)
				continue
			}

			parsed, _ := strconv.Atoi(stringNumbers[i])
			numbers = append(numbers, parsed)
		}

	}

	result := SleightOfHand(numbers, k)

	fmt.Fprintf(writer, "%v\n", result)
	writer.Flush()
}
