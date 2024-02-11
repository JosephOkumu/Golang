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

func linkedLen(l *List) int {
	count := 0        // initialize a counter variable
	current := l.head // set the current node to the head of the list

	// traverse the list and increment the counter for each node
	for current != nil {
		count++
		current = current.next
	}

	return count // return the count as the length of the list
}

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
	link := &List{}
	addNodeFront(link, "Why")
	addNodeFront(link, "are people")
	addNodeFront(link, "acting so inhumane")
	addNodeFront(link, "I don't know")
	fmt.Println(linkedLen(link))
}
