package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StackMaxEffective struct {
	stack    []int
	maxStack []int
}

func NewStack() *StackMaxEffective {
	return &StackMaxEffective{
		stack: []int{},
	}
}

func (s *StackMaxEffective) Push(x int) {
	s.stack = append(s.stack, x)

	if len(s.maxStack) == 0 {
		s.maxStack = append(s.maxStack, x)
		return
	}

	max := s.maxStack[len(s.maxStack)-1]
	if x > max {
		max = x
	}

	s.maxStack = append(s.maxStack, max)
}

func (s *StackMaxEffective) Pop() (int, error) {
	if len(s.stack) == 0 {
		return 0, errors.New("stack empty")
	}

	x := s.stack[len(s.stack)-1]

	s.stack = s.stack[:len(s.stack)-1]
	s.maxStack = s.maxStack[:len(s.maxStack)-1]

	return x, nil
}

func (s *StackMaxEffective) GetMax() (int, error) {
	if len(s.stack) == 0 {
		return 0, errors.New("stack empty")
	}

	max := s.maxStack[len(s.maxStack)-1]

	return max, nil
}

func main() {
	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)

	stack := NewStack()

	for i := 1; i <= n; i++ {
		scanner.Scan()
		line := scanner.Text()

		values := strings.Split(line, " ")
		command := values[0]

		switch command {
		case "get_max":
			max, err := stack.GetMax()
			if err != nil {
				fmt.Fprintln(writer, "None")
				continue
			}
			fmt.Fprintln(writer, max)
		case "push":
			x, _ := strconv.Atoi(values[1])
			stack.Push(x)
		case "pop":
			_, err := stack.Pop()
			if err != nil {
				fmt.Fprintln(writer, "error")
				continue
			}
		}
	}

	writer.Flush()
}
