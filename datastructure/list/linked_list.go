package list

import (
	"fmt"
	"strings"
)

type LinkedList struct {
	SingleLinkedList
	last *_Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		SingleLinkedList: SingleLinkedList{
			_AbstractList: _AbstractList{
				size: 0,
			},
			first: nil,
		},
		last: nil,
	}
}

func (l *LinkedList) Clear() {
	l.size = 0
	l.first = nil
	l.last = nil
}

func (l *LinkedList) Contains(e interface{}) bool {
	return l.IndexOfElement(e) != ElementNotFound
}

func (l *LinkedList) Add(e interface{}) {
	l.AddForIndex(l.size, e)
}

func (l *LinkedList) AddForIndex(index int, e interface{}) {
	l.rangeCheckForAdd(index)
	// size == 0
	// index == 0
	if l.size == index {
		oldLast := l.last
		l.last = newNode(e, oldLast, nil)
		if oldLast == nil {
			l.first = l.last
		} else {
			oldLast.next = l.last
		}
	} else {
		next := l.nodeForIndex(index)
		prev := next.prev
		node := newNode(e, prev, next)
		next.prev = node
		if prev == nil { // index == 0
			l.first = node
		} else {
			prev.next = node
		}
	}
	l.size += 1
}

func (l *LinkedList) Get(index int) interface{} {
	l.rangeCheck(index)
	return l.nodeForIndex(index).element
}

func (l *LinkedList) Set(index int, e interface{}) interface{} {
	l.rangeCheck(index)
	node := l.nodeForIndex(index)
	old := node.element
	node.element = e
	return old
}

func (l *LinkedList) Remove(index int) interface{} {
	l.rangeCheck(index)
	node := l.nodeForIndex(index)
	prev := node.prev
	next := node.next
	if prev == nil { // index == 0
		l.first = next
	} else {
		prev.next = next
	}

	if next == nil { // index == size - 1
		l.last = prev
	} else {
		next.prev = prev
	}
	l.size -= 1
	return node.element
}

func (l *LinkedList) RemoveOfElement(e interface{}) {
	if index := l.IndexOfElement(e); index != ElementNotFound {
		l.Remove(index)
	}
}

func (l *LinkedList) IndexOfElement(e interface{}) int {
	node := l.first
	for i := 0; i < l.size; i++ {
		if l.Equal != nil {
			if l.Equal(node.element, e) {
				return i
			}
		} else if node.element == e {
			return i
		}
		node = node.next
	}
	return ElementNotFound
}

func (l *LinkedList) nodeForIndex(index int) *_Node {
	l.rangeCheck(index)
	if index < (l.size >> 1) {
		node := l.first
		for i := 0; i < index; i++ {
			node = node.next
		}
		return node
	} else {
		node := l.last
		for i := l.size - 1; i > index; i-- {
			node = node.prev
		}
		return node
	}
}

func (l *LinkedList) String() string {
	builder := &strings.Builder{}
	fmt.Fprintf(builder, "LinkedList: size=%v, [", l.size)
	node := l.first
	for i := 0; i < l.size; i++ {
		if i != 0 {
			builder.WriteString(", ")
		}
		if node.prev != nil {
			fmt.Fprint(builder, node.prev.element)
		} else {
			builder.WriteString("nil")
		}
		fmt.Fprintf(builder, "_%v_", node.element)
		if node.next != nil {
			fmt.Fprint(builder, node.next.element)
		} else {
			builder.WriteString("nil")
		}
		node = node.next
	}
	builder.WriteString("]")
	return builder.String()
}
