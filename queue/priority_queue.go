package queue

// 优先队列
// 优先队列：出队顺序和入队顺序无关，和优先级相关
// 可以使用不同的底层实现

type PQueue interface {
	Enqueue(e interface{}) // 入队
	Dequeue() interface{}  // 出队  应该是优先级最高的元素优先出队
	GetFront() interface{} // 获取头元素   优先级最高的元素
	GetSize() int          // 获取队列长度
	IsEmpty() bool         // 判断队列是否为空

}
