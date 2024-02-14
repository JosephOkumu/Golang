package main

import "fmt"

type Trnode struct {
	parent *Trnode
	left   *Trnode
	right  *Trnode
	data   string
}

// Function to apply a function to each node in a binary search tree in inorder traversal
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

func bstreeInsert(root *Trnode, data string) *Trnode {

	// Base case: if the root is nil, create a new node
	if root == nil {
		return &Trnode{data: data}
	}

	// Recursive case: compare the data with the root's data
	if data < root.data {
		// If the data is smaller, insert into the left subtree
		if root.left == nil {
			// Create a new node and set its parent
			root.left = &Trnode{data: data, parent: root}
		} else {
			// Recursively call the insert function for the left subtree
			root.left = bstreeInsert(root.left, data)
		}
	} else {
		// If the data is larger, insert into the right subtree
		if root.right == nil {
			// Create a new node and set its parent
			root.right = &Trnode{data: data, parent: root}
		} else {
			// Recursively call the insert function for the right subtree
			root.right = bstreeInsert(root.right, data)
		}
	}

	// Return the root node
	return root
}

func main() {
	root := &Trnode{data: "4"}
	bstreeInsert(root, "1")
	bstreeInsert(root, "7")
	bstreeInsert(root, "5")

	BsTreeApplyInorder(root, fmt.Println)
}
