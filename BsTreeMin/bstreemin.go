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

func bsTreeMin(root *Trnode) *Trnode {
	if root == nil {
		return nil
	}

	minNode := root

	//Traverse the right subtree first
	if root.right != nil {
		rightmin := bsTreeMin(root.right)
		if rightmin.data < minNode.data {
			minNode = rightmin
		}
	}

	//Traverse the current node
	if root.data < minNode.data {
		minNode = root
	}

	//Traverse the left subtree
	if root.left != nil {
		leftmin := bsTreeMin(root.left)
		if leftmin.data < minNode.data {
			minNode = leftmin
		}
	}
	return minNode
}

func main() {
	root := &Trnode{data: "4"}
	bstreeInsert(root, "1")
	bstreeInsert(root, "7")
	bstreeInsert(root, "5")
	min := bsTreeMin(root)
	fmt.Println(min.data)
}
