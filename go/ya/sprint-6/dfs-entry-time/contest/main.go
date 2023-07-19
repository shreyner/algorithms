package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Color int

const (
	WHITE Color = iota
	GRAY
	BLACK
)

func convertToList(vCount int, edges [][]int) [][]int {
	adj := make([][]int, vCount+1)

	for _, edge := range edges {
		u, v := edge[0], edge[1]

		adj[u] = append(adj[u], v)
	}

	for _, ints := range adj {
		sort.Slice(ints, func(i, j int) bool {
			return ints[i] > ints[j]
		})
	}

	return adj
}

func dfs(vertexList [][]int, startVertex int) [][]int {
	time := 0
	stack := make([]int, 0, len(vertexList))
	color := make([]Color, len(vertexList))
	entry := make(map[int]int, len(vertexList))
	leave := make(map[int]int, len(vertexList))

	stack = append(stack, startVertex)
	entry[startVertex] = time

	for len(stack) != 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if color[v] == WHITE {
			color[v] = GRAY
			stack = append(stack, v)
			entry[v] = time
			time += 1

			for _, w := range vertexList[v] {
				if color[w] != 0 {
					continue
				}

				stack = append(stack, w)
			}
		} else if color[v] == GRAY {
			color[v] = BLACK
			leave[v] = time
			time += 1
		}
	}

	result := make([][]int, len(vertexList))
	for i := 1; i < len(vertexList); i++ {
		result[i] = []int{entry[i], leave[i]}
	}

	return result
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

	vertexList := convertToList(vCount, edges)
	times := dfs(vertexList, 1)

	for i := 1; i < len(times); i++ {
		time := times[i]
		fmt.Fprintf(writer, "%v %v\n", time[0], time[1])
	}

	writer.Flush()
}
