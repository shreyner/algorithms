package main

type Example struct {
	name string
}

func main() {
	// Empty check
	//d := deque.NewDeque(5)
	//fmt.Println(d.IsEmpty())
	//
	//_ = d.PushBack(10)
	//n, _ := d.PopFront()
	//fmt.Println(n, d.IsEmpty())
	//
	//// Re circled
	//d = deque.NewDeque(5)
	//
	//_ = d.PushBack(11)
	//_ = d.PushBack(12)
	//_ = d.PushBack(13)
	//_ = d.PushBack(14)
	//_ = d.PushBack(15)
	//
	//_, _ = d.PopFront()
	//_, _ = d.PopFront()
	//
	//_ = d.PushBack(16)
	//_ = d.PushBack(17)
	//
	//fmt.Println(d.GetQueue())
	//
	//// Err is fulled
	//d = deque.NewDeque(1)
	//_ = d.PushBack(10)
	//err := d.PushBack(10)
	//fmt.Println("Need error:", err)
	//
	//err = d.PushBack(10)
	//fmt.Println("Need error:", err)
	//
	//n, err = d.PopFront()
	//fmt.Println("no need error:", n, err)
	//
	//err = d.PushBack(10)
	//fmt.Println("no need error:", err)
	//err = d.PushBack(10)
	//fmt.Println("Need error:", err)
	//
	//// Err is empty
	//d = deque.NewDeque(1)
	//_ = d.PushBack(10)
	//n, err = d.PopFront()
	//fmt.Println("No need error:", n, err)
	//
	//_, err = d.PopFront()
	//fmt.Println("need error:", err)
	//_, err = d.PopFront()
	//fmt.Println("need error:", err)
	//
	//_ = d.PushBack(10)
	//n, err = d.PopFront()
	//fmt.Println("No need error:", n, err)
	//_, err = d.PopFront()
	//fmt.Println("Need error:", err)

	//d := deque.NewDeque(5)
	//_ = d.PushFront(10)
	//n, err := d.PopFront()
	//fmt.Println(n, err, d.GetQueue())

	//
	//d := deque.NewDeque(5)
	//_ = d.PushFront(10)
	//_ = d.PushFront(11)
	//_ = d.PushFront(12)
	//_ = d.PushFront(13)
	//_ = d.PushFront(14)
	//fmt.Println(d.GetQueue())
	//n, _ := d.PopBack()
	//fmt.Println(n)
	//n, _ = d.PopBack()
	//fmt.Println(n)
	//fmt.Println(d.GetQueue())
	//_ = d.PushFront(15)
	//_ = d.PushFront(16)
	//fmt.Println(d.GetQueue())
	//
	//n, _ = d.PopFront()
	//fmt.Println(n)
	//n, _ = d.PopFront()
	//fmt.Println(n)
	//fmt.Println(d.GetQueue())

	//fmt.Println(correct_bracket_seq.IsCorrectBracketSeq([]rune("{[()]}")))
	//fmt.Println(correct_bracket_seq.IsCorrectBracketSeq([]rune("()()")))
	//fmt.Println(correct_bracket_seq.IsCorrectBracketSeq([]rune("{}{}()")))
	//fmt.Println(correct_bracket_seq.IsCorrectBracketSeq([]rune("{}{}()[][]")))
	//fmt.Println(correct_bracket_seq.IsCorrectBracketSeq([]rune("")))
	//
	//fmt.Println("=====")
	//
	//fmt.Println(correct_bracket_seq.IsCorrectBracketSeq([]rune("{[()}")))
	//fmt.Println(correct_bracket_seq.IsCorrectBracketSeq([]rune("{")))
	//fmt.Println(correct_bracket_seq.IsCorrectBracketSeq([]rune("}")))

	//stackMax := stack_max_effective.NewStack()
	//stackMax.Push(1)
	//stackMax.Push(10)
	//stackMax.Push(100)
	//stackMax.Push(99)
	//stackMax.Push(50)
	//stackMax.Push(1000)
	//_, _ = stackMax.Pop()
	//println(stackMax.GetMax())

	//node3 := list_reverse.NewListNode("node3", nil, nil)
	//node2 := list_reverse.NewListNode("node2", node3, nil)
	//node1 := list_reverse.NewListNode("node1", node2, nil)
	//node0 := list_reverse.NewListNode("node0", node1, nil)
	//node3.SetPrev(node2)
	//node2.SetPrev(node1)
	//node1.SetPrev(node0)
	//newHead := list_reverse.Solution(node0)
	//list_reverse.PrintSolution(newHead)

	//node3 := caring_mother.NewListNode("node3", nil)
	//node2 := caring_mother.NewListNode("node2", node3)
	//node1 := caring_mother.NewListNode("node1", node2)
	//node0 := caring_mother.NewListNode("node0", node1)
	//fmt.Println(caring_mother.Solution(node0, "node222"))
	//fmt.Println(caring_mother.Solution(node0, "node0"))
	//fmt.Println(caring_mother.Solution(node0, "node1"))
	//fmt.Println(caring_mother.Solution(node0, "node2"))
	//fmt.Println(caring_mother.Solution(node0, "node3"))
	//fmt.Println(caring_mother.Solution(node0, "node4"))
	// result is : idx == 2

	//node3 := unloved_business.NewListNode("node3", nil)
	//node2 := unloved_business.NewListNode("node2", node3)
	//node1 := unloved_business.NewListNode("node1", node2)
	//node0 := unloved_business.NewListNode("node0", node1)
	//newHead := unloved_business.Solution(node0, 3)
	//
	//unloved_business.PrintSolution(newHead)
	// result is : node0 -> node2 -> node3

	//fmt.Println(sleight_of_hand.SleightOfHand([]int{1, 2, 3, 1, 2, -1, -1, 2, 2, -1, -1, 2, 2, -1, -1, 2}, 1))
	//fmt.Println(sleight_of_hand.SleightOfHand([]int{1, 1, 1, 1, 9, 9, 9, 9, 1, 1, 1, 1, 9, 9, 1, 1}, 4))
	//fmt.Println(sleight_of_hand.SleightOfHand([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 4))

	//fmt.Println(nearest_zero.NearestZero([]int{0, 1, 4, 9, 0}))
	//fmt.Println(nearest_zero.NearestZero([]int{0, 7, 9, 4, 8, 20}))
	//fmt.Println(nearest_zero.NearestZero([]int{3, 1, 0, 2, 4}))
	//fmt.Println(nearest_zero.NearestZero([]int{3, 1, 0, 0, 2, 4}))
	//fmt.Println(nearest_zero.NearestZero([]int{1, 0, 2}))
	//
	//fmt.Println(nearest_zero.NearestZero([]int{}))
	//fmt.Println(nearest_zero.NearestZero([]int{0}))

	//fmt.Println(factorization.Factorization(525))
	//
	//fmt.Println(find_pow_four.FindPowFour(4096))
	//
	//fmt.Println(add_ox.AddOx("1010", "1011")) // 10101
	//fmt.Println(add_ox.AddOx("1000", "1011")) // 10011
	//fmt.Println(add_ox.AddOx("10", "0"))      // 10
	//fmt.Println(add_ox.AddOx("0", "0"))       // 10
	//
	//fmt.Println(convert_to_ox.ConvertDecToOx(14))
	//
	//fmt.Println(palindrome.Palindrome("A man, a plan, a canal: Panama"))
	//fmt.Println(palindrome.Palindrome("zo"))
	//fmt.Println(palindrome.Palindrome("zoz"))
	//
	//fmt.Println(max_len_words.MaxLenWords(strings.Split("i love segment tree", " ")))   // segment, 7
	//fmt.Println(max_len_words.MaxLenWords(strings.Split("frog jumps from river", " "))) // jumps, 5
	//fmt.Println(max_len_words.MaxLenWords([]string{""}))                                // "", 0
	//
	//fmt.Println(random_weather.RandomWeather([]int{-1, -10, -8, 0, 2, 0, 5})) // 3
	//fmt.Println(random_weather.RandomWeather([]int{1, 2, 5, 4, 8}))           // 2
	//fmt.Println(random_weather.RandomWeather([]int{}))                        // 0
	//fmt.Println(random_weather.RandomWeather([]int{1}))                       // 1
	//
	//fmt.Println(neighbors.Neighbors([][]int{{1, 2, 3}, {0, 2, 6}, {7, 4, 1}, {2, 7, 0}}, []int{0, 0})) // [0 2]
	//fmt.Println(neighbors.Neighbors([][]int{{1, 2, 3}, {0, 2, 6}, {7, 4, 1}, {2, 7, 0}}, []int{3, 0})) // [7 7]
	//
	//fmt.Println(neighbors.Neighbors([][]int{{1, 2, 3}, {0, 2, 6}, {7, 4, 1}, {2, 7, 0}}, []int{0, 2})) // [6 2]
	//fmt.Println(neighbors.Neighbors([][]int{{1, 2, 3}, {0, 2, 6}, {7, 4, 1}, {2, 7, 0}}, []int{3, 2})) // [7 1]
	//
	//fmt.Println(
	//	neighbors.Neighbors(
	//		[][]int{
	//			{4, -10, 4, -9, 9, 5, -7, 1, 4, -3},
	//			{-3, 0, -1, -6, -6, 2, 3, 3, 4, 0},
	//			{-1, -5, 1, -9, -9, -6, 3, -1, -10, -7},
	//		},
	//		[]int{1, 0},
	//	),
	//) // [-1 0 4]
	//
	////fmt.Println(neighbors.Neighbors())
	//
	//result_sum := sum.Sum(1, 2)
	//timeNow := time.Now()
	//
	//fmt.Printf("Sum 1 + 2 = %v\n", result_sum)
	//fmt.Println("The time is", timeNow)
	//
	//myList := list.New()
	//
	//myList.PushBack(1)
	//myList.PushBack(2)
	//myList.PushBack(3)
	//myList.PushBack(4)
	//
	//element := myList.Front()
	//for i := 1; i < myList.Len(); i++ {
	//	fmt.Printf("Index: %v Value: %v\n", i, element.Value)
	//	element = element.Next()
	//}
	//
	//q := deque.New()
	//
	//q.PushBack(Example{name: "Hello"})
	//q.PushBack(Example{name: "Hello 2"})
	//q.PushBack(Example{name: "Hello 3"})
	//
	//for q.Len() > 0 {
	//	e := q.PopFront()
	//
	//	fmt.Printf("value: %v\n", e)
	//}
}
