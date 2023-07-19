package main

// Comment it before submitting
//type Node struct {
//	value int
//	left  *Node
//	right *Node
//}

func Solution(root *Node) int {
	stack := make([]*Node, 0, 100)
	max := root.value

	stack = append(stack, root)

	for len(stack) != 0 {
		iLast := len(stack) - 1

		node := stack[iLast]
		stack = stack[:iLast]

		if node.left != nil {
			stack = append(stack, node.left)
		}

		if node.right != nil {
			stack = append(stack, node.right)
		}

		if max < node.value {
			max = node.value
		}
	}

	return max
}

func test() {
	node1 := Node{1, nil, nil}
	node2 := Node{-5, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{2, &node3, nil}
	if Solution(&node4) != 3 {
		panic("WA")
	}
}
