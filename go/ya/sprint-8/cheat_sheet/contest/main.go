package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const BufferSize = 1000000000

type Node struct {
	Symbol     rune
	IsTerminal bool
	Child      map[rune]*Node
}

func addString(root *Node, str string) {
	charArray := []rune(str)
	currentRoot := root

	for i := len(charArray) - 1; i >= 0; i-- {
		symbol := charArray[i]

		if currentRoot.Child[symbol] == nil {
			currentRoot.Child[symbol] = &Node{
				Symbol: symbol,
				Child:  make(map[rune]*Node),
			}
		}

		currentRoot = currentRoot.Child[symbol]
	}

	currentRoot.IsTerminal = true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	scanner.Buffer(make([]byte, BufferSize), BufferSize)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	arrayText := []rune(scanner.Text())

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	trieRoot := Node{
		Child: make(map[rune]*Node),
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		addString(&trieRoot, scanner.Text())
	}

	dp := make([]bool, len(arrayText)+1)
	dp[0] = true

	for i := 0; i < len(arrayText); i++ {
		currentNode := &trieRoot

		for j := i; j >= 0; j-- {
			currentNode = currentNode.Child[arrayText[j]]

			if currentNode == nil {
				break
			}

			if currentNode.IsTerminal && dp[j] == true {
				dp[i+1] = true
			}
		}
	}

	if dp[len(dp)-1] {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}
