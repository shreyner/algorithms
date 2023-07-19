package main

// ListNode Comment it before submitting
type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode, elem string) int {
	index := 0
	for head != nil {
		if head.data == elem {
			return index
		}

		head = head.next
		index += 1
	}

	return -1
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	/*idx :=*/ Solution(&node0, "node2")
	// result is : idx == 2
}
