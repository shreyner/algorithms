package main

// ListNode - Comment it before submitting
type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode, idx int) *ListNode {
	if idx == 0 {
		return head.next
	}

	prevNode := head
	targetNode := head.next
	index := 1

	for index < idx {
		prevNode = targetNode
		targetNode = targetNode.next
		index += 1
	}

	prevNode.next = targetNode.next

	return head
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	/*newHead :=*/ Solution(&node0, 1)
	// result is : node0 -> node2 -> node3
}
