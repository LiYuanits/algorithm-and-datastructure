package linked_list

import (
	"bytes"
	"errors"
	"fmt"
)

// 双向链表

type DoubleLinkedList struct {
	dummyHead *doubleNode
	size      int
}

// nil -> 1 -> 2 -> 3

// 实例化双向链表
func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		dummyHead: &doubleNode{},
	}
}

type doubleNode struct {
	e          interface{}
	next, prev *doubleNode
}

func NewDoubleNode(e interface{}, next *doubleNode, prev *doubleNode) *doubleNode {
	return &doubleNode{
		e:    e,
		next: next,
		prev: prev,
	}
}

func (l *DoubleLinkedList) GetSize() int {
	return l.size
}

func (l *DoubleLinkedList) IsEmpty() bool {
	return l.size == 0
}

// 添加元素
func (l *DoubleLinkedList) Add(index int, e interface{}) error {
	// 判断索引是否越界
	if index < 0 || index > l.size {
		return errors.New("添加元素索引错误")
	}
	// 头部添加元素
	// 尾部添加元素
	// 身体中间添加元素
	// 拿到头，头一般是不存数据的用作定位
	prev := l.dummyHead
	// 遍历获取到要插入的前一个节点,
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	ii := prev.e
	// NewLinkListNode 构造node
	prev.next = NewDoubleNode(e, prev.next, prev.prev)
	prev.prev = NewDoubleNode(ii, prev.next, prev.next)
	l.size++
	return nil

}

func (l *DoubleLinkedList) String() string {
	buffer := bytes.Buffer{}
	cur := l.dummyHead.next
	for i := 0; i < l.size; i++ {
		buffer.WriteString(fmt.Sprint(cur.e) + "->")
		cur = cur.next
	}
	buffer.WriteString("NULL")
	return buffer.String()
}
