package list

import (
	"fmt"
	"testing"
)

func TestNewCircleLinkedList(t *testing.T) {
	list := NewCircleLinkedList()

	testList(t, list)
	list.Clear()

	for i := 0; i <= 8; i++ {
		list.Add(i)
	}
	// 指向头结点
	list.Rest()
	fmt.Println(list)
	for !list.IsEmpty() {
		fmt.Println(list.Next())
		fmt.Println(list.Next())
		//list.RemoveCurrent()
		fmt.Println(list.RemoveCurrent())
	}
	fmt.Println(list)

}
