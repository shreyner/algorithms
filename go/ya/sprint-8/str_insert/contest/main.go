package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const BufferSize = 100000000

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	scanner.Buffer(make([]byte, BufferSize), BufferSize)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	text := scanner.Text()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	mapWords := make(map[int]string, n)
	keys := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		values := strings.Split(scanner.Text(), " ")

		index, _ := strconv.Atoi(values[1])
		mapWords[index] = values[0]
		keys[i] = index
	}

	sort.Ints(keys)

	start := 0
	for i := 0; i < n; i++ {
		pos := keys[i]
		fmt.Fprint(writer, text[start:pos])
		fmt.Fprint(writer, mapWords[pos])

		start = pos
	}

	if start < len(text) {
		fmt.Fprint(writer, text[start:])
	}
}
