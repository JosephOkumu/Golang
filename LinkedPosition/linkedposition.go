package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

// List represents the linked list
type List struct {
	head *Node // pointer to the first node
	tail *Node // pointer to the last node
}

// addNodEnd adds a new node to the end of the list
func addNodEnd(l *List, data interface{}) {
	newNode := &Node{data: data, next: nil}

	if l.head == nil {
		l.head = newNode // set the new node as the head
		l.tail = newNode // set the new node as the tail
	} else {
		l.tail.next = newNode // set the next of the current tail to the new node
		l.tail = newNode      // update the tail to the new node
	}
}

// linkedPosition returns the node at a specific position in the list
func linkedPosition(N *Node, pos int) *Node {
	//set iterator to the head of the node
	iterator := N
	count := 0

	// traverse the list until the iterator reaches the end or the desired position
	for iterator != nil && count < pos {
		iterator = iterator.next
		count++
	}

	if count == pos { // if the desired position is found
		return iterator // return the node at that position
	}
	return nil // otherwise, return nil
}

func main() {
	link := &List{}

	// Add elements to the list
	addNodEnd(link, "Hello")
	addNodEnd(link, "How are")
	addNodEnd(link, "you")
	addNodEnd(link, 1)

	// Retrieve and print the node at specific positions
	fmt.Println(linkedPosition(link.head, 3).data) // Print the node at position 3
	fmt.Println(linkedPosition(link.head, 1).data) // Print the node at position 1
	fmt.Println(linkedPosition(link.head, 7))      // Print the node at position 7 (which is nil)
}
