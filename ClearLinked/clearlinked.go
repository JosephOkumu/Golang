package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

// The code block addNodEnd adds a node at the end of the linked list
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

// This function here is what clears the linked list
func clearLinked(l *List) {
	l.head = nil
	l.tail = nil
}

// This block prints the linked list
func printList(l *List) {
	iterator := l.head
	for iterator != nil {
		fmt.Print(iterator.data, " -> ")
		iterator = iterator.next
	}
	fmt.Println(nil)
}

func main() {
	link := &List{}

	addNodEnd(link, "I")
	addNodEnd(link, 7)
	addNodEnd(link, "Christian")
	addNodEnd(link, 35)

	fmt.Println("------list------")
	printList(link)
	clearLinked(link)
	fmt.Println("------updated list------")
	printList(link)
}
