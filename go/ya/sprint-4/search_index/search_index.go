package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var resultLimit = 5

func CreateSearchIndex(documents []string) map[string]map[int]int {
	searchIndex := make(map[string]map[int]int)

	for indexDocument, document := range documents {
		for _, word := range strings.Split(document, " ") {
			if _, ok := searchIndex[word]; !ok {
				searchIndex[word] = make(map[int]int)
			}

			searchIndex[word][indexDocument] += 1
		}
	}

	return searchIndex
}

func FindAllDocumentsByIndex(documents []string, searchIndex map[string]map[int]int, search string) []int {
	relatedDocuments := make(map[int]int, len(documents))

	searchWords := make(map[string]struct{})
	for _, word := range strings.Split(search, " ") {
		searchWords[word] = struct{}{}
	}

	for word, _ := range searchWords {
		for indexDocument, count := range searchIndex[word] {
			relatedDocuments[indexDocument] += count
		}
	}

	result := make([]int, 0, resultLimit)
	for i := 0; i < resultLimit; i++ {
		maxIndexDocument := -1
		maxSumRelated := 0

		for indexDocument, sumRelated := range relatedDocuments {
			if maxSumRelated < sumRelated || (maxSumRelated == sumRelated && (maxIndexDocument == -1 || maxIndexDocument > indexDocument)) {
				maxSumRelated = sumRelated
				maxIndexDocument = indexDocument
			}
		}

		if maxIndexDocument == -1 {
			break
		}

		delete(relatedDocuments, maxIndexDocument)
		result = append(result, maxIndexDocument)
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	documents := make([]string, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		documents[i] = scanner.Text()
	}

	searchIndex := CreateSearchIndex(documents)

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		search := scanner.Text()

		result := FindAllDocumentsByIndex(documents, searchIndex, search)

		for _, indexDducment := range result {
			fmt.Fprintf(writer, "%v ", indexDducment+1)
		}
		fmt.Fprintf(writer, "\n")
	}

	writer.Flush()
}
