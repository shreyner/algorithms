package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const BufferSize = 100000000

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	scanner.Buffer(make([]byte, BufferSize), BufferSize)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	var strABuilder strings.Builder
	for _, charA := range scanner.Text() {
		if charA%2 == 0 {
			strABuilder.WriteRune(charA)
		}
	}

	scanner.Scan()
	var strBBuilder strings.Builder
	for _, charB := range scanner.Text() {
		if charB%2 == 0 {
			strBBuilder.WriteRune(charB)
		}
	}

	strA := strABuilder.String()
	strB := strBBuilder.String()

	if strA == strB {
		fmt.Fprintln(writer, "0")
		return
	}

	if strA < strB {
		fmt.Fprintln(writer, "-1")
		return
	}

	if strA > strB {
		fmt.Fprintln(writer, "1")
		return
	}
}
