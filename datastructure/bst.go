package datastructure

import (
	"fmt"
	"github.com/emirpasic/gods/queues/linkedlistqueue"
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

// 二分搜索树 Binary Search Tree

// 二叉树 动态数据结构，除了元素还有左子树跟右子树，右子树比左子树大
// 二叉树每个节点最多有两个孩子

/*      28              15
      /    \
     /      \
    /        \
   16        30
   /\         /\
  /  \       /  \
 /    \     /    \
13     22   29   42*/

// 前序遍历 28 16 13 22 30 29 42
// 中序遍历 13 16 22 28 29 30 42
// 后序遍历 13 22 16 29 42 30 28

/**
	二分搜索树，条件
    1、是一颗二叉树
    2、二分搜索树的每个节点的值：
       大于其左子树的所有节点的值
	   小于其右子树的所有节点的值
    3、每一颗子树也是二分搜索树
    4、 存储的元素必须有可比较性的
*/

type Bst struct {
	root *bstNode
	size int
}

type bstNode struct {
	e     int
	left  *bstNode
	right *bstNode
}

// new 一个节点
func NewBstNde(e int) *bstNode {
	return &bstNode{
		e:     e,
		left:  &bstNode{},
		right: &bstNode{},
	}
}

// new 一个二叉树根节点
func NewBst() *Bst {
	return &Bst{
		root: &bstNode{},
		size: 0,
	}
}

// 往二分搜索树添加元素
func (b *Bst) Add(e int) {
	b.root = b.add(b.root, e)
}

// 返回插入新节点后二分搜索树的根
func (b *Bst) add(node *bstNode, e int) *bstNode {

	// 如果这颗二分搜索树还没有节点，生成根节点
	// 这里是递归终止的条件
	if node.e == 0 {
		b.size++
		return NewBstNde(e)
	}

	// 如果插入的元素比根节点的元素小就是往左子树插入数据
	// 如果插入的元素比根节点小就是往右子树添加元素
	if node.e > e {
		// 这里就是更二分搜索树的数据
		node.left = b.add(node.left, e)
	} else if node.e < e {
		// 这里就是更二分搜索树的数据
		node.right = b.add(node.right, e)
	}
	return node
}

func (b *Bst) GetSize() int {
	return b.size
}

func (b *Bst) IsEmpty() bool {
	return b.size == 0
}

// 查询元素是否存在二分搜索树
// 递归
func (b *Bst) Contains(e int) bool {

	return b.contains(b.root, e)

	// 判断二分搜索树是否遍历到底
	//if node.e == 0 {
	//	return false
	//}
	//
	//// 判断当前节点元素是否相等
	//if node.e == e {
	//	return true
	//}
	//
	//if node.e > e {
	//	b.add(node.left, e)
	//} else if node.e < e {
	//	b.add(node.right, e)
	//}
	//
	//return false
}

// 看以node为根节点的二分搜索树中是否包含元素e， 递归算法
func (b *Bst) contains(node *bstNode, e int) bool {

	// 判断二分搜索树是否遍历到底
	if node.e == 0 {
		return false
	}

	// 看当前节点元素跟目标元素是否相等
	if node.e == e {
		return true
	} else if node.e > e {
		// 判断目标元素是否小于当前元素，是的话目标元素就在当二分搜索树的左节点
		return b.contains(node.left, e)
	} else {

		return b.contains(node.right, e)
	}

}

// 二分搜索树遍历
// 树结构的遍历
// 前序遍历: 从跟节点开始，先访问跟节点再访问左子树，再访问右子树  根->左 -> 右
// 先遍历左子树再遍历右子树

func (b *Bst) PreOrder() {
	b.preOrder(b.root)

}

func (b *Bst) preOrder(node *bstNode) {
	if node.e == 0 {
		return
	}
	fmt.Println(node.e)
	b.preOrder(node.left)
	b.preOrder(node.right)
}

// 前序遍历非递归实现
func (b *Bst) PreOrderdl() {

	// 使用栈这种数据结构
	// 先把右子树push进去再push左子树，直到叶子节点为止
	s := linkedliststack.New()
	s.Push(b.root)
	for {
		if !s.Empty() {
			cur, ok := s.Pop()
			if !ok {
				fmt.Println("取出数据错误")
			}
			c := cur.(*bstNode)
			fmt.Println(c.e)

			if c.right.e != 0 {
				s.Push(c.right)
			}
			if c.left.e != 0 {
				s.Push(c.left)
			}
		} else {
			break
		}
	}
}

// 中序遍历
// 左 -> 根 -> 右
func (b *Bst) InOrder() {
	b.inOrder(b.root)
}

func (b *Bst) inOrder(node *bstNode) {

	/*   28              15
	   /    \
	  /      \
	 /        \
	16        30
	*/

	if node.e == 0 {
		return
	}
	b.inOrder(node.left)
	fmt.Println(node.e)
	b.inOrder(node.right)
}

// 后序遍历 左 -> 右 ->根
func (b *Bst) AfterOrder() {
	b.afterOrder(b.root)
}

func (b *Bst) afterOrder(node *bstNode) {
	if node.e == 0 {
		return
	}
	b.afterOrder(node.left)
	b.afterOrder(node.right)
	fmt.Println(node.e)
}

//  二分搜索树的前序中序后序遍历都是深度优先遍历
// 层序遍历别称广度遍历

func (b *Bst) LevelOrder() {
	// 使用队列数据结构先进先出
	// 先把跟节点放进去，在把左子树右子树放进去知道遍历到叶子节点
	q := linkedlistqueue.New()
	q.Enqueue(b.root)
	for {
		// 如果不等于空就代表队列里面还有元素
		if !q.Empty() {
			cur, ok := q.Dequeue()
			if !ok {
				fmt.Println("取数据错误")
				return
			}
			c := cur.(*bstNode)
			fmt.Println(c.e)
			if c.left.e != 0 {
				q.Enqueue(c.left)
			}
			if c.right.e != 0 {
				q.Enqueue(c.right)
			}
		} else {
			break
		}
	}
}

// 二分搜索树最小值,往左子树找，直到左子树为空
//  二分搜索树最大值，往右子树找，直到右子树为空

// 找出最小的元素
func (b *Bst) Minimum() {

	if b.size == 0 {
		panic("二分搜索树没有值")
	}
	ret := b.minimum(b.root)
	fmt.Printf("%v", ret)
}

func (b *Bst) minimum(node *bstNode) *bstNode {
	if node.left.e == 0 {
		return node
	}
	return b.minimum(node.left)
}

// 找出最大的元素
func (b *Bst) Maximum() {

	if b.size == 0 {
		panic("二分搜索树没有值")
	}
	ret := b.maximum(b.root)
	fmt.Printf("%v", ret)
}

func (b *Bst) maximum(node *bstNode) *bstNode {
	if node.right.e == 0 {
		return node
	}
	return b.maximum(node.right)
}

// 删除二分搜索树最小元素，并且返回最小元素
func (b *Bst) RemoveMin() *bstNode {
	ret := b.minimum(b.root)
	b.root = b.removeMin(b.root)
	return ret
}

func (b *Bst) removeMin(node *bstNode) *bstNode {
	// node.left.e  代表着已经遍历到底了
	if node.left.e == 0 {
		rightNode := node.right
		node.right = nil
		b.size--
		return rightNode
	}
	node.left = b.removeMin(node.left)
	return node
}

// 删除二分搜索树最小元素，并且返回最小元素

func (b *Bst) RemoveMax() *bstNode {
	// 查找出最大的元素
	ret := b.maximum(b.root)
	b.root = b.removeMax(b.root)
	return ret
}

func (b *Bst) removeMax(node *bstNode) *bstNode {
	// node.left.e  代表着已经遍历到底了
	// 有可能当前节点的左子树有值把他放在右子树
	if node.right.e == 0 {
		leftNode := node.left
		node.right = nil
		b.size--
		return leftNode
	}
	node.right = b.removeMax(node.right)
	return node
}

// 删除二分搜索树任意元素
func (b *Bst) Remove(e int) {
	root := b.remove(b.root, e)
	fmt.Println(root)
}

func (b *Bst) remove(node *bstNode, e int) *bstNode {

	if node.e == 0 {
		return nil
	}

	// 要删除的元素小于当前节点元素就往左子树找
	if node.e > e {
		node.left = b.remove(node.left, e)
		return node
	} else if node.e < e {
		// 要删除元素大于当前节点元素往右子树找
		node.right = b.remove(node.right, e)
		return node
	} else {

		// 如果当前节点等于要删除的元素

		// 要删除的节点只有右孩子
		if node.left.e == 0 {
			rightNode := node.right
			node.right = nil
			b.size--
			return rightNode
		}

		// 要删除的节点只有左孩子
		if node.right.e == 0 {
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

func (b *Bst) String() string {
	fmt.Println("执行了string")

	b.PreOrder()
	return "string"
}
