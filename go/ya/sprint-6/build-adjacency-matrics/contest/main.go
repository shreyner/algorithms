package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertTo(vCount int, edges [][]int) [][]int {
	matrix := make([][]int, vCount)
	for i := 0; i < vCount; i++ {
		matrix[i] = make([]int, vCount)
	}

	for _, edge := range edges {
		matrix[edge[0]-1][edge[1]-1] = 1
	}

	return matrix
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)
	scanner.Scan()
	values := strings.Split(scanner.Text(), " ")
	vCount, _ := strconv.Atoi(values[0])
	n, _ := strconv.Atoi(values[1])

	edges := make([][]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")
		u, _ := strconv.Atoi(line[0])
		v, _ := strconv.Atoi(line[1])

		edges[i] = []int{u, v}
	}

	matrics := convertTo(vCount, edges)

	for i := 0; i < len(matrics); i++ {
		for j := 0; j < len(matrics[i]); j++ {
			fmt.Fprintf(writer, "%v ", matrics[i][j])
		}
		fmt.Fprint(writer, "\n")
	}

	writer.Flush()
}
