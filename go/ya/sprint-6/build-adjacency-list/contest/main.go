package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertTo(edges [][]int) map[int][]int {
	adj := make(map[int][]int, len(edges))

	for _, edge := range edges {
		u, v := edge[0], edge[1]

		adj[u] = append(adj[u], v)
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

	adj := convertTo(edges)

	for i := 0; i < vCount; i++ {
		edge := adj[i+1]

		if len(edge) == 0 {
			fmt.Fprintln(writer, 0)
			continue
		}

		fmt.Fprintf(writer, "%v ", len(edge))
		for _, e := range edge {
			fmt.Fprintf(writer, "%v ", e)
		}
		fmt.Fprint(writer, "\n")
	}

	writer.Flush()
}
