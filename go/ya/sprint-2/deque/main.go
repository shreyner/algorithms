package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deque struct {
	queue []int

	max  int
	size int

	head int
	tail int
}

func NewDeque(maxSize int) *Deque {
	return &Deque{
		queue: make([]int, maxSize),
		max:   maxSize,
	}
}

func (d *Deque) IsEmpty() bool {
	return d.size == 0
}

func (d *Deque) IsFull() bool {
	return d.size == d.max
}

func (d *Deque) incIndex(i int) int {
	return (i + 1) % d.max
}

func (d *Deque) decIndex(i int) int {
	i -= 1
	if i < 0 {
		return d.max - 1
	}

	return i
}

func (d *Deque) PushBack(n int) error {
	if d.size == d.max {
		return errors.New("deque is fulled")
	}

	d.queue[d.tail] = n

	d.tail = d.incIndex(d.tail)
	d.size += 1

	return nil
}

func (d *Deque) PushFront(n int) error {
	if d.size == d.max {
		return errors.New("deque is fulled")
	}

	d.head = d.decIndex(d.head)

	d.queue[d.head] = n

	d.size += 1

	return nil
}

func (d *Deque) PopFront() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque is empty")
	}

	n := d.queue[d.head]

	d.head = d.incIndex(d.head)
	d.size -= 1

	return n, nil
}

func (d *Deque) PopBack() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque is empty")
	}

	d.tail = d.decIndex(d.tail)

	n := d.queue[d.tail]

	d.size -= 1

	return n, nil
}

type ExecutePushFunc func(n int) error
type ExecutePopFunc func() (int, error)

func main() {
	var countCommand int
	fmt.Scan(&countCommand)
	var maxDeque int
	fmt.Scan(&maxDeque)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)

	deque := NewDeque(maxDeque)

	commandPushMap := map[string]ExecutePushFunc{
		"push_back":  deque.PushBack,
		"push_front": deque.PushFront,
	}

	commandPopMap := map[string]ExecutePopFunc{
		"pop_front": deque.PopFront,
		"pop_back":  deque.PopBack,
	}

	for i := 1; i <= countCommand; i++ {
		scanner.Scan()
		line := scanner.Text()

		values := strings.Split(line, " ")
		command := values[0]

		commandPushExecuter, ok := commandPushMap[command]
		if ok {
			n, _ := strconv.Atoi(values[1])

			if err := commandPushExecuter(n); err != nil {
				fmt.Fprintln(writer, "error")
			}

			continue
		}

		commandPopExecuter, ok := commandPopMap[command]
		if ok {
			n, err := commandPopExecuter()

			if err != nil {
				fmt.Fprintln(writer, "error")
				continue
			}

			fmt.Fprintln(writer, n)
			continue
		}
	}

	writer.Flush()
}
