package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func QuadraticCalc(a, x, b, c int) int {
	return a*(x*x) + b*x + c
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)

	var a, x, b, c int
	scanner.Scan()
	line := scanner.Text()

	values := strings.Split(line, " ")
	a, _ = strconv.Atoi(values[0])
	x, _ = strconv.Atoi(values[1])
	b, _ = strconv.Atoi(values[2])
	c, _ = strconv.Atoi(values[3])

	result := QuadraticCalc(a, x, b, c)
	output_string := strconv.Itoa(result)
	writer.WriteString(output_string)
	writer.WriteString("\n")

	writer.Flush()
}
