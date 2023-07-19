package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Factorization(n int) (res []int) {
	d := 2

	for d*d <= n {
		if n%d == 0 {
			res = append(res, d)
			n = n / d
		} else {
			d += 1
		}
	}

	if n > 1 {
		res = append(res, n)
	}
	return
}

func main() {
	var num int
	fmt.Scan(&num)

	writer := bufio.NewWriter(os.Stdout)

	res := Factorization(num)

	for i := 0; i < len(res); i++ {
		writer.WriteString(strconv.Itoa(res[i]))

		if i == len(res)-1 {
			writer.WriteString("\n")
		} else {
			writer.WriteString(" ")
		}

	}

	writer.Flush()
}
