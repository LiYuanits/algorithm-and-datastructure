package datastructure

// 集合特性:每个元素只能存在一次，不能重复。
// 代表集合不能添加重复元素
// 二分搜索树集合，底层数据结构是二分搜索树

type Set interface {
	Add(e int)
	Remove(e int)
	Contains(e int) bool
	GetSize() int
	IsEmpty() bool
}

type Tset struct {
	bst *Bst
}

func NewTset() *Tset {
	return &Tset{bst: NewBst()}
}

// 集合是否为空
func (s *Tset) IsEmpty() bool {
	return s.bst.IsEmpty()
}

// 集合的长度
func (s *Tset) GetSize() int {
	return s.bst.GetSize()
}

// 查询元素是否存在
func (s *Tset) Contains(e int) bool {
	return s.bst.Contains(e)
}

// 往集合添加元素
func (s *Tset) Add(e int) {
	s.bst.Add(e)
}

// 删除集合中的元素
func (s *Tset) Remove(e int) {
	s.bst.Remove(e)

}
