package datastructure

import "algorithm-and-datastructure/linked_list"

// 集合特性:每个元素只能存在一次，不能重复。
// 代表集合不能添加重复元素
// 基于链表的集合

//
//type Set interface {
//	Add(e int)
//	Remove(e int)
//	Contains(e int) bool
//	GetSize() int
//	IsEmpty() bool
//}

type LinkedListSet struct {
	linkedList *linked_list.LinkedList
}

func NewLinkedListSet() *LinkedListSet {
	return &LinkedListSet{linkedList: linked_list.NewLinkList()}
}

//
func (l *LinkedListSet) IsEmpty() bool {
	return l.linkedList.IsEmpty()
}

func (l *LinkedListSet) GetSize() int {
	return l.linkedList.GetSize()
}

// 查询元素是否存在链表
func (l *LinkedListSet) Contains(e interface{}) bool {
	return l.linkedList.Contains(e)
}

// 往集合添加元素
func (l *LinkedListSet) Add(e interface{}) {
	click := l.Contains(e)
	if !click {
		l.linkedList.AddFirst(e)
	}
}

// 删除元素
func (l *LinkedListSet) Remove(e interface{}) {
	l.linkedList.RemoveElement(e)
}
