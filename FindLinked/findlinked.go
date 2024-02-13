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

// addNodEnd adds a new node to the end of the linked list
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

// compStr compares two strings and returns a boolean value
func compStr(a, b interface{}) bool {
	return a == b
}

// findLinked searches for a node in the linked list that matches a given reference value using a comparison function
func findLinked(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	iterator := l.head

	for iterator != nil {
		if comp(iterator.data, ref) {
			return &iterator.data
		}
		iterator = iterator.next
	}
	return nil
}

func main() {
	link := &List{}

	addNodEnd(link, "Hello")
	addNodEnd(link, "Hello1")
	addNodEnd(link, "Hello2")
	addNodEnd(link, "Hello3")

	found := findLinked(link, interface{}("Hello2"), compStr)

	fmt.Println(found)
	fmt.Println(*found)
}
