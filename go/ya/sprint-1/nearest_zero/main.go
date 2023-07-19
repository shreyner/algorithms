package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func NearestZero(arr []int) []int {
	resultArr := make([]int, len(arr))

	count := -1

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			count = 0
		}

		if count == -1 {
			continue
		}

		resultArr[i] = count

		count += 1
	}

	count = -1

	for j := len(arr) - 1; j >= 0; j-- {
		if arr[j] == 0 {
			count = 0
		}

		if count == -1 {
			continue
		}

		if resultArr[j] == 0 || resultArr[j] > count {
			resultArr[j] = count
		}

		count += 1
	}

	return resultArr
}

func main() {
	var n int
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)

	houseNumbers := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		stringHouseNumber := scanner.Text()

		number, _ := strconv.Atoi(stringHouseNumber)

		houseNumbers[i] = number
	}

	result := NearestZero(houseNumbers)

	for _, v := range result {
		fmt.Fprintf(writer, "%v ", v)
	}

	writer.WriteString("\n")
	writer.Flush()
}
