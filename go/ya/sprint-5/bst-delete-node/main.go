package main

/**
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
