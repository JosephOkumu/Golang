package main

import "fmt"

// Node represents a node in the linked list
type Node struct {
	data interface{}
	next *Node
}

// List represents the linked list
type List struct {
	head *Node
	tail *Node
}

// addNodEnd adds a new node to the end of the list
func addNodEnd(l *List, data interface{}) {
	newNode := &Node{data: data, next: nil}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
}

// listForEach applies a function to each node in the list
func listForEach(l *List, f func(*Node)) {
	iterator := l.head
	for iterator != nil {
		f(iterator)
		iterator = iterator.next
	}
}

// add2Node adds 2 to the node's data if it's an int, or appends "2" if it's a string
func add2Node(node *Node) {
	switch node.data.(type) {
	case int:
		node.data = node.data.(int) + 2
	case string:
		node.data = node.data.(string) + "2"
	}
}

func main() {
	link := &List{}

	// Adding nodes to the list
	addNodEnd(link, "1")
	addNodEnd(link, "2")
	addNodEnd(link, "3")
	addNodEnd(link, "5")

	// Applying add2Node function to each node in the list
	listForEach(link, add2Node)

	// Printing the data in each node
	it := link.head
	for it != nil {
		fmt.Println(it.data)
		it = it.next
	}
}
