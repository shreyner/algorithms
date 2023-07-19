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
	Cost   int
	Weight int
}

func main() {
	var capacity int
	fmt.Scan(&capacity)

	var n int
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	goldenHeaps := make([]*GoldenHeap, n)

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

	sort.Slice(goldenHeaps, func(i, j int) bool {
		return goldenHeaps[i].Cost > goldenHeaps[j].Cost
	})

	maxCost := 0
	for i := 0; i < n && capacity > 0; i++ {
		weight := goldenHeaps[i].Weight
		if capacity < weight {
			weight = capacity
		}

		maxCost += weight * goldenHeaps[i].Cost
		capacity -= weight
	}

	fmt.Fprintln(writer, maxCost)
}
