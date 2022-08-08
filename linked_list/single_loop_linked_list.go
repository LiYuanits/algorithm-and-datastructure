package linked_list

import (
	"bytes"
	"fmt"
)

type SingleLoopLinkedList struct {
	dummyHead *SingleLoopLinkedListNode
	size      int
}

type SingleLoopLinkedListNode struct {
	e    interface{}
	next *SingleLoopLinkedListNode
}

func NewSingleLoopLinkedListNode(e interface{}, next *SingleLoopLinkedListNode) *SingleLoopLinkedListNode {
	return &SingleLoopLinkedListNode{
		e:    e,
		next: next,
	}
}

func NewSingleLoopLinkedList() *SingleLoopLinkedList {
	return &SingleLoopLinkedList{
		dummyHead: &SingleLoopLinkedListNode{},
	}
}

func (l *SingleLoopLinkedList) Add(index int, e interface{}) {

	if index < 0 || index > l.size {
		panic("参数错误")
	}

	prev := l.dummyHead

	if l.size == 0 {
		newNode := NewSingleLoopLinkedListNode(e, prev.next)
		l.dummyHead.next = newNode
		newNode.next = l.dummyHead.next
		l.size++
		return
	}

	// 插入头
	if l.size != 0 && index == 0 {
		newNode := NewSingleLoopLinkedListNode(e, prev.next)
		prev.next = newNode
		lastNode := l.findLastNode()
		lastNode.next = newNode
		l.size++
		return
	}

	// 插入尾巴
	if l.size != 0 && index == l.size {
		lastNode := l.findLastNode()
		lastNode.next = NewSingleLoopLinkedListNode(e, lastNode.next)
		l.size++
		return
	}

	// 中间插入元素
	for i := 0; i < l.size; i++ {
		prev = prev.next
	}
	prev.next = NewSingleLoopLinkedListNode(e, prev.next)
	l.size++

}

func (l *SingleLoopLinkedList) findLastNode() *SingleLoopLinkedListNode {
	prev := l.dummyHead.next
	for i := 0; i < l.size; i++ {
		prev = prev.next
	}
	return prev
}

func (l *SingleLoopLinkedList) String() string {
	buffer := bytes.Buffer{}
	cur := l.dummyHead.next
	for i := 0; i < l.size; i++ {
		buffer.WriteString(fmt.Sprint(cur.e) + "->")
		cur = cur.next
	}
	buffer.WriteString("NULL")
	return buffer.String()
}
