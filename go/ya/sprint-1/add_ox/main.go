package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func PadLeft(str string, length int, pad string) string {
	var buf strings.Builder

	for i := 0; i < length-len(str); i++ {
		buf.WriteString(pad)
	}

	buf.WriteString(str)
	return buf.String()
}

func AddOx(a string, b string) string {
	var result []string

	var buf int

	max := Max(len(a), len(b))

	a = PadLeft(a, max, "0")
	b = PadLeft(b, max, "0")

	for i := len(a) - 1; i >= 0; i-- {
		aI, _ := strconv.Atoi(string(a[i]))
		bI, _ := strconv.Atoi(string(b[i]))

		v := buf + aI + bI
		buf = v / 2

		result = append(result, strconv.Itoa(v%2))
	}

	if buf != 0 {
		result = append(result, "1")
	}

	var sb strings.Builder

	for i := len(result) - 1; i >= 0; i-- {
		sb.WriteString(result[i])
	}

	return sb.String()
}

func main() {
	var numA, numB string
	fmt.Scan(&numA)
	fmt.Scan(&numB)

	writer := bufio.NewWriter(os.Stdout)

	output_string := AddOx(numA, numB)

	writer.WriteString(output_string)
	writer.WriteString("\n")

	writer.Flush()
}
