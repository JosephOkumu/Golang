package main

import "fmt"

// Define the TreeNode type for the binary search tree
type Trnode struct {
	data   string
	parent *Trnode
	left   *Trnode
	right  *Trnode
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

// Function to demonstrate the usage of the bstreeInsert function
func main() {
	// Create the root node with the value 4
	root := &Trnode{data: "4"}

	// Insert nodes with values 1, 7, and 5 into the tree
	bstreeInsert(root, "1")
	bstreeInsert(root, "7")
	bstreeInsert(root, "5")

	// Print the data of the left child of the root
	fmt.Println(root.left.data)

	// Print the data of the root
	fmt.Println(root.data)

	// Print the data of the left child of the right child of the root
	fmt.Println(root.right.left.data)

	// Print the data of the right child of the root
	fmt.Println(root.right.data)
}
