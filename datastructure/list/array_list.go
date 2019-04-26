package list

import (
	"fmt"
	"strings"
)

type ArrayList struct {
	AbstractList
	// 存放元素slice
	elements []interface{}
}

// 创建 ArrayList, capacity 参数为初始容量, 如果小于10 则默认为10
func NewArrayList(capacity int) *ArrayList {
	if capacity < DefaultArrayListCapacity {
		capacity = DefaultArrayListCapacity
	}
	return &ArrayList{
		AbstractList: AbstractList{
			size: 0,
		},
		elements: make([]interface{}, capacity),
	}
}

// 清除ArrayList所有元素
func (a *ArrayList) Clear() {
	a.size = 0
	if cap(a.elements) > DefaultArrayListCapacity {
		a.elements = make([]interface{}, DefaultArrayListCapacity)
	}
}

// 检查 ArrayList 是否包含指定元素
func (a *ArrayList) Contains(e interface{}) bool {
	return a.IndexOfElement(e) != ElementNotFound
}

// 添加一个元素到 ArrayList 中
func (a *ArrayList) Add(e interface{}) {
	a.AddForIndex(a.size, e)
}

// 添加一个元素到 ArrayList 指定索引, 请注意索引是否越界
func (a *ArrayList) AddForIndex(index int, e interface{}) {
	a.rangeCheckForAdd(index)

	a.ensureCapacity(a.size + 1)
	for i := a.size; i > a.size; i-- {
		a.elements[i] = a.elements[i-1]
	}
	a.elements[index] = e
	a.size += 1
}

// 获取ArrayList中指定索引元素, 请注意索引是否越界
func (a *ArrayList) Get(index int) interface{} {
	a.rangeCheck(index)
	return a.elements[index]
}

// 设置ArrayList中指定索引元素, 请注意索引是否越界
func (a *ArrayList) Set(index int, e interface{}) interface{} {
	a.rangeCheck(index)
	old := a.elements[index]
	a.elements[index] = e
	return old
}

// 移除ArrayList中指定索引元素, 请注意索引是否越界
func (a *ArrayList) Remove(index int) interface{} {
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

// 移除ArrayList中指定元素, 如有则移除
func (a *ArrayList) RemoveOfElement(e interface{}) {
	if index := a.IndexOfElement(e); index != ElementNotFound {
		a.Remove(index)
	}
}

// 获取指定元素在ArrayList中对应索引, 如没有则返回 ElementNotFound
func (a *ArrayList) IndexOfElement(e interface{}) int {
	for i := 0; i < a.size; i++ {
		if a.Equal != nil {
			if a.Equal(a.elements[i], e) {
				return i
			}
		} else if a.elements[i] == e {
			return i
		}
	}
	return ElementNotFound
}

// ArrayList 扩容, 以1.5 倍扩容
func (a *ArrayList) ensureCapacity(capacity int) {
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

// ArrayList 扩容, 以原始容量0.5 倍缩容
func (a *ArrayList) trim() {
	oldCapacity := cap(a.elements)
	newCapacity := oldCapacity >> 1
	if a.size > newCapacity || oldCapacity <= DefaultArrayListCapacity {
		return
	}
	newElements := make([]interface{}, newCapacity)
	n := len(a.elements)
	copy(newElements, a.elements[0:n])
	a.elements = newElements
	fmt.Println(oldCapacity, " 缩容为 ", newCapacity)
}

func (a *ArrayList) String() string {
	builder := &strings.Builder{}
	fmt.Fprintf(builder, "ArrayList: capacity=%v, size=%v, [", cap(a.elements), a.size)
	for i := 0; i < a.size; i++ {
		if i != 0 {
			builder.WriteString(", ")
		}
		fmt.Fprint(builder, a.elements[i])
	}
	builder.WriteString("]")
	return builder.String()
}
