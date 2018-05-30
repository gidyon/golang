package lists

func reverseIterative(head *Node) *Node {
	var previous, current, next *Node
	current = head
	for current != nil {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}
	head = previous
	return head
}
