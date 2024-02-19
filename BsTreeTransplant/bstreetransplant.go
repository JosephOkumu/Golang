package main

import "fmt"

type Trnode struct {
	parent *Trnode
	left   *Trnode
	right  *Trnode
	data   string
}

func bstreeInsert(root *Trnode, data string) *Trnode {
	if root == nil {
		return nil
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

	//compare the data of root node with elem
	if root.data == elem {
		return root
	}

	//recursively call bstreeSearchElem on the left and the right children of the root.
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
	// Base case: if the root is nil, return
	if root != nil {
		// Recursive case: traverse the left subtree
		BsTreeApplyInorder(root.left, f)

		// Apply the function to the current node's data
		_, err := f(root.data)

		// If the function returns an error, return
		if err != nil {
			return
		}

		// Recursive case: traverse the right subtree
		BsTreeApplyInorder(root.right, f)
	}
}

func BsTreeTransplant(root, node, rplc *Trnode) *Trnode {
	if node == nil || root == nil {
		return root
	}

	if node == root {
		root = rplc
	} else {
		if node.parent.left == node {
			node.parent.left = rplc
		} else {
			node.parent.right = rplc
		}
	}
	if rplc != nil {
		rplc.parent = node.parent
	}
	return root

}

func main() {
	root := &Trnode{data: "4"}
	bstreeInsert(root, "1")
	bstreeInsert(root, "7")
	bstreeInsert(root, "5")
	node := bstreeSearchElem(root, "1")
	rplc := &Trnode{data: "3"}
	root = BsTreeTransplant(root, node, rplc)
	BsTreeApplyInorder(root, fmt.Println)
}
