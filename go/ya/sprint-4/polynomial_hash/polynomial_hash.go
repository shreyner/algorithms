package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	fmt.Println(PolynomialHash(scanner.Text(), a, m))
}
