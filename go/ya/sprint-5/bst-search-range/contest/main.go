package main

import "fmt"

// Comment it before submitting
type Node struct {
	value int
	left  *Node
	right *Node
}

func printRange(root *Node, left, right int) {
	if root.left != nil && left <= root.value {
		printRange(root.left, left, right)
	}

	if left <= root.value && right >= root.value {
		fmt.Printf("%v ", root.value)
	}

	if root.right != nil && right >= root.value {
		printRange(root.right, left, right)
	}
}

func test() {
	node0 := Node{0, nil, nil}
	node1 := Node{2, nil, nil}
	node2 := Node{1, &node0, &node1}
	node3 := Node{8, nil, nil}
	node4 := Node{8, nil, &node3}
	node5 := Node{9, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node2, &node6}

	printRange(&node7, 2, 8)
	// expected output: 2 5 8 8
	//printRange(&node7, 0, 5)
}

//func main() {
//	test()
//}
