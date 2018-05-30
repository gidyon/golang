package lists

import "fmt"

func insertAtBeginning(head *Node, data int) (newHead *Node) {
	newHead = new(Node)
	newHead.data = data
	newHead.next = head
	return
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
