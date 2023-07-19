package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ConvertDecToOx(num int) string {
	if num == 0 {
		return "0"
	}

	var sb []int
	lastNum := num

	for lastNum > 0 {
		sb = append(sb, lastNum%2)
		lastNum = lastNum / 2
	}

	var reverseSb strings.Builder
	for i := len(sb) - 1; 0 <= i; i-- {
		reverseSb.WriteString(strconv.Itoa(sb[i]))
	}

	return reverseSb.String()
}

func main() {
	var num int
	fmt.Scan(&num)

	writer := bufio.NewWriter(os.Stdout)

	output_string := ConvertDecToOx(num)

	writer.WriteString(output_string)
	writer.WriteString("\n")

	writer.Flush()
}
