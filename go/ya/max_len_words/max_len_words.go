package max_len_words

import (
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
