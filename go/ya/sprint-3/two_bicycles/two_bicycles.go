package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TwoBicycles(arr []int, x int) (int, int) {
	if len(arr) == 0 {
		return -1, -1
	}

	_, firstIndex := binarySearch(arr, x, 0, len(arr))
	_, secondIndex := binarySearch(arr, x*2, 0, len(arr))

	if firstIndex != -1 {
		firstIndex += 1
	}

	if secondIndex != -1 {
		secondIndex += 1
	}

	return firstIndex, secondIndex
}

func binarySearch(arr []int, x, left, right int) (int, int) {
	if right <= left {
		return 0, -1
	}

	mid := (left + right) / 2
	item := arr[mid]

	if x <= item {
		i, index := binarySearch(arr, x, left, mid)

		if index != -1 && index < mid {
			return i, index
		}

		return item, mid
	}

	return binarySearch(arr, x, mid+1, right)
}

func main() {
	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		word := scanner.Text()

		value, _ := strconv.Atoi(word)
		arr[i] = value
	}

	scanner.Scan()
	word := scanner.Text()
	target, _ := strconv.Atoi(word)

	first, second := TwoBicycles(arr, target)

	fmt.Fprintf(writer, "%v %v\n", first, second)

	writer.Flush()
}
