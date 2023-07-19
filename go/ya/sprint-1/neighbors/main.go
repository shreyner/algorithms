package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Neighbors(matrix [][]int, coordinates []int) (result []int) {
	x, y := coordinates[0], coordinates[1]

	if x+1 < len(matrix) {
		result = append(result, matrix[x+1][y])
	}

	if y-1 >= 0 {
		result = append(result, matrix[x][y-1])
	}

	if y+1 < len(matrix[x]) {
		result = append(result, matrix[x][y+1])
	}

	if x-1 >= 0 {
		result = append(result, matrix[x-1][y])
	}

	sort.Ints(result)

	return
}

func main() {
	var numLines, numRows int
	fmt.Scan(&numLines)
	fmt.Scan(&numRows)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)

	matrix := make([][]int, 0, numLines)

	for i := 1; i <= numLines; i++ {
		row := make([]int, 0, numRows)

		scanner.Scan()
		line := scanner.Text()

		values := strings.Split(line, " ")

		for j := 0; j < numRows; j++ {
			value, _ := strconv.Atoi(values[j])
			row = append(row, value)
		}

		matrix = append(matrix, row)
	}

	coordinate := make([]int, 2)

	for i := 0; i < 2; i++ {
		scanner.Scan()
		line := scanner.Text()

		value, _ := strconv.Atoi(line)
		coordinate[i] = value
	}

	result := Neighbors(matrix, coordinate)

	result2 := make([]string, len(result))

	for i, v := range result {
		result2[i] = strconv.Itoa(v)
	}

	writer.WriteString(strings.Join(result2, " "))
	writer.WriteString("\n")
	writer.Flush()
}
