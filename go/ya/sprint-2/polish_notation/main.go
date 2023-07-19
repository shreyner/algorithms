package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	stack []int
}

func NewStack() *Stack {
	return &Stack{
		stack: []int{},
	}
}

func (s *Stack) Len() int {
	return len(s.stack)
}

func (s *Stack) Push(x int) {
	s.stack = append(s.stack, x)
}

func (s *Stack) Pop() (int, error) {
	if len(s.stack) == 0 {
		return 0, errors.New("stack empty")
	}

	x := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return x, nil
}

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
	stack := NewStack()

	for _, v := range str {
		operator, ok := mapOperators[v]

		if !ok {
			n, _ := strconv.Atoi(v)
			stack.Push(n)

			continue
		}

		b, _ := stack.Pop()
		a, _ := stack.Pop()

		c := operator(a, b)

		stack.Push(c)
	}

	if stack.Len() == 0 {
		return 0
	}

	c, _ := stack.Pop()

	return c
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)

	scanner.Scan()
	line := scanner.Text()

	res := Calculate(line)

	fmt.Fprintln(writer, res)

	writer.Flush()
}
