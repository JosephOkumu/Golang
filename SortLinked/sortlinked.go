package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func sortLinked(l *Node) *Node {
	if l == nil || l.next == nil {
		return l
	}
	middle := findMiddle(l)
	nextMiddle := middle.next
	middle.next = nil
	left := sortLinked(l)
	right := sortLinked(nextMiddle)
	return merge(left, right)
}
func findMiddle(l *Node) *Node {
	slow := l
	fast := l
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
func merge(l1 *Node, l2 *Node) *Node {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.data <= l2.data {
		l1.next = merge(l1.next, l2)
		return l1
	} else {
		l2.next = merge(l1, l2.next)
		return l2
	}
}

func printList(l *Node) {
	it := l
	for it != nil {
		fmt.Print(it.data, " -> ")
		it = it.next
	}
	fmt.Print(nil, "\n")
}

func addNodEnd(l *Node, data int) *Node {
	n := &Node{data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.next != nil {
		iterator = iterator.next
	}
	iterator.next = n
	return l
}

func main() {
	var link *Node

	link = addNodEnd(link, 5)
	link = addNodEnd(link, 4)
	link = addNodEnd(link, 3)
	link = addNodEnd(link, 2)
	link = addNodEnd(link, 1)

	printList(sortLinked(link))
}
