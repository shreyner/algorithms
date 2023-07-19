package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ListForm(first []string, k int) (res []string) {
	x, _ := strconv.Atoi(strings.Join(first, ""))

	y := x + k

	return strings.Split(strconv.Itoa(y), "")
}

func main() {
	var numLength int
	fmt.Scan(&numLength)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)

	scanner.Scan()
	line := scanner.Text()
	nums := strings.Split(line, " ")

	scanner.Scan()
	lineX := scanner.Text()
	x, _ := strconv.Atoi(lineX)

	res := ListForm(nums, x)

	for i := 0; i < len(res); i++ {
		writer.WriteString(res[i])

		if i == len(res)-1 {
			writer.WriteString("\n")
		} else {
			writer.WriteString(" ")
		}

	}

	writer.Flush()
}
