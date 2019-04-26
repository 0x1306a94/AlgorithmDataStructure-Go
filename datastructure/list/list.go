package list

import "fmt"

type List interface {
	Clear()

	Size() int

	IsEmpty() bool

	Contains(e interface{}) bool

	Add(e interface{})

	AddForIndex(index int, e interface{})

	Get(index int) interface{}

	Set(index int, e interface{}) interface{}

	Remove(index int) interface{}

	RemoveOfElement(e interface{})

	IndexOfElement(e interface{}) int
}

type _AbstractList struct {
	// 已经存储的元素数量
	size int
	// 自定义元素判等函数
	Equal func(v1, v2 interface{}) bool
}

func (a *_AbstractList) Size() int {
	return a.size
}

func (a *_AbstractList) IsEmpty() bool {
	return a.size == 0
}

func (a *_AbstractList) outOfBounds(index int) {
	panic(fmt.Sprintf("Index: %v, Size: %v", index, a.size))
}

func (a *_AbstractList) rangeCheck(index int) {
	if index < 0 || index >= a.size {
		a.outOfBounds(index)
	}
}

func (a *_AbstractList) rangeCheckForAdd(index int) {
	if index < 0 || index > a.size {
		a.outOfBounds(index)
	}
}
