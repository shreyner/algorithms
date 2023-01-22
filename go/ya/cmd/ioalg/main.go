package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)
	for i := 1; i <= n; i++ {
		var value_1, value_2 int
		scanner.Scan()
		line := scanner.Text()
		values := strings.Split(line, " ")
		value_1, _ = strconv.Atoi(values[0])
		value_2, _ = strconv.Atoi(values[1])
		result := value_1 + value_2
		output_string := strconv.Itoa(result)
		writer.WriteString(output_string)
		writer.WriteString("\n")
	}
	writer.Flush()
}
