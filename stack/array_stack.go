package stack

import "datastructures/array"

// 栈的特性：后进先出
//type Stack interface {
//	Push(e interface{})
//	Pop() interface{}
//	Peek() interface{}
//	GetSize() int
//	IsEmpty() bool
//}

type ArrayStack struct {
	array *array.Array
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{array: array.NewArray(10)}
}

func (a *ArrayStack) Push(e interface{}) error {
	err := a.array.AddLast(e)
	return err
}

func (a *ArrayStack) Pop() (item interface{}, err error) {
	item, err = a.array.RemoveLast()
	return
}

func (a *ArrayStack) Peek() (item interface{}, err error) {
	return a.array.GetLast()
}

func (a *ArrayStack) GetSize() int {
	return a.array.Size()
}

func (a *ArrayStack) IsEmpty() bool {
	return a.array.IsEmpty()
}
