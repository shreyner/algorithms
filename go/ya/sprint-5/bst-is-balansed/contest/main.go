package main

import "math"

/*
*
Comment it before submitting
*/
//type Node struct {
//	value int
//	left  *Node
//	right *Node
//}

func isBSTBalanced(root *Node) (bool, int) {
	if root == nil {
		return true, 0
	}

	validLeft, leftHeight := isBSTBalanced(root.left)
	validRight, rightHeight := isBSTBalanced(root.right)

	if !validLeft || !validRight {
		return false, 0
	}

	if math.Abs(float64(leftHeight-rightHeight)) <= 1 {
		if leftHeight > rightHeight {
			return true, leftHeight + 1
		}

		return true, rightHeight + 1
	}

	return false, 0
}

func Solution(root *Node) bool {
	isValid, _ := isBSTBalanced(root)

	return isValid
}

func test() {
	//node1 := Node{1, nil, nil}
	//node2 := Node{-5, nil, nil}
	//node3 := Node{3, &node1, &node2}
	//node4 := Node{10, nil, nil}
	//node5 := Node{2, &node3, &node4}
	//
	//if !Solution(&node5) {
	//	panic("WA")
	//}

	//node5 := Node{5, nil, nil}
	//node6 := Node{6, nil, nil}
	//node3 := Node{3, &node5, &node6}
	//node1 := Node{1, &node3, nil}
	//
	//node9 := Node{9, nil, nil}
	//node8 := Node{8, nil, &node9}
	//node7 := Node{7, nil, nil}
	//node4 := Node{4, &node7, &node8}
	//node2 := Node{2, nil, &node4}
	//
	//node0 := Node{0, &node1, &node2}
	//
	//if Solution(&node0) {
	//	panic("WA")
	//}
}

//func main() {
//	test()
//}
