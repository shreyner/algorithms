package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printlnArr(arr []int) {
	var sb strings.Builder
	for _, v := range arr {
		fmt.Fprintf(&sb, "%v ", v)
	}
	fmt.Println(sb.String())
}

func BubbleSort(arr []int) {
	n := len(arr)
	sorted := false

	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}

		if swapped == false {
			sorted = true
		}

		if swapped == true {
			printlnArr(arr)
		}

		if sorted == true && n == len(arr) {
			printlnArr(arr)
		}

		n = n - 1
	}
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

	BubbleSort(arr)

	writer.Flush()
}
