package main

import "fmt"

// Comment it before submitting
type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode) {
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	Solution(&node0)
	/*
	   Output is:
	   node0
	   node1
	   node2
	   node3
	*/
}
