// This program checks if a tree follows the binary search tree properties.
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

func bsTreeIsBinary(root *Trnode) bool {
	return isBinaryHelper(root, nil, nil)

}

func isBinaryHelper(node *Trnode, min *Trnode, max *Trnode) bool {
	if node == nil {
		return true
	}
	if (min != nil && node.data <= min.data) || (max != nil && node.data >= max.data) {
		return false
	}
	return isBinaryHelper(node.left, min, node) && isBinaryHelper(node.right, node, max)
}

func main() {
	root := &Trnode{data: "4"}
	bstreeInsert(root, "1")
	bstreeInsert(root, "7")
	bstreeInsert(root, "5")
	fmt.Println(bsTreeIsBinary(root))
}
