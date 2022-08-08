package _map

import (
	"errors"
	"fmt"
)

//type Map1 interface {
//	Add(k interface{}, v interface{})
//
//	Remove(k interface{}) interface{}
//
//	Contains(k interface{}) bool
//
//	Get(k interface{}) interface{}
//
//	Set(K interface{}, v interface{})
//
//	GetSize() int
//
//	IsEmpty() bool
//}

// 基于二分搜索树的map映射(字典)：键值对存储数据
//

type BstMap struct {
	root *bstMapNode
	size int
}

// new 一个二叉树根节点
func NewBstMap() *BstMap {
	return &BstMap{
		root: &bstMapNode{},
		size: 0,
	}
}

type bstMapNode struct {
	key   int
	value interface{}
	left  *bstMapNode
	right *bstMapNode
}

// new 一个节点
func NewBstNde(key int, value interface{}) *bstMapNode {
	return &bstMapNode{
		key:   key,
		value: value,
		left:  &bstMapNode{},
		right: &bstMapNode{},
	}
}

func (b *BstMap) IsEmpty() bool {
	return b.size == 0
}

func (b *BstMap) GetSize() int {
	return b.size
}

// 往二分搜索树添加元素(key int , value interface{})
func (b *BstMap) Add(k int, v interface{}) {
	b.root = b.add(b.root, k, v)
}

// 返回插入新节点后二分搜索树的根
func (b *BstMap) add(node *bstMapNode, k int, v interface{}) *bstMapNode {

	// 如果这颗二分搜索树还没有节点，生成根节点
	// 这里是递归终止的条件
	if node.key == 0 {
		b.size++
		return NewBstNde(k, v)
	}

	// 如果插入的元素比根节点的元素小就是往左子树插入数据
	// 如果插入的元素比根节点小就是往右子树添加元素
	if node.key > k {
		// 这里就是更二分搜索树的数据
		node.left = b.add(node.left, k, v)
	} else if node.key < k {
		// 这里就是更二分搜索树的数据
		node.right = b.add(node.right, k, v)
	} else {
		node.value = v
	}
	return node
}

// 辅助函数查询当前key在哪个节点
func (b *BstMap) getNode(node *bstMapNode, key int) (n *bstMapNode, err error) {

	// 判断二分搜索树是否遍历到底
	if node.key == 0 {
		return nil, errors.New("没找到")
	}
	// 看当前节点元素跟目标元素是否相等
	if node.key == key {
		return node, nil
	} else if node.key > key {
		// 判断目标元素是否小于当前元素，是的话目标元素就在当二分搜索树的左节点
		return b.getNode(node.left, key)
	} else {
		return b.getNode(node.right, key)
	}
}

// 查找元素是否存在
func (b *BstMap) Contains(key int) bool {
	_, err := b.getNode(b.root, key)
	if err != nil {
		return false
	}
	return true
}

//	Get(k interface{}) interface{}
//
//	Set(K interface{}, v interface{})

func (b *BstMap) Get(k int) interface{} {
	ret, err := b.getNode(b.root, k)
	if err != nil {
		return nil
	}
	return ret.value
}

// 设置元素
func (b *BstMap) Set(k int, v interface{}) {
	ret, err := b.getNode(b.root, k)
	if err == nil {
		ret.value = v
	}
}

func (b *BstMap) minimum(node *bstMapNode) *bstMapNode {
	if node.left.key == 0 {
		return node
	}
	return b.minimum(node.left)
}

// 删除二分搜索树最小元素，并且返回最小元素
func (b *BstMap) RemoveMin() *bstMapNode {
	ret := b.minimum(b.root)
	b.root = b.removeMin(b.root)
	return ret
}

func (b *BstMap) removeMin(node *bstMapNode) *bstMapNode {
	// node.left.e  代表着已经遍历到底了
	if node.left.key == 0 {
		rightNode := node.right
		node.right = nil
		b.size--
		return rightNode
	}
	node.left = b.removeMin(node.left)
	return node
}

// 删除二分搜索树任意元素
func (b *BstMap) Remove(e int) {
	root := b.remove(b.root, e)
	fmt.Println(root)
}

func (b *BstMap) remove(node *bstMapNode, e int) *bstMapNode {

	if node.key == 0 {
		return nil
	}

	// 要删除的元素小于当前节点元素就往左子树找
	if node.key > e {
		node.left = b.remove(node.left, e)
		return node
	} else if node.key < e {
		// 要删除元素大于当前节点元素往右子树找
		node.right = b.remove(node.right, e)
		return node
	} else {

		// 如果当前节点等于要删除的元素

		// 要删除的节点只有右孩子
		if node.left.key == 0 {
			rightNode := node.right
			node.right = nil
			b.size--
			return rightNode
		}

		// 要删除的节点只有左孩子
		if node.right.key == 0 {
			leftNode := node.left
			node.left = nil
			b.size--
			return leftNode
		}

		// 待删除节点左右节点不为空
		// 找到待删除节点的最小节点，即左子树
		// 用这个节点顶替待删除节点的位置

		// 找到最小元素
		successor := b.minimum(node.right)
		successor.right = b.RemoveMin()
		successor.left = node.left
		node.left = nil
		node.right = nil
		return successor

	}

}
