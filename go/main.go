package main

import (
	"container/list"
	"fmt"
	"strings"
	"time"

	"github.com/gammazero/deque"

	"algorithms/sum"
	"algorithms/ya/add_ox"
	"algorithms/ya/convert_to_ox"
	"algorithms/ya/factorization"
	"algorithms/ya/find_pow_four"
	"algorithms/ya/max_len_words"
	"algorithms/ya/neighbors"
	"algorithms/ya/palindrome"
	"algorithms/ya/random_weather"
)

type Example struct {
	name string
}

func main() {
	fmt.Println(factorization.Factorization(525))

	fmt.Println(find_pow_four.FindPowFour(4096))

	fmt.Println(add_ox.AddOx("1010", "1011")) // 10101
	fmt.Println(add_ox.AddOx("1000", "1011")) // 10011
	fmt.Println(add_ox.AddOx("10", "0"))      // 10
	fmt.Println(add_ox.AddOx("0", "0"))       // 10

	fmt.Println(convert_to_ox.ConvertDecToOx(14))

	fmt.Println(palindrome.Palindrome("A man, a plan, a canal: Panama"))
	fmt.Println(palindrome.Palindrome("zo"))
	fmt.Println(palindrome.Palindrome("zoz"))

	fmt.Println(max_len_words.MaxLenWords(strings.Split("i love segment tree", " ")))   // segment, 7
	fmt.Println(max_len_words.MaxLenWords(strings.Split("frog jumps from river", " "))) // jumps, 5
	fmt.Println(max_len_words.MaxLenWords([]string{""}))                                // "", 0

	fmt.Println(random_weather.RandomWeather([]int{-1, -10, -8, 0, 2, 0, 5})) // 3
	fmt.Println(random_weather.RandomWeather([]int{1, 2, 5, 4, 8}))           // 2
	fmt.Println(random_weather.RandomWeather([]int{}))                        // 0
	fmt.Println(random_weather.RandomWeather([]int{1}))                       // 1

	fmt.Println(neighbors.Neighbors([][]int{{1, 2, 3}, {0, 2, 6}, {7, 4, 1}, {2, 7, 0}}, []int{0, 0})) // [0 2]
	fmt.Println(neighbors.Neighbors([][]int{{1, 2, 3}, {0, 2, 6}, {7, 4, 1}, {2, 7, 0}}, []int{3, 0})) // [7 7]

	fmt.Println(neighbors.Neighbors([][]int{{1, 2, 3}, {0, 2, 6}, {7, 4, 1}, {2, 7, 0}}, []int{0, 2})) // [6 2]
	fmt.Println(neighbors.Neighbors([][]int{{1, 2, 3}, {0, 2, 6}, {7, 4, 1}, {2, 7, 0}}, []int{3, 2})) // [7 1]

	fmt.Println(
		neighbors.Neighbors(
			[][]int{
				{4, -10, 4, -9, 9, 5, -7, 1, 4, -3},
				{-3, 0, -1, -6, -6, 2, 3, 3, 4, 0},
				{-1, -5, 1, -9, -9, -6, 3, -1, -10, -7},
			},
			[]int{1, 0},
		),
	) // [-1 0 4]

	//fmt.Println(neighbors.Neighbors())

	result_sum := sum.Sum(1, 2)
	timeNow := time.Now()

	fmt.Printf("Sum 1 + 2 = %v\n", result_sum)
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
