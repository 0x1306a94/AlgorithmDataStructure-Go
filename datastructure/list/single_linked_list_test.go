package list

import (
	"fmt"
	"testing"
)

func TestNewSingleLinkedList(t *testing.T) {

	var list List
	list = NewSingleLinkedList()

	testList(t, list)
	list.Clear()
	for i := 0; i < 10; i++ {
		list.Add(i)
	}
	// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
	fmt.Println(list)

	{
		expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		for i, e := range expected {
			if list.Get(i) != e {
				t.Fatal("测试未通过")
			}
		}
	}

	{
		// 0, 1, 2, 3, 4, 6, 7, 8, 9
		list.RemoveOfElement(5)
		// 0, 1, 3, 4, 6, 7, 8, 9
		list.RemoveOfElement(2)
		// 0, 1, 3, 4, 6, 7, 9
		list.RemoveOfElement(8)
		fmt.Println(list)
		expected := []int{0, 1, 3, 4, 6, 7, 9}
		for i, e := range expected {
			if list.Get(i) != e {
				t.Fatal("测试未通过")
			}
		}
	}
}
