package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const formatString = "%s "

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	words := strings.Split(scanner.Text(), " ")

	for i := len(words) - 1; i >= 0; i-- {
		fmt.Fprintf(writer, formatString, words[i])
	}
}
