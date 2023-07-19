package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func MaxLenWords(words []string) (string, int) {
	//words := strings.Split(line, " ")

	maxLength := 0
	firstMaxLengthWord := ""

	for _, word := range words {
		count := utf8.RuneCountInString(word)

		if count > maxLength {
			maxLength = count
			firstMaxLengthWord = word
		}

	}

	return firstMaxLengthWord, maxLength
}

func main() {
	var numWords int
	fmt.Scan(&numWords)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)

	words := make([]string, 0)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	word, maxLength := MaxLenWords(words)

	fmt.Fprintf(writer, "%v\n%v\n", word, maxLength)
	writer.Flush()
}
