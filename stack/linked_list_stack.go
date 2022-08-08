package stack

import "datastructures/linked_list"

//  栈后进先出

type Stack interface {
	Push(e interface{})
	Pop() interface{}
	Peek() interface{}
	GetSize() int
	IsEmpty() bool
}

type LinkedListStack struct {
	linkedlist *linked_list.LinkedList
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{linkedlist: linked_list.NewLinkList()}
}

func (s *LinkedListStack) Push(e interface{}) {
	s.linkedlist.AddFirst(e)
}

func (s *LinkedListStack) Pop() interface{} {
	return s.linkedlist.RemoveFirst()
}

func (s *LinkedListStack) Peek() interface{} {
	return s.linkedlist.GetFirst()
}

func (s *LinkedListStack) GetSize() int {
	return s.linkedlist.GetSize()
}

func (s *LinkedListStack) IsEmpty() bool {
	return s.linkedlist.IsEmpty()
}
