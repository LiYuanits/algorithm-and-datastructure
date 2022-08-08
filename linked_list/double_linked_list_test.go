package linked_list

import "testing"

func Test(t *testing.T) {

	l := NewDoubleLinkedList()

	t.Run("Add", func(t *testing.T) {
		l.Add(0, 1)
		l.Add(1, 2)
		l.Add(2, 3)
		l.Add(3, 4)

		t.Log(l)
		//   1  - > 2 -> 3 -> 4 ->nil
		//   1  <-2 <- 3 <- 4

	})

}
