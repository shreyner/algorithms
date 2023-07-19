package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Login string
	Tasks int
	Fine  int
}

func userLessCompare(a, b *User) bool {
	if a.Tasks != b.Tasks {
		return a.Tasks < b.Tasks
	} else if a.Fine != b.Fine {
		return a.Fine > b.Fine
	} else if a.Login != b.Login {
		return a.Login > b.Login
	}

	return false
}

type LessComparator[T any] func(a, b T) bool

type Heapsort[T any] struct {
	size int
	heap []T
	less LessComparator[T]
}

func NewHeap[T any](cap int, less LessComparator[T]) *Heapsort[T] {
	return &Heapsort[T]{
		heap: make([]T, cap+1),
		less: less,
	}
}

func (h *Heapsort[T]) shiftUp(idx int) {
	if idx == 1 {
		return
	}

	parentIdx := idx / 2

	if h.less(h.heap[parentIdx], h.heap[idx]) {
		h.heap[parentIdx], h.heap[idx] = h.heap[idx], h.heap[parentIdx]

		h.shiftUp(parentIdx)
		return
	}
}

func (h *Heapsort[T]) shiftDown(idx int) {
	leftIdx := 2 * idx
	rightIds := 2*idx + 1

	if h.size < leftIdx {
		return
	}

	var indexLarge int
	if (rightIds <= h.size) && h.less(h.heap[leftIdx], h.heap[rightIds]) {
		indexLarge = rightIds
	} else {
		indexLarge = leftIdx
	}

	if h.less(h.heap[idx], h.heap[indexLarge]) {
		h.heap[indexLarge], h.heap[idx] = h.heap[idx], h.heap[indexLarge]

		h.shiftDown(indexLarge)
		return
	}
}

func (h *Heapsort[T]) Add(d T) {
	h.size += 1
	idx := h.size

	h.heap[idx] = d
	h.shiftUp(idx)
}

func (h *Heapsort[T]) Get() T {
	result := h.heap[1]
	h.heap[1] = h.heap[h.size]
	h.size -= 1
	h.shiftDown(1)
	return result
}

func main() {
	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)

	hs := NewHeap[*User](n, userLessCompare)

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")
		name, tasksStr, fineStr := line[0], line[1], line[2]

		tasks, _ := strconv.Atoi(tasksStr)
		fine, _ := strconv.Atoi(fineStr)

		user := &User{
			Login: name,
			Tasks: tasks,
			Fine:  fine,
		}

		hs.Add(user)
	}

	for i := 0; i < n; i++ {
		u := hs.Get()

		fmt.Fprintln(writer, u.Login)
	}

	writer.Flush()
}
