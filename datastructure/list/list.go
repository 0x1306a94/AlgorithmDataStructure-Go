package list

// 元素索引未找到
const ElementNotFound int = -1

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
