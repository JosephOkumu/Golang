// This program finds the node with maximum value stored in binary search tree
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

func bsTreeMax(root *Trnode) *Trnode {
	if root == nil {
		return nil
	}

	maxnode := root

	//Traverse the right subtree first
	if root.right != nil {
		rightmax := bsTreeMax(root.right)
		if rightmax.data > maxnode.data {
			maxnode = rightmax
		}
	}

	//Traverse the current node
	if root.data > maxnode.data {
		maxnode = root
	}

	//Traverse the left subtree
	if root.left != nil {
		leftmax := bsTreeMax(root.left)
		if leftmax.data > maxnode.data {
			maxnode = leftmax
		}
	}
	return maxnode
}

func main() {
	root := &Trnode{data: "4"}
	bstreeInsert(root, "1")
	bstreeInsert(root, "7")
	bstreeInsert(root, "5")
	max := bsTreeMax(root)
	fmt.Println(max.data)
}
