package heap

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {

	t.Run("add", func(t *testing.T) {

		h := NewMaxHeap(20)

		for i := 1; i < 11; i++ {
			h.Add(i)
		}
		fmt.Printf("%v", h.heap)

	})

}
