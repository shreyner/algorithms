package unloved_business

import "fmt"

// ListNode - Comment it before submitting
type ListNode struct {
	data string
	next *ListNode
}

func NewListNode(data string, node *ListNode) *ListNode {
	return &ListNode{
		data: data,
		next: node,
	}
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

func PrintSolution(head *ListNode) {
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}
