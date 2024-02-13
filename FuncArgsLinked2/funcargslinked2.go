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

// isPositiveNode returns true if the node's data is a positive number
func isPositiveNode(node *Node) bool {
	switch node.data.(type) {
	case int, float32, float64, byte:
		return node.data.(int) > 0
	default:
		return false
	}
}

// isAnode returns true if the node's data is not a number
func isAnode(node *Node) bool {
	switch node.data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}

// listForEachIf applies a function to each node in the list that satisfies a given condition
func listForEachIf(l *List, f func(*Node), cond func(*Node) bool) {
	iterator := l.head
	for iterator != nil {
		if cond(iterator) {
			f(iterator)
		}
		iterator = iterator.next
	}

}

// printElem prints the data in a node
func printElem(node *Node) {
	fmt.Println(node.data)
}

// stringToInt sets the node's data to 2
func stringToInt(node *Node) {
	node.data = 2
}

// printList prints the data in each node of the list
func printList(l *List) {
	it := l.head
	for it != nil {
		fmt.Print(it.data, "->")
		it = it.next
	}
	fmt.Print("nil", "\n")
}

func main() {
	link := &List{}

	// Adding nodes to the list
	addNodEnd(link, 1)
	addNodEnd(link, "Hello")
	addNodEnd(link, 3)
	addNodEnd(link, "there")
	addNodEnd(link, 23)
	addNodEnd(link, "!")
	addNodEnd(link, 54)

	// Printing the list before applying functions
	printList(link)

	// Applying printElem function to each node in the list that satisfies isPositiveNode condition
	fmt.Println("-------- function applied --------")
	listForEachIf(link, printElem, isPositiveNode)

	// Applying stringToInt function to each node in the list that satisfies isAnode condition
	listForEachIf(link, stringToInt, isAnode)

	// Printing the list after applying functions
	fmt.Println("--------- function applied --------")
	printList(link)

	fmt.Println()
}
