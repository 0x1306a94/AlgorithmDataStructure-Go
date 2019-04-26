package list

type CircleLinkedList struct {
	LinkedList

	current *_Node
}

func NewCircleLinkedList() *CircleLinkedList {
	return &CircleLinkedList{
		LinkedList: LinkedList{
			SingleLinkedList: SingleLinkedList{
				_AbstractList: _AbstractList{
					size: 0,
				},
				first: nil,
			},
			last: nil,
		},
		current: nil,
	}
}

func (l *CircleLinkedList) Clear() {
	l.size = 0
	l.first = nil
	l.last = nil
	l.current = nil
}

func (l *CircleLinkedList) Rest() {
	l.current = l.first
}

func (l *CircleLinkedList) Next() interface{} {
	if l.current == nil {
		return nil
	}
	next := l.current.next
	//old := l.current.element
	l.current = next
	return l.current.element
}

func (l *CircleLinkedList) Add(e interface{}) {
	l.AddForIndex(l.size, e)
}

func (l *CircleLinkedList) AddForIndex(index int, e interface{}) {
	l.rangeCheckForAdd(index)
	// size == 0
	// index == 0
	if l.size == index {
		oldLast := l.last
		l.last = newNode(e, oldLast, l.first)
		if oldLast == nil { // 添加第一个元素
			l.first = l.last
			l.first.next = l.first
			l.first.prev = l.first
		} else {
			oldLast.next = l.last
			l.first.prev = l.last
		}
	} else {
		next := l.nodeForIndex(index)
		prev := next.prev
		node := newNode(e, prev, next)
		next.prev = node
		prev.next = node

		if next == l.first { // index == 0
			l.first = node
		}
	}
	l.size += 1
}

func (l *CircleLinkedList) RemoveCurrent() interface{} {
	if l.current == nil {
		return nil
	}
	next := l.current.next
	oldElement := l.removeOfNode(l.current)
	if l.size == 0 {
		l.current = nil
	} else {
		l.current = next
	}
	return oldElement
}

func (l *CircleLinkedList) Remove(index int) interface{} {
	l.rangeCheck(index)
	return l.removeOfNode(l.nodeForIndex(index))
}

func (l *CircleLinkedList) removeOfNode(node *_Node) interface{} {
	if l.size == 1 {
		l.first = nil
		l.last = nil
		l.current = nil
	} else {
		prev := node.prev
		next := node.next
		prev.next = next
		next.prev = prev

		if node == l.first { // index == 0
			l.first = next
		}

		if node == l.last { // index == size - 1
			l.last = prev
		}
	}
	l.size -= 1
	return node.element
}
