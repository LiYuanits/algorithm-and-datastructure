package array

import (
	"errors"
	"fmt"
)

type Array struct {
	data []interface{}
	size int
}

// 实例化数组
func NewArray(capacity int) *Array {
	return &Array{
		data: make([]interface{}, capacity),
		size: 0,
	}
}

func (a *Array) Size() int {
	return a.size
}

func (a *Array) GetCapacity() int {
	return len(a.data)
}

func (a *Array) IsEmpty() bool {
	return a.size == 0
}

// 插入元素
func (a *Array) Add(index int, e interface{}) error {
	// 判断索引
	if index < 0 || index > a.size {
		return errors.New("索引参数错误")
	}
	// 扩容数组，以两倍扩容
	if a.size == len(a.data) {
		a.resize(a.size * 2)
	}

	// 插入元素,判断索引是否小于数组长度如果是就代表就代表是往数组中间插入元素
	// 要把数组整体往后移动，
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = e
	a.size++
	return nil
}

// 往头部插入元素
func (a *Array) AddFirst(e interface{}) error {
	err := a.Add(0, e)
	if err != nil {
		return err
	}
	return nil
}

// 往尾部插入元素
func (a *Array) AddLast(e interface{}) error {
	err := a.Add(a.size, e)
	if err != nil {
		return err
	}

	return nil
}

// 数组扩容
func (a *Array) resize(capacity int) {
	newData := make([]interface{}, capacity)
	for i := 0; i < len(a.data); i++ {
		newData[i] = a.data[i]
	}
	a.data = newData
}

func (a *Array) Remove(index int) (interface{}, error) {

	if index < 0 || index >= a.size {
		return nil, errors.New("删除元素索引错误")
	}

	// 要删除的元素
	ret := a.data[index]

	// 把数组元素往前移动
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}

	// 数组长度减一
	a.size--
	a.data[a.size] = 0
	return ret, nil
}

func (a *Array) RemoveFirst() (i interface{}, err error) {
	i, err = a.Remove(0)
	if err != nil {
		return nil, err
	}
	return
}

func (a *Array) RemoveLast() (i interface{}, err error) {
	i, err = a.Remove(a.size - 1)
	if err != nil {
		return nil, err
	}
	return
}

func (a *Array) RemoveElement(e int) error {
	i := a.Find(e)
	if i != -1 {
		_, err := a.Remove(i)
		if err != nil {
			return err
		}
	}
	return nil
}

// 查找当前元素是否数组中
func (a *Array) Find(e interface{}) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return i
		}
	}
	return -1
}

func (a *Array) Get(index int) (item interface{}, err error) {
	if index < 0 || index >= a.size {
		return nil, errors.New("索引错误")

	}
	return a.data[index], nil
}

func (a *Array) GetFirst() (item interface{}, err error) {
	item, err = a.Get(0)
	return
}

func (a *Array) GetLast() (item interface{}, err error) {
	item, err = a.Get(a.size - 1)
	return
}

func (a *Array) Set(index int, e int) {
	if index < 0 || index >= a.size {
		panic("参数错误")
	}
	a.data[index] = e
}

// 交换元素
func (a *Array) Swap(i int, j int) {

	if i < 0 || i >= a.size || j < 0 || j >= a.size {
		panic("索引错误")
	}
	t := a.data[i]
	a.data[i] = a.data[j]
	a.data[j] = t
}

//打印所有元素值
func (a *Array) PrintData() {
	for i := 0; i < a.size; i++ {
		if a.data[i] != nil {
			fmt.Print(a.data[i], " ")
		}
	}
	fmt.Println()
}
