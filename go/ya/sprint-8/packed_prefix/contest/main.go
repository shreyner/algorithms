package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func unpack(str string) []rune {
	numbersStack := make([]int, 0)
	strStack := make([][]rune, 0)

	strStack = append(strStack, make([]rune, 0))

	for _, char := range []rune(str) {
		if unicode.IsDigit(char) {
			numbersStack = append(numbersStack, int(char-'0'))
			continue
		}

		if char == '[' {
			strStack = append(strStack, make([]rune, 0))
			continue
		}

		if char == ']' {
			number := numbersStack[len(numbersStack)-1]
			chars := strStack[len(strStack)-1]

			numbersStack = numbersStack[:len(numbersStack)-1]
			strStack = strStack[:len(strStack)-1]

			for i := 0; i < number; i++ {
				for _, r := range chars {
					strStack[len(strStack)-1] = append(strStack[len(strStack)-1], r)
				}
			}

			continue
		}

		strStack[len(strStack)-1] = append(strStack[len(strStack)-1], char)
	}

	return strStack[len(strStack)-1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	word := unpack(scanner.Text())

	for i := 1; i < n; i++ {
		scanner.Scan()
		str := unpack(scanner.Text())
		lengthPrefix := 0

		for j := 0; j < len(word); j++ {
			if word[j] != str[j] || j > len(word) || j > len(str) {
				break
			}

			lengthPrefix += 1
		}

		word = word[:lengthPrefix]
	}

	fmt.Fprintln(writer, string(word))
}
