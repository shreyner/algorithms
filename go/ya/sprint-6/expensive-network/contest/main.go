package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge []int

func NewEdge(edge, weight int) Edge {
	return Edge{edge, weight}
}

func (e Edge) Edge() int {
	return e[0]
}

func (e Edge) Weight() int {
	return e[1]
}

type EdgesHeapSort []Edge

func (h EdgesHeapSort) Len() int {
	return len(h)
}

func (h EdgesHeapSort) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h EdgesHeapSort) Less(i, j int) bool {
	return h[i].Weight() > h[j].Weight()
}

func (h *EdgesHeapSort) Push(x any) {
	*h = append(*h, x.(Edge))
}

func (h *EdgesHeapSort) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

var (
	maxSpanningTree = 0
	added           = make(map[int]struct{}, 0)
	edgesSort       = make(EdgesHeapSort, 0)
)

func convertToList(vCount int, edges [][]int) [][]Edge {
	adjList := make([][]Edge, vCount+1)

	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]

		adjList[u] = append(adjList[u], NewEdge(v, w))
		adjList[v] = append(adjList[v], NewEdge(u, w))
	}

	return adjList
}

func addVertex(adjList [][]Edge, v int) {
	added[v] = struct{}{}

	for _, edge := range adjList[v] {
		if _, ok := added[edge[0]]; ok {
			continue
		}

		heap.Push(&edgesSort, edge)
	}
}

func findMST(adjList [][]Edge) error {
	v := 1

	addVertex(adjList, v)

	for len(added) != len(adjList)-1 && edgesSort.Len() != 0 {
		e := heap.Pop(&edgesSort).(Edge)

		if _, ok := added[e[0]]; !ok {
			maxSpanningTree += e.Weight()
			addVertex(adjList, e[0])
		}
	}

	if len(added) != len(adjList)-1 {
		return errors.New("Oops! I did it again")
	}

	return nil
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

	err := findMST(adjList)

	if err != nil {
		fmt.Fprintln(writer, err)
	} else {
		fmt.Fprintln(writer, maxSpanningTree)
	}

	writer.Flush()
}
