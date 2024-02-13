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

func linkedMerge(l1 *List, l2 *List) {
	if l2.head == nil {
		return
	}
	if l1.head == nil {
		l1.head = l2.head
		l1.tail = l2.tail
	} else {
		l1.tail.next = l2.head
		l1.tail = l2.tail

	}
}

func printList(l *List) {
	it := l.head

	for it != nil {
		fmt.Print(it.data, " -> ")
		it = it.next
	}
	fmt.Print(nil, "\n")
}

func main() {
	link := &List{}
	link2 := &List{}

	addNodEnd(link, "a")
	addNodEnd(link, "b")
	addNodEnd(link, "c")
	addNodEnd(link, "d")
	fmt.Println("-----first List-----")
	printList(link)

	addNodEnd(link2, "e")
	addNodEnd(link2, "f")
	addNodEnd(link2, "g")
	addNodEnd(link2, "h")
	fmt.Println("-----second List-----")
	printList(link2)

	fmt.Println("-----Merged List-----")
	linkedMerge(link, link2)
	printList(link)
}
