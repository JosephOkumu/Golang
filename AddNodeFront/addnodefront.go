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

// addNodeFront adds a new node to the front of the list
func addNodeFront(l *List, data interface{}) {
	newNode := &NodeL{
		data: data,
		next: nil,
	}
	if l.head == nil { // if the list is empty
		l.head = newNode // set the new node as the head
		l.tail = newNode // set the new node as the tail
	} else {
		newNode.next = l.head // set the next of the new node to the current head
		l.head = newNode      // update the head to the new node
	}
}

func main() {
	link := &List{} // create a new empty list

	// add elements to the list
	addNodeFront(link, "Some")
	addNodeFront(link, "people")
	addNodeFront(link, "can pretend!")

	// print the elements of the list
	it := link.head
	for it != nil {
		fmt.Print(it.data, " ")
		it = it.next
	}
}
