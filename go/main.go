package main

import (
	"algorithms/sum"
	"fmt"
	"time"
)

func main() {
	sum := sum.Sum(1, 2)
	timeNow := time.Now()

	fmt.Println("Hello world")
	fmt.Printf("Sum 1 + 2 = %v\n", sum)
	fmt.Println("The time is", timeNow)
}
