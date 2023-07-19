package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func convertToList(vCount int, edges [][]int) [][]int {
	adj := make([][]int, vCount+1)

	for _, edge := range edges {
		u, v := edge[0], edge[1]

		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	for _, ints := range adj {
		sort.Slice(ints, func(i, j int) bool {
			return ints[i] > ints[j]
		})
	}

	return adj
}

func dfs(vertexMatrix [][]int, startVertex int) []int {
	stack := make([]int, 0, len(vertexMatrix))
	path := make([]int, 0, len(vertexMatrix))
	color := make([]int, len(vertexMatrix))

	stack = append(stack, startVertex)

	for len(stack) != 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if color[v] == 0 {
			color[v] = 1
			stack = append(stack, v)
			path = append(path, v)

			for _, w := range vertexMatrix[v] {
				if color[w] != 0 {
					continue
				}

				stack = append(stack, w)
			}
		} else if color[v] == 1 {
			color[v] = 2
		}
	}

	return path
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

	scanner.Scan()
	startVertexStr, _ := strconv.Atoi(scanner.Text())

	vertexList := convertToList(vCount, edges)
	historyPath := dfs(vertexList, startVertexStr)

	for _, pathVertex := range historyPath {
		fmt.Fprintf(writer, "%v ", pathVertex)
	}

	writer.Flush()
}
