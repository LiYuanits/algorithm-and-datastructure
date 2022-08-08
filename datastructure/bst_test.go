package datastructure

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	b := NewBst()

	t.Run("test1", func(t *testing.T) {

		b.Add(28)
		b.Add(16)
		b.Add(30)
		b.Add(13)
		b.Add(22)
		b.Add(29)
		b.Add(42)

		// 前序遍历
		//b.PreOrder()

		// 中序遍历
		// b.InOrder()

		// 后序遍历
		// b.AfterOrder()

		// 28 16 30 13 22 29 42   层序遍历 广度优先遍历
		// b.LevelOrder()

		// 前序遍历非递归实现  // 28 16 13 22 30 29 42
		// b.PreOrderdl()

		// b.Minimum()

		/*      28              15
		      /    \
		     /      \
		    /        \
		   16        30
		   /\         /\
		  /  \       /  \
		 /    \     /    \
		13     22   29   42*/

		// b.Maximum()
		b.Remove(42)
		fmt.Println(b)
	})

}

/*      28              15
      /    \
     /      \
    /        \
   16        30
   /\         /\
  /  \       /  \
 /    \     /    \
13     22   29   42*/

/**/
