package list

import (
	"fmt"
	"strings"
)

type SingleLinkedList struct {
	AbstractList
	// 头结点
	first *_Node
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{
		AbstractList: AbstractList{
			size: 0,
		},
		first: nil,
	}
}

func (s *SingleLinkedList) Clear() {
	s.size = 0
	s.first = nil
}

func (s *SingleLinkedList) Contains(e interface{}) bool {
	return s.IndexOfElement(e) != ElementNotFound
}

func (s *SingleLinkedList) Add(e interface{}) {
	s.AddForIndex(s.size, e)
}

func (s *SingleLinkedList) AddForIndex(index int, e interface{}) {
	s.rangeCheckForAdd(index)
	if index == 0 {
		s.first = newNode(e, nil, s.first)
	} else {
		prev := s.nodeForIndex(index - 1)
		prev.next = newNode(e, nil, prev.next)
	}
	s.size += 1
}

func (s *SingleLinkedList) Get(index int) interface{} {
	node := s.nodeForIndex(index)
	return node.element
}

func (s *SingleLinkedList) Set(index int, e interface{}) interface{} {
	node := s.nodeForIndex(index)
	old := node.element
	node.element = e
	return old
}

func (s *SingleLinkedList) Remove(index int) interface{} {
	s.rangeCheck(index)
	node := s.first
	if index == 0 {
		s.first = node.next
	} else {
		prev := s.nodeForIndex(index - 1)
		node = prev.next
		prev.next = node.next
	}
	s.size -= 1
	return node.element
}

func (s *SingleLinkedList) RemoveOfElement(e interface{}) {
	if index := s.IndexOfElement(e); index != ElementNotFound {
		s.Remove(index)
	}
}

func (s *SingleLinkedList) IndexOfElement(e interface{}) int {
	node := s.first
	for i := 0; i < s.size; i++ {
		if s.Equal != nil {
			if s.Equal(node.element, e) {
				return i
			}
		} else if node.element == e {
			return i
		}
		node = node.next
	}
	return ElementNotFound
}

func (s *SingleLinkedList) nodeForIndex(index int) *_Node {
	s.rangeCheck(index)
	node := s.first
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (s *SingleLinkedList) String() string {
	builder := &strings.Builder{}
	fmt.Fprintf(builder, "SingleLinkedList: size=%v, [", s.size)
	node := s.first
	for i := 0; i < s.size; i++ {
		if i != 0 {
			builder.WriteString(", ")
		}
		fmt.Fprint(builder, node.element)
		node = node.next
	}
	builder.WriteString("]")
	return builder.String()
}
