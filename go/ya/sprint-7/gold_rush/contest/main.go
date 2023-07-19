package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type GoldenHeap struct {
	Weight int
	Cost   int
}

type GoldenHeaps []*GoldenHeap

func (s GoldenHeaps) Len() int {
	return len(s)
}

func (s GoldenHeaps) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s GoldenHeaps) Less(i, j int) bool {
	return s[i].Cost > s[j].Cost
}

func main() {
	var backpackCap int
	fmt.Scan(&backpackCap)

	var n int
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)

	goldenHeaps := make(GoldenHeaps, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		values := strings.Split(scanner.Text(), " ")

		cost, _ := strconv.Atoi(values[0])
		weight, _ := strconv.Atoi(values[1])

		goldenHeaps[i] = &GoldenHeap{
			Cost:   cost,
			Weight: weight,
		}
	}

	sort.Sort(goldenHeaps)

	maxCost := 0
	for backpackCap > 0 && n != 0 {
		goldenHeap := goldenHeaps[len(goldenHeaps)-n]

		weight := 0
		if backpackCap-goldenHeap.Weight <= 0 {
			weight = backpackCap
		} else {
			weight = goldenHeap.Weight
		}

		maxCost += weight * goldenHeap.Cost
		backpackCap -= weight

		n--
	}

	fmt.Fprintln(writer, maxCost)
	writer.Flush()
}
