package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	WHITE = 0
	GREY  = -1
)

func convertToList(vCount int, edges [][]int) [][]int {
	adj := make([][]int, vCount+1)

	for _, edge := range edges {
		u, v := edge[0], edge[1]

		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	for _, ints := range adj {
		sort.Ints(ints)
	}

	return adj
}

func dfsRelationsComponents(vertexList [][]int) (int, []int) {
	stack := make([]int, 0, len(vertexList))
	colors := make([]int, len(vertexList))
	countColor := 0

	for i := 1; i < len(vertexList); i++ {
		if colors[i] == WHITE {
			stack = append(stack, i)
			countColor += 1
		}

		for len(stack) != 0 {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if colors[v] == WHITE {
				colors[v] = GREY

				stack = append(stack, v)

				for _, vertex := range vertexList[v] {
					if colors[vertex] != WHITE {
						continue
					}

					stack = append(stack, vertex)
				}
			} else if colors[v] == GREY {
				colors[v] = countColor
			}
		}
	}

	return countColor, colors
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

	countColor, colors := dfsRelationsComponents(vertexList)

	fmt.Fprintf(writer, "%v\n", countColor)

	for i := 1; i <= countColor; i++ {
		for j := 1; j < len(vertexList); j++ {
			if colors[j] != i {
				continue
			}

			fmt.Fprintf(writer, "%v ", j)
		}

		fmt.Fprintf(writer, "\n")
	}

	writer.Flush()
}
