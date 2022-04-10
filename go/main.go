package main

import (
	"algorithms/sum"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world")
	fmt.Printf("Sum 1 + 2 = %v\n", sum.Sum(1, 2))
	fmt.Println("The time is", time.Now())
}
