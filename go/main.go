package main

import (
	"algorithms/sum"
	"container/list"
	"fmt"
	"time"

	"github.com/gammazero/deque"
)

type Example struct {
	name string
}

func main() {
	sum := sum.Sum(1, 2)
	timeNow := time.Now()

	fmt.Println("Hello world")
	fmt.Printf("Sum 1 + 2 = %v\n", sum)
	fmt.Println("The time is", timeNow)

	myList := list.New()

	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	myList.PushBack(4)

	element := myList.Front()
	for i := 1; i < myList.Len(); i++ {
		fmt.Printf("Index: %v Value: %v\n", i, element.Value)
		element = element.Next()
	}

	q := deque.New()

	q.PushBack(Example{name: "Hello"})
	q.PushBack(Example{name: "Hello 2"})
	q.PushBack(Example{name: "Hello 3"})

	for q.Len() > 0 {
		e := q.PopFront()

		fmt.Printf("value: %v\n", e)
	}
}
