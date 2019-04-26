package list

import (
	"fmt"
	"testing"
)

func testList(t *testing.T, list List) {
	list.Add(11)
	list.Add(22)
	list.Add(33)
	list.Add(44)
	// [11, 22, 33, 44]
	fmt.Println(list)

	list.AddForIndex(0, 55) // [55, 11, 22, 33, 44]
	fmt.Println(list)

	list.AddForIndex(2, 66) // [55, 11, 66, 22, 33, 44]
	fmt.Println(list)

	list.AddForIndex(list.Size(), 77) // [55, 11, 66, 22, 33, 44, 77]
	fmt.Println(list)

	list.Remove(0) // [11, 66, 22, 33, 44, 77]
	fmt.Println(list)

	list.Remove(2) // [11, 66, 33, 44, 77]
	fmt.Println(list)

	list.Remove(list.Size() - 1) // [11, 66, 33, 44]
	fmt.Println(list)

	if list.IndexOfElement(44) != 3 {
		t.Fatal("测试未通过")
	}

	if list.IndexOfElement(22) != ElementNotFound {
		t.Fatal("测试未通过")
	}

	if !list.Contains(33) {
		t.Fatal("测试未通过")
	}

	if list.Get(0) != 11 {
		t.Fatal("测试未通过")
	}

	if list.Get(1) != 66 {
		t.Fatal("测试未通过")
	}

	if list.Get(list.Size()-1) != 44 {
		t.Fatal("测试未通过")
	}
}
