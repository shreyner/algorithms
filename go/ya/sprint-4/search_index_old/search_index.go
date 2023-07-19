package search_index_old

import (
	"sort"
	"strings"
)

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
	sss := make([]int, len(documents))

	searchTokens := make(map[string]struct{})
	for _, word := range strings.Split(search, " ") {
		searchTokens[word] = struct{}{}
	}

	for word, _ := range searchTokens {
		for indexDocument, count := range searchIndex[word] {
			sss[indexDocument] += count
		}
	}

	result := make([]int, 0)
	for i, count := range sss {
		if count == 0 {
			continue
		}

		result = append(result, i)
	}

	sort.Slice(result, func(i, j int) bool {
		return sss[result[i]] > sss[result[j]]
	})

	if len(result) > 5 {
		return result[:5]
	}

	return result
}

//package main
//
//import (
//"bufio"
//"fmt"
//"os"
//"sort"
//"strconv"
//"strings"
//)
//
//func CreateSearchIndex(documents []string) map[string]map[int]int {
//	searchIndex := make(map[string]map[int]int)
//
//	for indexDocument, document := range documents {
//		for _, word := range strings.Split(document, " ") {
//			if _, ok := searchIndex[word]; !ok {
//				searchIndex[word] = make(map[int]int)
//			}
//
//			searchIndex[word][indexDocument] += 1
//		}
//	}
//
//	return searchIndex
//}
//
//func FindAllDocumentsByIndex(documents []string, searchIndex map[string]map[int]int, search string) []int {
//	sss := make([]int, len(documents))
//
//	searchTokens := make(map[string]struct{})
//	for _, word := range strings.Split(search, " ") {
//		searchTokens[word] = struct{}{}
//	}
//
//	for word, _ := range searchTokens {
//		for indexDocument, count := range searchIndex[word] {
//			sss[indexDocument] += count
//		}
//	}
//
//	result := make([]int, 0)
//	for i, count := range sss {
//		if count == 0 {
//			continue
//		}
//
//		result = append(result, i)
//	}
//
//	sort.SliceStable(result, func(i, j int) bool {
//		return sss[result[i]] > sss[result[j]]
//	})
//
//	if len(result) > 5 {
//		return result[:5]
//	}
//
//	return result
//}
//
//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	writer := bufio.NewWriter(os.Stdout)
//	scanner := bufio.NewScanner(reader)
//	scanner.Split(bufio.ScanLines)
//
//	scanner.Scan()
//	n, _ := strconv.Atoi(scanner.Text())
//
//	documents := make([]string, n)
//
//	for i := 0; i < n; i++ {
//		scanner.Scan()
//		documents[i] = scanner.Text()
//	}
//
//	searchIndex := CreateSearchIndex(documents)
//
//	scanner.Scan()
//	n, _ = strconv.Atoi(scanner.Text())
//
//	for i := 0; i < n; i++ {
//		scanner.Scan()
//		search := scanner.Text()
//
//		result := FindAllDocumentsByIndex(documents, searchIndex, search)
//
//		for _, indexDducment := range result {
//			fmt.Fprintf(writer, "%v ", indexDducment+1)
//		}
//		fmt.Fprintf(writer, "\n")
//	}
//
//	writer.Flush()
//}
