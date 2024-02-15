package main

import "fmt"

type Trnode struct {
	parent *Trnode
	left   *Trnode
	right  *Trnode
	data   string
}

func BTreeApplyPostorder(root *Trnode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyPostorder(root.left, f)

		BTreeApplyPostorder(root.right, f)

		_, err := f(root.data)

		if err != nil {
			return
		}
	}
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

func main() {
	root := &Trnode{data: "4"}
	bstreeInsert(root, "1")
	bstreeInsert(root, "7")
	bstreeInsert(root, "5")
	BTreeApplyPostorder(root, fmt.Println)
}
