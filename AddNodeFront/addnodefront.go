package main

import "fmt"

type NodeL struct {
	data interface{}
	next *NodeL
}

type List struct {
	head *NodeL
	tail *NodeL
}

func addNodeFront(l *List, data interface{}) {
	newNode := &NodeL{
		data: data,
		next: nil,
	}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head = newNode

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
