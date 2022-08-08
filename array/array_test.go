package array

import "testing"

func Test1(t *testing.T) {

	arr := NewArray(10)
	t.Run("Add", func(t *testing.T) {

		arr.Add(0, 1)
		arr.Add(1, 2)
	})
}
