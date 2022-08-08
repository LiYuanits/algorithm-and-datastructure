package queue

import "testing"

func Test1(t *testing.T) {

	arrayQueue := NewArrayQueue(10)

	t.Run("EnQueue", func(t *testing.T) {
		arrayQueue.EnQueue("6666")
		arrayQueue.EnQueue("111")

		t.Log(arrayQueue)

	})

}
