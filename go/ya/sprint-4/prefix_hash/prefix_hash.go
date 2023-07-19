package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PolynomialHash(arrStr string, a, m int) (sum int) {
	if len(arrStr) == 0 {
		return
	}

	for _, c := range arrStr {
		sum = (sum*a + int(c)) % m
	}

	return
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	str := scanner.Text()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		values := strings.Split(scanner.Text(), " ")

		var left, right int
		left, _ = strconv.Atoi(values[0])
		right, _ = strconv.Atoi(values[1])

		fmt.Println(PolynomialHash(str[left-1:right], a, m))
	}
}
