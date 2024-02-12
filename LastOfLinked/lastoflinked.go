// This program displays the last interface of the linked list.
package main

import "fmt"

type Node struct {
	data interface{}
	next *Node // pointer to the next node
}

// List is the linked list
type List struct {
	head *Node // pointer to the first node
	tail *Node // pointer to the last node
}

// lastLinked returns the last element of the list
func lastLinked(l *List) interface{} {
	if l.head == nil {
		return nil
	}

	iterator := l.head
	for iterator.next != nil { // traverse the list until the last node
		iterator = iterator.next
	}
	return iterator.data // return the data of the last node
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
	link2 := &List{}

	addNodEnd(link, "three")
	addNodEnd(link, 3)
	addNodEnd(link, "1")

	fmt.Println(lastLinked(link))  // print the last element of link
	fmt.Println(lastLinked(link2)) // print the last element of link2 (which is nil)
}
