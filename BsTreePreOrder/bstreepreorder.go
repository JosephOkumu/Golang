package main

import "fmt"

// Define the TreeNode type for the binary search tree
type Trnode struct {
	parent *Trnode
	left   *Trnode
	right  *Trnode
	data   string
}

// Function to apply a function to each node in a binary search tree in preorder traversal
func BTreeApplyPreorder(root *Trnode, f func(...interface{}) (int, error)) {
	// Base case: if the root is nil, return
	if root != nil {
		// Apply the function to the current node's data
		_, err := f(root.data)

		// If the function returns an error, return
		if err != nil {
			return
		}

		// Recursive case: traverse the left subtree
		BTreeApplyPreorder(root.left, f)

		// Recursive case: traverse the right subtree
		BTreeApplyPreorder(root.right, f)
	}
}

// Function to insert a node into the binary search tree
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

// Function to demonstrate the usage of the bstreeInsert and BTreeApplyPreorder functions
func main() {
	// Create the root node with the value 4
	root := &Trnode{data: "4"}

	// Insert nodes with values 1, 7, and 5 into the tree
	bstreeInsert(root, "1")
	bstreeInsert(root, "7")
	bstreeInsert(root, "5")

	// Apply the fmt.Println function to each node in the tree in preorder traversal
	BTreeApplyPreorder(root, fmt.Println)
}
