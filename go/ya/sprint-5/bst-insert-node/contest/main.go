package main

// Comment it before submitting
type Node struct {
	value int
	left  *Node
	right *Node
}

func insert(root *Node, key int) *Node {
	if key < root.value {
		if root.left == nil {
			root.left = &Node{key, nil, nil}
		} else {
			insert(root.left, key)
		}
	}

	if root.value <= key {
		if root.right == nil {
			root.right = &Node{key, nil, nil}
		} else {
			insert(root.right, key)
		}
	}

	return root
}

func test() {
	node1 := Node{7, nil, nil}
	node2 := Node{8, &node1, nil}
	node3 := Node{7, nil, &node2}
	newHead := insert(&node3, 6)

	if newHead != &node3 {
		panic("WA")
	}

	if newHead.left.value != 6 {
		panic("WA")
	}
}

func main() {
	test()
}
