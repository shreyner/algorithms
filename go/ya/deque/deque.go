package deque

import (
	"errors"
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

func (d *Deque) PushBack(n int) error {
	if d.IsFull() {
		return errors.New("deque is fulled")
	}

	d.queue[d.tail] = n

	d.tail = (d.tail + 1) % d.max
	d.size += 1

	return nil
}

func (d *Deque) PushFront(n int) error {
	if d.IsFull() {
		return errors.New("deque is fulled")
	}

	d.head -= 1
	if d.head < 0 {
		d.head = d.max - 1
	}

	d.queue[d.head] = n

	d.size += 1

	return nil
}

func (d *Deque) PopFront() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque is empty")
	}

	n := d.queue[d.head]

	d.head = (d.head + 1) % d.max
	d.size -= 1

	return n, nil
}

func (d *Deque) PopBack() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque is empty")
	}

	d.tail -= 1
	if d.tail < 0 {
		d.tail = d.max - 1
	}

	n := d.queue[d.tail]

	d.size -= 1

	return n, nil
}

//package main
//
//import (
//"bufio"
//"errors"
//"fmt"
//"os"
//"strconv"
//"strings"
//)
//
//type Deque struct {
//	queue []int
//
//	max  int
//	size int
//
//	head int
//	tail int
//}
//
//func NewDeque(maxSize int) *Deque {
//	return &Deque{
//		queue: make([]int, maxSize),
//		max:   maxSize,
//	}
//}
//
//func (d *Deque) IsEmpty() bool {
//	return d.size == 0
//}
//
//func (d *Deque) IsFull() bool {
//	return d.size == d.max
//}
//
//func (d *Deque) PushBack(n int) error {
//	if d.size == d.max {
//		return errors.New("deque is fulled")
//	}
//
//	d.queue[d.tail] = n
//
//	d.tail = (d.tail + 1) % d.max
//	d.size += 1
//
//	return nil
//}
//
//func (d *Deque) PushFront(n int) error {
//	if d.size == d.max {
//		return errors.New("deque is fulled")
//	}
//
//	d.head -= 1
//	if d.head < 0 {
//		d.head = d.max - 1
//	}
//
//	d.queue[d.head] = n
//
//	d.size += 1
//
//	return nil
//}
//
//func (d *Deque) PopFront() (int, error) {
//	if d.IsEmpty() {
//		return 0, errors.New("deque is empty")
//	}
//
//	n := d.queue[d.head]
//
//	d.head = (d.head + 1) % d.max
//	d.size -= 1
//
//	return n, nil
//}
//
//func (d *Deque) PopBack() (int, error) {
//	if d.IsEmpty() {
//		return 0, errors.New("deque is empty")
//	}
//
//	d.tail -= 1
//	if d.tail < 0 {
//		d.tail = d.max - 1
//	}
//
//	n := d.queue[d.tail]
//
//	d.size -= 1
//
//	return n, nil
//}
//
//func main() {
//	var countCommand int
//	fmt.Scan(&countCommand)
//	var maxDeque int
//	fmt.Scan(&maxDeque)
//
//	reader := bufio.NewReader(os.Stdin)
//	scanner := bufio.NewScanner(reader)
//	scanner.Split(bufio.ScanLines)
//
//	writer := bufio.NewWriter(os.Stdout)
//
//	deque := NewDeque(maxDeque)
//
//	for i := 1; i <= countCommand; i++ {
//		scanner.Scan()
//		line := scanner.Text()
//
//		values := strings.Split(line, " ")
//		command := values[0]
//
//		switch command {
//		case "push_back":
//			if deque.IsFull() {
//				fmt.Fprintln(writer, "error")
//				continue
//			}
//
//			n, _ := strconv.Atoi(values[1])
//			deque.PushBack(n)
//		case "push_front":
//			if deque.IsFull() {
//				fmt.Fprintln(writer, "error")
//				continue
//			}
//			n, _ := strconv.Atoi(values[1])
//			deque.PushFront(n)
//		case "pop_front":
//			if deque.IsEmpty() {
//				fmt.Fprintln(writer, "error")
//				continue
//			}
//
//			n, _ := deque.PopFront()
//			fmt.Fprintln(writer, n)
//		case "pop_back":
//			if deque.IsEmpty() {
//				fmt.Fprintln(writer, "error")
//				continue
//			}
//
//			n, _ := deque.PopBack()
//			fmt.Fprintln(writer, n)
//		}
//	}
//
//	writer.Flush()
//}
