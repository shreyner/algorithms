package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	StatusWIN  = "WIN"
	StatusFAIL = "FAIL"
)

func CheckWinEvent(a, b, c int) string {
	aE, bE, cE := a%2 == 0, b%2 == 0, c%2 == 0

	if aE == bE && bE == cE && cE == aE {
		return StatusWIN
	}

	return StatusFAIL
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)

	var a, b, c int
	scanner.Scan()
	line := scanner.Text()

	values := strings.Split(line, " ")
	a, _ = strconv.Atoi(values[0])
	b, _ = strconv.Atoi(values[1])
	c, _ = strconv.Atoi(values[2])

	result := CheckWinEvent(a, b, c)

	writer.WriteString(result)
	writer.WriteString("\n")

	writer.Flush()
}
