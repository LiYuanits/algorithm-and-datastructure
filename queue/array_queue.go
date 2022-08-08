package queue

import (
	"bytes"
	"datastructures/array"
	"fmt"
)

// 线性数据结构
// 普通队列特性：先进先出。
// 只能从一端（队尾）添加元素，只能从另一端（队首）取出元素

type Queue interface {
	// 入队
	EnQueue(e interface{}) bool
	// 出队
	Dequeue() (i interface{}, err error)

	GetSize() int
	IsEmpty() bool
}

type ArrayQueue struct {
	array *array.Array
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		array: array.NewArray(capacity),
	}
}

func (q *ArrayQueue) GetSize() int {
	return q.array.Size()
}

func (q *ArrayQueue) IsEmpty() bool {

	return q.array.IsEmpty()
}

// 入队
func (q *ArrayQueue) EnQueue(e interface{}) error {

	if q.array.Size() == 0 {

		return q.array.AddFirst(e)

	} else {
		return q.array.AddLast(e)
	}

}

// 出队
func (q *ArrayQueue) Dequeue() (i interface{}, err error) {
	return q.array.RemoveFirst()
}

func (q *ArrayQueue) String() string {
	buffer := bytes.Buffer{}
	for i := 0; i < q.array.Size(); i++ {
		r, _ := q.array.Get(i)
		buffer.WriteString(fmt.Sprint(r) + "->")
	}
	buffer.WriteString("NULL")
	return buffer.String()
}
