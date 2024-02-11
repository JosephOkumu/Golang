package main

import "fmt"

// NodeL is a node in the linked list
type NodeL struct {
	data interface{} // data of any type
	next *NodeL      // pointer to the next node
}

// List is the linked list
type List struct {
	head *NodeL // pointer to the first node
	tail *NodeL // pointer to the last node
}

// addNodEnd adds a new node to the end of the list
func addNodEnd(l *List, data interface{}) {
	newNode := &NodeL{data: data, next: nil} // create a new node with the given data

	if l.head == nil { // if the list is empty
		l.head = newNode // set the new node as the head
		l.tail = newNode // set the new node as the tail
	} else {
		l.tail.next = newNode // set the next of the current tail to the new node
		l.tail = newNode      // update the tail to the new node
	}
}

func main() {
	link := &List{} // create a new empty list

	// add elements to the list
	addNodEnd(link, "Hello")
	addNodEnd(link, "man")
	addNodEnd(link, "how are you")

	// print the elements of the list
	for link.head != nil {
		fmt.Println(link.head.data) // print the data of the current node
		link.head = link.head.next  // move to the next node
	}
}
