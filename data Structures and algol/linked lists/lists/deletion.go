package lists

import (
	"errors"
)

func deleteAtBeginning(head *Node) *Node {
	temp := head
	temp = temp.next
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

func deleteAtIndex(head *Node, data int, pos int) (*Node, error) {
	if pos == 1 {
		deleteAtBeginning(head)
		return head, nil
	}
	if pos == 0 {
		return nil, errors.New("cannot delete at position 0")
	}
	if pos > lengthOfList(head) {
		return nil, errors.New("position exceeds the length of the list")
	}
	// we'll modify head via the copy
	temp := head
	for i := 0; i < pos-2; i++ {
		// temp points to (n-1)th Node
		temp = temp.next
	}
	nodeToDelete := temp.next     // (n)th Node
	temp.next = nodeToDelete.next // we'll have deleted the target Node
	return head, nil
}
