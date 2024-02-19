package main

import "fmt"

type Trnode struct {
	parent *Trnode
	left   *Trnode
	right  *Trnode
	data   string
}

func bsTreeDeleteNode(root, node *Trnode) *Trnode {
	if root == nil || node == nil {
		return root
	}

	// Case 1: Node is a leaf node
	if node.left == nil && node.right == nil {
		if node.parent != nil {
			if node.parent.left == node {
				node.parent.left = nil
			} else {
				node.parent.right = nil
			}
		} else {
			root = nil
		}
		return root
	}

	// Case 2: Node has only one child
	if node.left == nil {
		child := node.right
		replaceNode(root, node, child)
		return root
	}
	if node.right == nil {
		child := node.left
		replaceNode(root, node, child)
		return root
	}

	// Case 3: Node has two children
	successor := getMinimumNode(node.right)
	node.data = successor.data
	root = bsTreeDeleteNode(root, successor)
	return root
}

// Helper function to replace a node with another node
func replaceNode(root, node, replace *Trnode) {
	if node.parent != nil {
		if node.parent.left == node {
			node.parent.left = replace
		} else {
			node.parent.right = replace
		}

	} else {
		root = replace
	}

	if replace != nil {
		replace.parent = node.parent
	}
}

// Helper function to find the minimum value node in a subtree
func getMinimumNode(node *Trnode) *Trnode {
	for node.left != nil {
		node = node.left
	}
	return node
}

func bstreeInsert(root *Trnode, data string) *Trnode {
	if root == nil {
		return &Trnode{data: data}
	}

	if data < root.data {
		if root.left == nil {
			root.left = &Trnode{data: data, parent: root}
		} else {
			root.left = bstreeInsert(root.left, data)
		}
	} else {
		if root.right == nil {
			root.right = &Trnode{data: data, parent: root}
		} else {
			root.right = bstreeInsert(root.right, data)
		}
	}
	return root
}

func bstreeSearchElem(root *Trnode, elem string) *Trnode {
	if root == nil {
		return nil
	}
	if root.data == elem {
		return root
	}

	left := bstreeSearchElem(root.left, elem)

	if left != nil {
		return left
	}
	right := bstreeSearchElem(root.right, elem)
	if right != nil {
		return right
	}
	return nil
}

func BsTreeApplyInorder(root *Trnode, f func(...interface{}) (int, error)) {
	if root != nil {
		BsTreeApplyInorder(root.left, f)

		_, err := f(root.data)

		if err != nil {
			return
		}
		BsTreeApplyInorder(root.right, f)
	}
}

func main() {
	root := &Trnode{data: "4"}
	bstreeInsert(root, "1")
	bstreeInsert(root, "7")
	bstreeInsert(root, "5")
	node := bstreeSearchElem(root, "4")

	fmt.Println("Before Deletion:")
	BsTreeApplyInorder(root, fmt.Println)
	root = bsTreeDeleteNode(root, node)
	fmt.Println("After Deletion:")
	BsTreeApplyInorder(root, fmt.Println)
}
