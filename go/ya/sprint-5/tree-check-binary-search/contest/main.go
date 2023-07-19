package main

// Comment it before submitting
//type Node struct {
//	value int
//	left  *Node
//	right *Node
//}

func isBST(node *Node) (valid bool, min *int, max *int) {
	if node == nil {
		return true, nil, nil
	}

	leftValid, leftMin, leftMax := isBST(node.left)
	rightValid, rightMin, rightMax := isBST(node.right)

	if !leftValid || !rightValid {
		return false, nil, nil
	}

	if (leftMax != nil && *leftMax >= node.value) || (rightMin != nil && *rightMin <= node.value) {
		return false, nil, nil
	}

	min = new(int)
	max = new(int)

	*min = node.value
	*max = node.value

	if leftMin != nil && *leftMin < *min {
		min = leftMin
	}

	if rightMin != nil && *rightMin < *min {
		min = rightMin
	}

	if rightMax != nil && *rightMax > *max {
		max = rightMax
	}

	if leftMax != nil && *leftMax > *max {
		max = leftMax
	}

	return true, min, max
}

func Solution(root *Node) bool {
	valid, _, _ := isBST(root)

	return valid
}

func test() {
	//node1 := Node{1, nil, nil}
	//node2 := Node{4, nil, nil}
	//node3 := Node{3, &node1, &node2}
	//node6 := Node{6, nil, nil}
	//node9 := Node{9, nil, nil}
	//node4 := Node{8, &node6, &node9}
	//node5 := Node{5, &node3, &node4}

	//if !Solution(&node5) {
	//	panic("WA")
	//}
	//node2.value = 5
	//node9.value = 3
	//if Solution(&node5) {
	//	panic("WA")
	//}

	//node1 := Node{1, nil, nil}
	//node2 := Node{3, nil, nil}
	//node3 := Node{3, &node1, &node2}
	//node6 := Node{6, nil, nil}
	//node9 := Node{9, nil, nil}
	//node4 := Node{8, &node6, &node9}
	//node5 := Node{5, &node3, &node4}
	//
	//if Solution(&node5) {
	//	panic("WA")
	//}
}

//func main() {
//	test()
//}
