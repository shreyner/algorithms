package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertToList(vCount int, edges [][]int) [][]int {
	adj := make([][]int, vCount)

	for _, edge := range edges {
		u, v := edge[0], edge[1]

		adj[u-1] = append(adj[u-1], v)
	}

	return adj
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

	grphList := convertToList(vCount, edges)

	for i := 0; i < len(grphList); i++ {
		if len(grphList[i]) == 0 {
			fmt.Fprintln(writer, "0")
			continue
		}

		fmt.Fprintf(writer, "%v ", len(grphList[i]))
		for j := 0; j < len(grphList[i]); j++ {
			fmt.Fprintf(writer, "%v ", grphList[i][j])
		}
		fmt.Fprint(writer, "\n")
	}

	writer.Flush()
}
