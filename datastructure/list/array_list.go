package list

import "fmt"

const defaultCapacity int = 10

type ArrayList struct {
	size int
	capacity int

	elements []interface{}
}

func New(capacity int) *ArrayList {
	if capacity < defaultCapacity {
		capacity = defaultCapacity
	}
	return &ArrayList{
		size: 0,
		capacity: capacity,
		elements: make([]interface{}, capacity),
	}
}

func (a *ArrayList) Clear()  {
	a.size = 0
	if cap(a.elements) > defaultCapacity {
		a.elements = make([]interface{}, defaultCapacity)
	}
}

func (a *ArrayList) Size() int {
	return a.size
}

func (a *ArrayList) IsEmpty(e interface{}) bool {
	return a.size == 0
}

func (a *ArrayList) Contains(e interface{}) bool {
	return a.IndexOfElement(e) != ELEMENT_NOT_FOUND
}

func (a *ArrayList) Add(e interface{})  {
	a.AddForIndex(a.size, e)
}

func (a *ArrayList) AddForIndex(index int, e interface{}) {
	a.rangeCheckForAdd(index)

	a.ensureCapacity(a.size + 1)
	for i := a.size; i > a.size; i-- {
		a.elements[i] = a.elements[i - 1]
	}
	a.elements[index] = e
	a.size += 1
}

func (a *ArrayList) Get(index int) interface{} {
	a.rangeCheck(index)
	return a.elements[index]
}

func (a *ArrayList) Set(index int, e interface{}) interface{} {
	a.rangeCheck(index)
	old := a.elements[index]
	a.elements[index] = e
	return old
}

func (a *ArrayList) Remove(index int) interface{}  {
	a.rangeCheck(index)
	old := a.elements[index]
	for i := index + 1; i < a.size; i++ {
		a.elements[i-1] = a.elements[i]
	}
	a.size -= 1
	a.elements[a.size] = nil
	a.trim()
	return old
}

func (a *ArrayList) RemoveOfElement(e interface{})  {
	a.Remove(a.IndexOfElement(e))
}

func (a *ArrayList) IndexOfElement(e interface{}) int {
	for i := 0; i < a.size; i++ {
		if a.elements[i] == e {
			return i
		}
	}
	return ELEMENT_NOT_FOUND
}

func (a *ArrayList) rangeCheck(index int)  {
	if index < 0 || index >= a.size {
		panic("rangeCheck: 索引越界")
	}
}

func (a *ArrayList) rangeCheckForAdd(index int)  {
	if index < 0 || index > a.size {
		panic("rangeCheckForAdd: 索引越界")
	}
}

func (a *ArrayList) ensureCapacity(capacity int)  {
	oldCapacity := cap(a.elements)
	if oldCapacity >= capacity {
		return
	}

	// 新容量为旧容量的1.5倍
	newCapacity := oldCapacity + (oldCapacity >> 1)
	newElements := make([]interface{}, newCapacity)
	n := len(a.elements)
	copy(newElements, a.elements[0:n])
	a.elements = newElements
	fmt.Println(oldCapacity, " 扩容为 ", newCapacity)
}

func (a *ArrayList) trim()  {
	oldCapacity := cap(a.elements)
	newCapacity := oldCapacity >> 1
	if a.size > newCapacity || oldCapacity <= defaultCapacity {
		return
	}
	newElements := make([]interface{}, newCapacity)
	n := len(a.elements)
	copy(newElements, a.elements[0:n])
	a.elements = newElements
	fmt.Println(oldCapacity, " 缩容为 ", newCapacity)
}

func (a *ArrayList) String() string {
	return fmt.Sprint(a.elements)
}