package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func convertToList(vCount int, edges [][]int) [][][]int {
	adjList := make([][][]int, vCount+1)

	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]

		adjList[u] = append(adjList[u], []int{v, w})
		adjList[v] = append(adjList[v], []int{u, w})
	}

	return adjList
}

func weight(adjList [][][]int, t, u int) int {
	for _, i := range adjList[t] {
		if i[0] != u {
			continue
		}

		return i[1]
	}

	return math.MaxInt
}

func relax(adjList [][][]int, dist []int, previous []int, t, u int) {
	if dist[u] > dist[t]+weight(adjList, t, u) {
		dist[u] = dist[t] + weight(adjList, t, u)
		previous[u] = t
	}
}

func getMinDistNotVisitedVertex(adjList [][][]int, dist []int, visited []bool) int {
	currentMinimum := math.MaxInt
	currentMinimumVertex := -1

	for i := 1; i < len(adjList); i++ {
		if !visited[i] && dist[i] < currentMinimum {
			currentMinimum = dist[i]
			currentMinimumVertex = i
		}
	}

	return currentMinimumVertex

}

func dijkstra(adjList [][][]int, s int) []int {
	dist := make([]int, len(adjList))
	previous := make([]int, len(adjList))
	visited := make([]bool, len(adjList))

	for v := range adjList {
		dist[v] = math.MaxInt
		previous[v] = -1
		visited[v] = false
	}

	dist[s] = 0

	for {
		t := getMinDistNotVisitedVertex(adjList, dist, visited)

		if t == -1 || dist[t] == math.MaxInt {
			break
		}

		visited[t] = true

		neighbours := adjList[t]

		for _, neighbour := range neighbours {
			relax(adjList, dist, previous, t, neighbour[0])
		}
	}

	return dist
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
		w, _ := strconv.Atoi(line[2])

		edges[i] = []int{u, v, w}
	}

	adjList := convertToList(vCount, edges)

	for i := 1; i < vCount+1; i++ {
		previous := dijkstra(adjList, i)

		for _, dist := range previous[1:] {
			if dist == math.MaxInt {
				dist = -1
			}

			fmt.Fprintf(writer, "%v ", dist)
		}

		fmt.Fprint(writer, "\n")
	}

	writer.Flush()
}
