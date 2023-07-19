package deque_old

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
//func (d *Deque) incIndex(i int) int {
//	return (i + 1) % d.max
//}
//
//func (d *Deque) decIndex(i int) int {
//	i -= 1
//	if i < 0 {
//		return d.max - 1
//	}
//
//	return i
//}
//
//func (d *Deque) PushBack(n int) error {
//	if d.size == d.max {
//		return errors.New("deque is fulled")
//	}
//
//	d.queue[d.tail] = n
//
//	d.tail = d.incIndex(d.tail)
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
//	d.head = d.decIndex(d.head)
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
//	d.head = d.incIndex(d.head)
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
//	d.tail = d.decIndex(d.tail)
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
//		if "push_back" == command ||  "push_front" == command {
//			if deque.IsFull() {
//				fmt.Fprintln(writer, "error")
//				continue
//			}
//
//			n, _ := strconv.Atoi(values[1])
//
//			switch command {
//			case "push_back":
//				deque.PushBack(n)
//			case "push_front":
//				deque.PushFront(n)
//			}
//		}
//
//		if "pop_front" == command || "pop_back" == command {
//			if deque.IsEmpty() {
//				fmt.Fprintln(writer, "error")
//				continue
//			}
//
//			var n int
//
//			switch command {
//			case "pop_front":
//				n, _ = deque.PopFront()
//			case "pop_back":
//				n, _ = deque.PopBack()
//			}
//
//			fmt.Fprintln(writer, n)
//		}
//	}
//
//	writer.Flush()
//}
