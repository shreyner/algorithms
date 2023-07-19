package main

import (
	"bufio"
	"fmt"
	"os"
)

func FindPowFour(num int) bool {
	for num > 1 {
		if num%4 != 0 {
			return false
		}

		num = num / 4
	}

	return true
}

func main() {
	var num int
	fmt.Scan(&num)

	writer := bufio.NewWriter(os.Stdout)

	res := FindPowFour(num)

	if res {
		writer.WriteString("True")
	} else {
		writer.WriteString("False")
	}

	writer.WriteString("\n")

	writer.Flush()
}
