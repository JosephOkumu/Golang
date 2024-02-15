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

func main() {
	root := &Trnode{data: "4"}
	bstreeInsert(root, "1")
	bstreeInsert(root, "7")
	bstreeInsert(root, "5")
	selected := bstreeSearchElem(root, "7")

	fmt.Print("item selected -> ")
	if selected != nil {
		fmt.Println(selected.data)
	} else {
		fmt.Println("nil")
	}

	fmt.Print("Parent of selected item -> ")
	if selected.parent != nil {
		fmt.Println(selected.parent.data)
	} else {
		fmt.Println("nil")
	}

	fmt.Print("Left child of selcected item -> ")
	if selected.left != nil {
		fmt.Println(selected.left.data)
	} else {
		fmt.Println("nil")
	}
	fmt.Print("Right child of selected item -> ")
	if selected.right != nil {
		fmt.Println(selected.right.data)
	} else {
		fmt.Println("nil")
	}
}
