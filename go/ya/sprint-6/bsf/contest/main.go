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
	GRAY        = iota
	BLACK       = iota
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

func BFS(adjList [][]int, s int) []int {
	color := make([]Color, len(adjList))
	path := make([]int, 0)
	distance := make([]int, len(adjList))

	planned := make([]int, 0, len(adjList))

	planned = append(planned, s)
	color[s] = GRAY
	distance[s] = 0

	for len(planned) != 0 {
		u := planned[0]
		planned = planned[1:]
		path = append(path, u)

		for _, v := range adjList[u] {
			if color[v] != WHITE {
				continue
			}

			distance[v] = distance[u] + 1
			color[v] = GRAY
			planned = append(planned, v)
		}

		color[u] = BLACK
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

	adjList := convertToList(vCount, edges)

	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())

	path := BFS(adjList, s)

	for _, i := range path {
		fmt.Fprintf(writer, "%v ", i)
	}

	writer.Flush()
}
