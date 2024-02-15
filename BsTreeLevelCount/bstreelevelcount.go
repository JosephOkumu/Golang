package main

import "fmt"

type Trnode struct {
	parent *Trnode
	left   *Trnode
	right  *Trnode
	data   string
}

func bsTreeLevelCount(root *Trnode) int {
	if root == nil {
		return 0
	}

	//Recursively call bsTreeLevelCount on the left and right children of the root
	left := bsTreeLevelCount(root.left)
	right := bsTreeLevelCount(root.right)

	//Return the maximum of left and right plus 1
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
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

func main() {
	root := &Trnode{data: "4"}
	bstreeInsert(root, "1")
	bstreeInsert(root, "7")
	bstreeInsert(root, "5")
	fmt.Println(bsTreeLevelCount(root))
}
