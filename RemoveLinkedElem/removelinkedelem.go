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

// Removes nodes with a specific value from the list
func listRemoveif(l *List, data_ref interface{}) {
	current := l.head
	previous := (*Node)(nil)

	for current != nil {
		if current.data == data_ref {
			if previous != nil {
				previous.next = current.next
			} else {
				l.head = current.next
			}
			current = current.next
			continue
		}
		previous = current
		current = current.next
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

	fmt.Println("---- normal state ----")
	addNodEnd(link2, 1)
	printList(link2)
	listRemoveif(link2, 1)
	fmt.Println("----- answer -----")
	printList(link2)
	fmt.Println()

	fmt.Println("---- normal state ----")
	addNodEnd(link, 1)
	addNodEnd(link, "Hello")
	addNodEnd(link, 1)
	addNodEnd(link, "There")
	addNodEnd(link, 1)
	addNodEnd(link, 1)
	addNodEnd(link, "How")
	addNodEnd(link, 1)
	addNodEnd(link, "are")
	addNodEnd(link, "you")
	addNodEnd(link, 1)
	printList(link)

	listRemoveif(link, 1)
	fmt.Println("------answer------")
	printList(link)

}
