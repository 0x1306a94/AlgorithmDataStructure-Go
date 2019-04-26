package list

import (
	"fmt"
	"testing"
)

func TestArrayList(t *testing.T) {

	list := New(10)

	for i := 0; i < 100; i++ {
		list.Add(i)
	}

	fmt.Println(list)

	for i := 0; i < 50; i++ {
		list.Remove(0)
	}

	fmt.Println(list)

}