package heap

import "algorithm-and-datastructure/array"

// 堆的特性：堆中某个节点的值总是不大于其父亲节点的值
//用数组实现的最大堆
// 堆就是一颗天然的二叉树
// 使用二叉树实现的堆，也称为二叉堆

// 根节点的元素为整个堆最大的元素也称为最大堆
// 根节点的元素为整个堆最小的元素也称为最小堆

// 满二叉树：除了叶子节点，所有的节点左右孩子都不为空
// 完全二叉树：不是一颗满的二叉树，确实的部分是在整棵树的右下侧。

// 通过一个索引计算节点元素公式
// 从索引为1 开始记录元素
// parent(i) = i / 2
// left child (i) = 2 * i
// right child (i) = 2 * i + 1

// 从索引为0 开始记录元素计算公式
// parent(i) = （i-1） / 2
// left child (i) = 2 * i +1
// right child (i) = 2 * i + 1 +2

type MaxHeap struct {
	heap *array.Array
}

//
func NewMaxHeap(capacity int) *MaxHeap {
	return &MaxHeap{
		heap: array.NewArray(capacity),
	}
}

//
func (h *MaxHeap) GetSize() int {
	return h.heap.Size()
}

//
func (h *MaxHeap) IsEmpty() bool {
	return h.heap.IsEmpty()
}

// 给定索引计算出父亲节点的索引
func (h *MaxHeap) parent(index int) int {
	if index == 0 {
		panic("0是没有索引的")
	}
	return (index - 1) / 2
}

// left child (i) = 2 * i +1
// right child (i) = 2 * i +2

// 返回完全二叉树数组表示左孩子索引
func (h *MaxHeap) leftChild(index int) int {
	if index == 0 {
		panic("0是没有索引的")
	}
	return index*2 + 1
}

// 返回完全二叉树数组表示右孩子索引
func (h *MaxHeap) rightChild(index int) int {
	if index == 0 {
		panic("0是没有索引的")
	}
	return index*2 + 2
}

// 向堆中添加元素
// 把元素添加进数组的尾巴
// 然后执行元素上浮判断下是否符合最大堆的规则
func (h *MaxHeap) Add(e interface{}) {
	h.heap.AddLast(e)
	h.siftUp(h.heap.Size() - 1)

}

// 元素上浮
// 插入元素总是往数组最后面插入，为了保证堆的特性
// 要比较下当前插入元素跟插入元素所在位置的父亲节点是否小于父亲节点
// 如果不是就要交换元素的位置
func (h *MaxHeap) siftUp(k int) {
	for k > 0 {
		// 由于数组的元素存放的是 interface{}类型所，所以要类型断言才能比较元素
		item, _ := h.heap.Get(k)
		parentItem, _ := h.heap.Get(h.parent(k))
		if item.(int) < parentItem.(int) {
			return
		}
		h.heap.Swap(k, h.parent(k))
		k = h.parent(k)
	}
}

// 查找堆中最大的元素
func (h *MaxHeap) FindMax() interface{} {
	if h.heap.Size() == 0 {
		panic("堆是空的")
	}
	ret, _ := h.heap.Get(0)
	return ret
}

// 取出堆中最大的元素
func (h *MaxHeap) ExtractMax() interface{} {

	ret := h.FindMax()
	// 最后一个元素接替成为最大的元素
	h.heap.Swap(0, h.heap.Size()-1)

	//删除最后一个元素
	h.heap.RemoveLast()
	h.siftDown(0)
	return ret

}

// siftDown 元素下沉
// 由于删除元素的关系（就是取出堆最大的元素）,取出了一个元素就不是一个符合规则的最大堆了
// 所有有了这个算法
// 把数组最后一个元素放到最大节点的位置，然后删除最后那个元素
// 跟节点跟左右俩个孩子做比较
// 把最大元素的孩子节点跟根节点交换位置
// 一直到节点不大于要交换的元素为止

func (h *MaxHeap) siftDown(k int) {

	for {
		// 判断索引是否越界
		if h.leftChild(k) > h.heap.Size() {
			break
		}

		// 左孩子右孩子比较后的索引
		j := h.leftChild(k)

		// 判断是否有右孩子
		if j+1 < h.heap.Size() {
			// 拿到左右孩子
			leftItem, _ := h.heap.Get(j)
			rightItem, _ := h.heap.Get(j + 1)
			// 比较如果左孩子小于右孩子j等于右孩子
			if leftItem.(int) < rightItem.(int) {
				j = h.rightChild(k)
			}
		}

		item, _ := h.heap.Get(k)

		i, _ := h.heap.Get(j)
		if item.(int) >= i.(int) {
			break
		}
		// 交换元素
		h.heap.Swap(k, j)
	}

}
