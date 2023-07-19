package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ContestFind(arr []int) int {
	mapGraph := make(map[int]int)

	max := 0
	y := 0

	mapGraph[0] = 0

	for i := 0; i < len(arr); i++ {
		x := i + 1

		if arr[i] == 0 {
			y -= 1
		} else {
			y += 1
		}

		prevX, ok := mapGraph[y]

		if !ok {
			mapGraph[y] = x

			continue
		}

		if max < x-prevX {
			max = x - prevX
		}
	}

	return max
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
		value := scanner.Text()

		num, _ := strconv.Atoi(value)
		arr[i] = num
	}

	res := ContestFind(arr)

	fmt.Println(res)

	writer.Flush()
}
