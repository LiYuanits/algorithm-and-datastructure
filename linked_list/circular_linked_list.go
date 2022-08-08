package linked_list

// 单循环链表

type CircularLinkedList struct {
	size int
}

func NewCircularLinkedList() *CircularLinkedList {

	return &CircularLinkedList{}

}

type circularNode struct {
	e    interface{}
	next *circularNode
}

func NewCircularNode(e interface{}, next *circularNode) *circularNode {
	return &circularNode{
		e:    e,
		next: nil,
	}
}
