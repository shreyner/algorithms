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
	GREY
	BLACK
)

func convertToList(vCount int, edges [][]int) [][]int {
	adj := make([][]int, vCount+1)

	for _, edge := range edges {
		u, v := edge[0], edge[1]

		adj[u] = append(adj[u], v)
	}

	for _, ints := range adj {
		sort.Ints(ints)
	}

	return adj
}

func dagSort(vertexList [][]int) []int {
	stack := make([]int, 0, len(vertexList))
	color := make([]Color, len(vertexList))
	dagVertexStack := make([]int, 0, len(vertexList))

	for i := len(vertexList) - 1; i > 0; i-- {
		if color[i] == WHITE {
			stack = append(stack, i)
		}

		for len(stack) != 0 {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if color[v] == WHITE {
				color[v] = GREY

				stack = append(stack, v)

				for _, vertex := range vertexList[v] {
					if color[vertex] != 0 {
						continue
					}

					stack = append(stack, vertex)
				}
			} else if color[v] == GREY {
				color[v] = BLACK
				dagVertexStack = append(dagVertexStack, v)
			}
		}
	}

	return dagVertexStack
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

	res := dagSort(vertexList)

	for i := len(res) - 1; i >= 0; i-- {
		fmt.Fprintf(writer, "%v ", res[i])
	}

	writer.Flush()
}
