package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`[\W]`)

func Palindrome(line string) bool {
	line = strings.ToLower(re.ReplaceAllString(line, ""))

	for i := 0; i < len(line)/2; i++ {
		if line[i] != line[len(line)-i-1] {
			return false
		}

	}

	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)

	scanner.Scan()
	line := scanner.Text()

	result := Palindrome(line)

	if result {
		writer.WriteString("True")
	} else {
		writer.WriteString("False")
	}

	writer.Flush()
}
