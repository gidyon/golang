package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func insertAtBeginning(head *Node, data int) *Node {
	node := new(Node)
	node.data = data
	if head == nil {
		node.next = nil
		return head
	}
	node.next = head
	head = node
	return head
}

func insertAtEnd(head *Node, data int) *Node {
	newNode := new(Node)
	newNode.data = data
	newNode.next = nil
	if lengthOfList(head) <= 1 {
		head = insertAtBeginning(head, data)
		return head
	}
	temp := head
	for temp.next != nil {
		temp = temp.next
	}
	fmt.Println(temp)
	temp.next = newNode
	return head
}

func lengthOfList(head *Node) (length int) {
	temp := head
	for temp.next != nil {
		length++
		temp = temp.next
	}
	return
}

func insertAtIndex(head *Node, data int, pos int) (newHead *Node) {
	if pos <= 1 {
		newHead = insertAtBeginning(head, data)
		return
	}
	if pos > lengthOfList(head) {
		newHead = insertAtEnd(head, data)
		return
	}
	newNode := new(Node)
	newNode.data = data
	temp := head
	for i := 0; i < pos-2; i++ {
		temp = temp.next
	}
	newNode.next = temp.next
	temp.next = newNode
	return head
}

func print(head *Node) {
	fmt.Print("The linked list is: { ")
	for head.next != nil {
		fmt.Print(head.data, " ")
		head = head.next
	}
	fmt.Print(" }\n")
}
func deleteAtBeginning(head *Node) *Node {
	temp := head.next
	head = temp
	return head
}
func deleteAtEnd(head *Node) *Node {
	prev, temp := head, head
	for temp.next != nil {
		prev = temp
		temp = temp.next
	}
	prev.next = nil
	return head
}
func deleteAtIndex(head *Node, pos int) *Node {
	// we'll modify head via the copy
	temp := head
	for i := 0; i < pos-2; i++ {
		// temp points to (n-1)th Node
		temp = temp.next
	}
	nodeToDelete := temp.next     // (n)th Node
	temp.next = nodeToDelete.next // we'll have deleted the target Node
	return head
}
func reverseIterative(head *Node) *Node {
	var previous, current, next *Node = nil, head, nil
	for current != nil {
		next = current.next
		fmt.Println("previous => ", previous, "\n", "current => ", current, "\n", "next => ", next, "\n\n")
		current.next = previous
		previous = current
		current = next
	}
	head = previous
	return head
}
func main() {
	head := new(Node)
	head = insertAtEnd(head, 12)
	head = insertAtEnd(head, 13)
	head = insertAtEnd(head, 14)
	head = insertAtEnd(head, 15)
	head = insertAtEnd(head, 16)
	head = insertAtEnd(head, 17)
	print(head)
	fmt.Println(head)
	fmt.Println("Reversing elements")
	head = reverseIterative(head)
	// fmt.Println(head)
	// fmt.Println(head.next)
	// fmt.Println(head.next.next)
	// fmt.Println(head.next.next.next)
	// fmt.Println(head.next.next.next.next)
	// fmt.Println(head.next.next.next.next.next)
	// fmt.Println(head.next.next.next.next.next.next)
	// fmt.Println(head.next.next.next.next.next.next.next)
	print(head)
}
