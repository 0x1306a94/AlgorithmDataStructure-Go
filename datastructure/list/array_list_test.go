package list

import (
	"fmt"
	"testing"
)

func TestArrayList(t *testing.T) {

	var list List
	list = NewArrayList(10)
	testList(t, list)
	list.Clear()

	for i := 0; i < 10; i++ {
		list.Add(i)
	}
	fmt.Println(list)
	for i := 10; i < 15; i++ {
		list.Add(i)
	}
	{
		expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
		for i, e := range expected {
			if list.Get(i) != e {
				t.Fatal("测试未通过")
			}
		}
	}

	fmt.Println(list)

	for i := 0; i < 10; i++ {
		list.Remove(0)
	}

	{
		expected := []int{10, 11, 12, 13, 14}
		for i, e := range expected {
			if list.Get(i) != e {
				t.Fatal("测试未通过")
			}
		}
	}

	fmt.Println(list)

}
