package _map

import "errors"

// 基于链表实现的map映射(字典)：键值对存储数据

type Map1 interface {
	Add(k interface{}, v interface{})

	Remove(k interface{}) interface{}

	Contains(k interface{}) bool

	Get(k interface{}) interface{}

	Set(K interface{}, v interface{})

	GetSize() int

	IsEmpty() bool
}

type LinkedListMap struct {
	dummyHead *node
	size      int
}

// new 一个链表
func NewLinkListMap() *LinkedListMap {
	return &LinkedListMap{
		dummyHead: &node{}, // 设置虚拟头节点固定位置
	}
}

// 链表节点
type node struct {
	key   interface{}
	value interface{}
	next  *node
}

// new 一个链表节点
func NewLinkListMapNode(key interface{}, value interface{}, next *node) *node {
	return &node{
		key:   key,
		value: value,
		next:  next,
	}
}

func (l *LinkedListMap) GetSize() int {
	return l.size
}

func (l *LinkedListMap) IsEmpty() bool {
	return l.size == 0
}

// 查找节点
func (l *LinkedListMap) GetNode(k interface{}) (n *node, err error) {
	cur := l.dummyHead.next
	for {
		if cur == nil {
			return nil, errors.New("节点不存在")
		}

		if cur.key == k {
			return cur, nil
		}

		cur = cur.next
	}
}

// 查询元素是否存在
func (l *LinkedListMap) Contains(k interface{}) bool {
	_, err := l.GetNode(k)
	if err != nil {
		return false
	}
	return true
}

// 获取node
func (l *LinkedListMap) Get(k interface{}) (v interface{}) {
	ret, err := l.GetNode(k)
	if err != nil {
		return nil
	}
	return ret.value
}

// 设置节点
func (l *LinkedListMap) Set(k interface{}, v interface{}) {
	ret, err := l.GetNode(k)
	if err != nil {
		panic("节点不存在")
	}
	ret.value = v
}

// 添加
func (l *LinkedListMap) Add(k interface{}, v interface{}) {
	node, err := l.GetNode(k)
	// 当前节点不存在代表添加新节点
	// 从头头部添加O(1)
	if err != nil {
		l.dummyHead.next = NewLinkListMapNode(k, v, l.dummyHead.next)
		l.size++
		return
	}
	// 表示当前节点存在要修改节点
	node.value = v
}

// 删除元素
func (l *LinkedListMap) Remove(k interface{}) interface{} {
	prev := l.dummyHead
	for {
		if prev.next == nil || prev.next.key == k {
			break
		}
		prev = prev.next
	}
	if prev.next != nil {
		retNode := prev.next
		prev.next = retNode.next
		retNode.next = nil
		l.size--
		return retNode.value
	}
	return nil
}
