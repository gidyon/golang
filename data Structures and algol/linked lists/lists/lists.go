package lists

import "fmt"

type Node struct {
	data int
	next *Node
}

func lengthOfList(head *Node) (length int) {
	temp := head
	for temp.next != nil {
		length++
		temp = temp.next
	}
	return
}

func print(head *Node) {
	fmt.Print("The linked list is: { ")
	for head.next != nil {
		fmt.Print(head.data, " ")
		head = head.next
	}
	fmt.Print(" }\n")
}
