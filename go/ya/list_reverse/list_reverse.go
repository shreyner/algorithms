package list_reverse

import "fmt"

// ListNode - Comment it before submitting
type ListNode struct {
	data string
	next *ListNode
	prev *ListNode
}

func (n *ListNode) SetPrev(prev *ListNode) {
	n.prev = prev
}

func NewListNode(data string, next *ListNode, prev *ListNode) *ListNode {
	return &ListNode{
		data: data,
		next: next,
		prev: prev,
	}
}

func Solution(head *ListNode) *ListNode {
	prevNode := head
	for head != nil {
		nextNode := head.next

		head.next = head.prev
		head.prev = nextNode

		prevNode = head
		head = nextNode
	}

	return prevNode
}

func PrintSolution(head *ListNode) {
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}
