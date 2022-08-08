package queue

import (
	"bytes"
	"fmt"
)

// 队列特性先进先出
// 实现了这个数据结构的特性

// 其实这个队列我也只需要使用AddLast() 添加元素 , RemoveFirst() 删除元素就行了

type QueueLinkedList struct {
	head *queueNode
	tail *queueNode
	size int
}

type queueNode struct {
	e    interface{}
	next *queueNode
}

func NewQueueLinkedList() *QueueLinkedList {
	return &QueueLinkedList{
		head: &queueNode{},
		tail: &queueNode{},
		size: 0,
	}
}

func newQueueNode(e interface{}) *queueNode {
	return &queueNode{
		e:    e,
		next: nil,
	}
}

func (l *QueueLinkedList) GetSize() int {
	return l.size
}

func (l *QueueLinkedList) IsEmpty() bool {
	return l.size == 0
}

// 入队 队尾入队
func (l *QueueLinkedList) EnQueue(e interface{}) {
	// 如果是第一个元素进来那这个链表的头和尾都应该是执行同一个元素

	if l.size == 0 {
		l.tail = newQueueNode(e)
		l.head = l.tail
	} else {
		l.tail.next = newQueueNode(e)
		l.tail = l.tail.next
	}
	l.size++
}

// 出队 队首出队
func (l *QueueLinkedList) Dequeue() interface{} {
	if l.IsEmpty() {
		panic("没有元素啊")
	}
	retNode := l.head
	l.head = retNode.next
	c := retNode.e
	retNode = nil
	if l.head == nil {
		l.tail = nil
	}
	l.size--
	return c
}

func (l *QueueLinkedList) GetFront() interface{} {
	if l.IsEmpty() {
		panic("Queue 没有元素啊")
	}
	return l.head.e
}

func (l *QueueLinkedList) String() string {
	buffer := bytes.Buffer{}
	cur := l.head
	buffer.WriteString("Queue:front: ")
	for i := 0; i < l.size; i++ {
		buffer.WriteString(fmt.Sprint(cur.e) + "->")
		cur = cur.next
	}
	buffer.WriteString("NULL: ")
	buffer.WriteString("Queue:tail")
	return buffer.String()
}
