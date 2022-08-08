package linked_list

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {

	linkList := NewLinkList()

	t.Run("size", func(t *testing.T) {
		linkList.Add(0, 1)
		linkList.Add(1, 2)
		linkList.Add(2, 3)
		linkList.Add(3, 4)
		//linkList.Reverse(linkList.dummyHead)
		T3(linkList.dummyHead.next)
		//linkList.Add(4, 4)
		//linkList.Add(5, 2)
		//linkList.RemoveElement(3)
	})

}

func T3(n *node) {
	if n.next == nil {
		return
	}
	T3(n.next)
	fmt.Println(n.e)
}
