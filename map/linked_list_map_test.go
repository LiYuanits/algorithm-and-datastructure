package _map

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	listMap := NewLinkListMap()

	t.Run("add", func(t *testing.T) {

		listMap.Add("k1", "v1")
		listMap.Add("k2", "v3")
		listMap.Add("k3", "v3")
		fmt.Printf("%+v", listMap)

	})

}
