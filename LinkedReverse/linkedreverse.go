package main

import "fmt"

// Node represents a node in the linked list
type Node struct {
	data interface{} // data of any type
	next *Node       // pointer to the next node
}

// List represents the linked list
type List struct {
	head *Node // pointer to the first node
	tail *Node // pointer to the last node
}

// linkedReverse reverses the linked list
func linkedReverse(l *List) {
	var prev *Node
	current := l.head
	temp := l.head

	// traverse the list and reverse the pointers
	for current != nil {
		temp = current.next
		current.next = prev
		prev = current
		current = temp
	}

	if l.head != nil {
		l.head = prev // update the head to the last node
	}
}

// addNodEnd adds a new node to the end of the list
func addNodEnd(l *List, data interface{}) {
	newNode := &Node{data: data, next: nil}

	if l.head == nil { // if the list is empty
		l.head = newNode // set the new node as the head
		l.tail = newNode // set the new node as the tail
	} else {
		l.tail.next = newNode // set the next of the current tail to the new node
		l.tail = newNode      // update the tail to the new node
	}
}

func main() {
	link := &List{}

	// Add elements to the list
	addNodEnd(link, 1)
	addNodEnd(link, 2)
	addNodEnd(link, 3)
	addNodEnd(link, 4)

	linkedReverse(link) // Reverse the list

	it := link.head

	// Print the elements of the reversed list
	for it != nil {
		fmt.Println(it.data)
		it = it.next
	}
}
