package linked_list

import (
	"bytes"
	"errors"
	"fmt"
)

// 链表 nil -> 1 -> 2 -> 3 -> 4 -> 5 -> nil
//  __
// |e = nil
// | next = node
// |__
type LinkedList struct {
	dummyHead *node
	size      int
}

type node struct {
	e    interface{}
	next *node
}

// new 一个链表
func NewLinkList() *LinkedList {
	return &LinkedList{
		dummyHead: &node{}, // 设置虚拟头节点固定位置
	}
}

// new 一个链表节点
func NewLinkListNode(e interface{}, next *node) *node {
	return &node{
		e:    e,
		next: next,
	}
}

func (l *LinkedList) GetSize() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) Add(index int, e interface{}) error {

	// 单向链表
	if index < 0 || index > l.size {
		return errors.New("索引参数错误")
	}

	// 拿到头，头一般是不存数据的用作定位
	prev := l.dummyHead

	// 遍历获取到要插入的前一个节点,
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	// NewLinkListNode 构造node
	prev.next = NewLinkListNode(e, prev.next)
	l.size++
	return nil
}

func (l *LinkedList) AddFirst(e interface{}) {
	l.Add(0, e)
}

func (l *LinkedList) AddLast(e interface{}) {
	l.Add(l.size, e)
}

func (l *LinkedList) Get(index int) interface{} {
	if index < 0 || index >= l.size {
		panic("参数错误")
	}
	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.e
}

func (l *LinkedList) GetFirst() interface{} {
	return l.Get(0)
}

func (l *LinkedList) GetLast() interface{} {
	return l.Get(l.size - 1)
}

func (l *LinkedList) Set(index int, e interface{}) {

	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.e = e
}

func (l *LinkedList) Contains(e interface{}) bool {
	cur := l.dummyHead.next
	for i := 0; i < l.size; i++ {
		if cur.e == e {
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *LinkedList) Remove(index int) interface{} {

	// 找到要删除节点的前一个节点，把前一个节点的next指向删除节点的next
	if index < 0 || index >= l.size {
		panic("参数错误")
	}
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	retNode := prev.next
	prev.next = retNode.next
	retNode.next = nil
	l.size--
	return retNode.e
}

func (l *LinkedList) RemoveFirst() interface{} {
	return l.Remove(0)
}
func (l *LinkedList) RemoveLast() interface{} {
	return l.Remove(l.size - 1)
}

func (l *LinkedList) RemoveElement(e interface{}) {
	prev := l.dummyHead

	for {
		if prev.next.e == e {
			break
		}
		prev = prev.next
	}

	if prev.next.e != nil {
		// 要删除的节点
		delNode := prev.next
		// 要删除节点的下个节点
		prev.next = delNode.next
		// 清空节点
		delNode.next = nil
		// 要减一
		l.size--
	}
}

func (l *LinkedList) String() string {
	buffer := bytes.Buffer{}
	cur := l.dummyHead.next
	for i := 0; i < l.size; i++ {
		buffer.WriteString(fmt.Sprint(cur.e) + "->")
		cur = cur.next
	}
	buffer.WriteString("NULL")
	return buffer.String()
}

func (l *LinkedList) Reverse(n *node) {
	if n == nil {
		return
	}
	l.Reverse(n)
	fmt.Println(n.e)
}
