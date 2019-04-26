package list

const ELEMENT_NOT_FOUND int = -1

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

