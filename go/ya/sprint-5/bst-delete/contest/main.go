package main

/*
*
Comment it before submitting
*/
type Node struct {
	value int
	left  *Node
	right *Node
}

func findMostLeft(node, parent *Node) (*Node, *Node) {
	for node != nil {
		if node.right == nil {
			return node, parent
		}

		parent = node
		node = node.right
	}

	return nil, nil
}

func findMostRight(node, parent *Node) (*Node, *Node) {
	for node != nil {
		if node.left == nil {
			return node, parent
		}

		parent = node
		node = node.left
	}

	return nil, nil
}

func findTarget(node, parent *Node, key int) (*Node, *Node) {
	for node != nil {
		if node == nil || node.value == key {
			return node, parent
		}

		parent = node
		if node.value > key {
			node = node.left
		} else {
			node = node.right
		}
	}

	return nil, nil
}

func isLeaf(node *Node) bool {
	return node.left == nil && node.right == nil
}

func replaceChild(parent, target, replace *Node) {
	if parent == nil {
		return
	}

	if parent.left == target {
		if replace != nil {
			parent.left = replace
		} else {
			parent.left = target.right
		}

	}

	if parent.right == target {
		if replace != nil {
			parent.right = replace
		} else {
			parent.right = target.left
		}
	}
}

func remove(root *Node, key int) *Node {
	if root == nil {
		return root
	}

	target, parent := findTarget(root, nil, key)

	if target == nil {
		return root
	}

	if isLeaf(target) {
		if parent == nil {
			return nil
		}

		replaceChild(parent, target, nil)

		return root
	}

	var replaceNode, replaceParent *Node

	if target.left != nil {
		replaceNode, replaceParent = findMostLeft(target.left, target)
	}

	if replaceNode == nil && target.right != nil {
		replaceNode, replaceParent = findMostRight(target.right, target)
	}

	if parent == nil {
		root = replaceNode
	}

	replaceChild(parent, target, replaceNode)
	replaceChild(replaceParent, replaceNode, nil)

	if target.left != nil {
		replaceNode.left = target.left
	}

	if target.right != nil {
		replaceNode.right = target.right
	}

	return root
}

func test() {
	node1 := Node{2, nil, nil}
	node2 := Node{3, &node1, nil}
	node3 := Node{1, nil, &node2}
	node4 := Node{6, nil, nil}
	node5 := Node{8, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node3, &node6}

	newHead := remove(&node7, 10)

	if newHead.value != 5 {
		panic("WA")
	}

	if newHead.right != &node5 {
		panic("WA")
	}

	if newHead.right.value != 8 {
		panic("WA")
	}
}

func test2() {
	node1 := Node{1, nil, nil}
	node13 := Node{13, nil, nil}
	node7 := Node{7, nil, nil}
	node9 := Node{9, nil, nil}
	node8 := Node{8, &node7, &node9}
	node10 := Node{10, &node8, nil}
	node12 := Node{12, &node10, &node13}
	node5 := Node{5, &node1, &node12}

	newHead := remove(&node5, 12)

	if newHead.value != 5 {
		panic("WA")
	}

	if newHead.right != &node10 {
		panic("WA")
	}

	if newHead.right.value != 10 {
		panic("WA")
	}
}

func test3() {
	node1 := Node{1, nil, nil}
	node13 := Node{13, nil, nil}
	node7 := Node{7, nil, nil}
	node9 := Node{9, nil, nil}
	node8 := Node{8, &node7, &node9}
	node10 := Node{10, &node8, nil}
	node12 := Node{12, &node10, &node13}
	node5 := Node{5, &node1, &node12}

	newHead := remove(&node5, 1)

	if newHead.value != 5 {
		panic("WA")
	}

	if newHead.left != nil {
		panic("WA")
	}
}

func test4() {
	node1 := Node{1, nil, nil}
	node13 := Node{13, nil, nil}
	node7 := Node{7, nil, nil}
	node9 := Node{9, nil, nil}
	node8 := Node{8, &node7, &node9}
	node10 := Node{10, &node8, nil}
	node12 := Node{12, &node10, &node13}
	node5 := Node{5, &node1, &node12}

	newHead := remove(&node5, 13)

	if newHead.value != 5 {
		panic("WA")
	}

	if node12.right != nil {
		panic("WA")
	}
}

func test5() {
	node1 := Node{2, nil, nil}
	node2 := Node{3, &node1, nil}
	node3 := Node{1, nil, &node2}
	node4 := Node{6, nil, nil}
	node5 := Node{8, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node3, &node6}

	newHead := remove(&node7, 5)

	if newHead.value != 3 {
		panic("WA")
	}

	if newHead.left.value != 1 {
		panic("WA")
	}

	if newHead.left.right.value != 2 {
		panic("WA")
	}

	if newHead.right.value != 10 {
		panic("WA")
	}
}

func test6() {
	node2 := Node{2, nil, nil}
	node1 := Node{1, nil, &node2}

	newHead := remove(&node1, 1)

	if newHead.value != 2 {
		panic("WA")
	}
}

func test7() {
	node1 := Node{2, nil, nil}
	node2 := Node{3, &node1, nil}
	node3 := Node{1, nil, &node2}
	node4 := Node{6, nil, nil}
	node5 := Node{8, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node3, &node6}

	_ = remove(&node7, 200)
}

func main() {
	test()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
}
