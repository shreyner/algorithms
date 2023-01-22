package palindrome

import (
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
