package stack

import (
	"testing"
)

func Test1(t *testing.T) {

	arrStack := NewArrayStack()

	t.Run("Push", func(t *testing.T) {
		arrStack.Push("11")
		t.Log(arrStack.Peek())

	})

}
