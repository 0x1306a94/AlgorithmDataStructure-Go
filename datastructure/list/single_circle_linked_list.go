package list

type SingleCircleLinkedList struct {
	SingleLinkedList
}

func NewSingleCircleLinkedList() *SingleCircleLinkedList {
	return &SingleCircleLinkedList{
		SingleLinkedList{
			AbstractList{
				size: 0,
			},
			nil,
		},
	}
}

func (s *SingleCircleLinkedList) Clear() {
	s.size = 0
	if s.first != nil {
		s.first.next = nil
	}
	s.first = nil
}

func (s *SingleCircleLinkedList) Add(e interface{}) {
	s.AddForIndex(s.size, e)
}

func (s *SingleCircleLinkedList) AddForIndex(index int, e interface{}) {
	s.rangeCheckForAdd(index)
	if index == 0 {
		newFirst := newNode(e, nil, s.first)
		var last *_Node
		if s.size == 0 {
			last = newFirst
		} else {
			last = s.nodeForIndex(s.size - 1)
		}
		last.next = newFirst
		s.first = newFirst
	} else {
		prev := s.nodeForIndex(index - 1)
		prev.next = newNode(e, nil, prev.next)
	}
	s.size += 1
}

func (s *SingleCircleLinkedList) Remove(index int) interface{} {
	s.rangeCheck(index)
	node := s.first
	if index == 0 {
		if s.size == 1 {
			s.first = nil
		} else {
			last := s.nodeForIndex(s.size - 1)
			s.first = s.first.next
			last.next = s.first
		}
	} else {
		prev := s.nodeForIndex(index - 1)
		node = prev.next
		prev.next = node.next
	}
	s.size -= 1
	return node.element
}
