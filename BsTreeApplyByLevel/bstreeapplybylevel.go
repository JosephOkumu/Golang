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

func BTreeApplyByLevel(root *Trnode, f func(...interface{}) (int, error)) {
	//check if the root node is nil
	if root == nil {
		return
	}

	//Create a queue to store nodes to be processed
	queue := []*Trnode{root}

	//process nodes in the queue until it is empty
	for len(queue) > 0 {
		//Dequeue a node from the front of the queue
		node := queue[0]
		queue = queue[1:]

		//Apply the given function to the node's data
		f(node.data)

		//Enqueue the left and right child nodes if they exist
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}

}

func main() {
	root := &Trnode{data: "4"}
	bstreeInsert(root, "1")
	bstreeInsert(root, "7")
	bstreeInsert(root, "5")
	BTreeApplyByLevel(root, fmt.Println)
}
