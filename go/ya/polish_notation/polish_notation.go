package polish_notation

import (
	"math"
	"strconv"
	"strings"
)

type Operator func(a, b int) int

var mapOperators = map[string]Operator{
	"+": func(a, b int) int {
		return a + b
	},
	"-": func(a, b int) int {
		return a - b
	},
	"*": func(a, b int) int {
		return a * b
	},
	"/": func(a, b int) int {
		return int(math.Floor(float64(a) / float64(b)))
	},
}

func Calculate(s string) int {
	str := strings.Split(s, " ")
	stack := make([]int, 0)

	for _, v := range str {
		operator, ok := mapOperators[v]

		if !ok {
			n, _ := strconv.Atoi(v)
			stack = append(stack, n)

			continue
		}

		b := stack[len(stack)-1]
		a := stack[len(stack)-2]

		stack = stack[:len(stack)-2]

		c := operator(a, b)

		stack = append(stack, c)
	}

	if len(stack) == 0 {
		return 0
	}

	c := stack[len(stack)-1]

	return c
}

//package main
//
//import (
//"bufio"
//"fmt"
//"math"
//"os"
//"strconv"
//"strings"
//)
//
//type Operator func(a, b int) int
//
//var mapOperators = map[string]Operator{
//	"+": func(a, b int) int {
//		return a + b
//	},
//	"-": func(a, b int) int {
//		return a - b
//	},
//	"*": func(a, b int) int {
//		return a * b
//	},
//	"/": func(a, b int) int {
//		return int(math.Floor(float64(a) / float64(b)))
//	},
//}
//
//func Calculate(s string) int {
//	str := strings.Split(s, " ")
//	stack := make([]int, 0)
//
//	for _, v := range str {
//		operator, ok := mapOperators[v]
//
//		if !ok {
//			n, _ := strconv.Atoi(v)
//			stack = append(stack, n)
//
//			continue
//		}
//
//		b := stack[len(stack)-1]
//		a := stack[len(stack)-2]
//
//		stack = stack[:len(stack)-2]
//
//		c := operator(a, b)
//
//		stack = append(stack, c)
//	}
//
//	if len(stack) == 0 {
//		return 0
//	}
//
//	c := stack[len(stack)-1]
//
//	return c
//}
//
//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	scanner := bufio.NewScanner(reader)
//	scanner.Split(bufio.ScanLines)
//
//	writer := bufio.NewWriter(os.Stdout)
//
//	scanner.Scan()
//	line := scanner.Text()
//
//	res := Calculate(line)
//
//	fmt.Fprintln(writer, res)
//
//	writer.Flush()
//}
//
