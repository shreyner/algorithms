package main

import (
	"bufio"
	"fmt"
	"os"
)

func UniqCups(arr []string) (result []string) {
	mapSet := make(map[string]bool, len(arr))

	for _, s := range arr {
		if _, ok := mapSet[s]; ok {
			continue
		}

		mapSet[s] = true
		result = append(result, s)
	}

	return result
}

func main() {
	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)
	arr := make([]string, n)

	for i := 0; i < n; i++ {
		scanner.Scan()

		arr[i] = scanner.Text()
	}

	res := UniqCups(arr)

	for _, str := range res {
		fmt.Println(str)
	}

	writer.Flush()
}
