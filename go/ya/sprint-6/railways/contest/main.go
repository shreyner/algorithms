package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Color int

const (
	WHITE Color = iota
	GREY
	BLACK
)

func dfs(vertexList [][]int) bool {
	stack := make([]int, 0, len(vertexList))
	colors := make([]Color, len(vertexList))

	for i := 0; i < len(vertexList); i++ {
		stack = append(stack, i)

		for len(stack) != 0 {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if colors[v] == WHITE {
				colors[v] = GREY

				stack = append(stack, v)

				for _, vertex := range vertexList[v] {
					if colors[vertex] == WHITE {
						stack = append(stack, vertex)
					} else if colors[vertex] == GREY {
						return true
					}
				}
			} else if colors[v] == GREY {
				colors[v] = BLACK
			}
		}
	}

	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	vertexList := make([][]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		values := strings.Split(scanner.Text(), "")

		for j, value := range values {
			if value == "B" {
				vertexList[i] = append(vertexList[i], i+j+1)
			}

			if value == "R" {
				vertexList[i+j+1] = append(vertexList[i+j+1], i)
			}
		}
	}

	res := dfs(vertexList)

	if res {
		fmt.Fprintln(writer, "NO")
	} else {
		fmt.Fprintln(writer, "YES")
	}

	writer.Flush()
}
