package queue

import "github.com/0x1306a94/AlgorithmDataStructure-Go/datastructure/list"

type Deque struct {
	list *list.LinkedList
}

func NewDeque() *Deque {
	return &Deque{
		list: list.NewLinkedList(),
	}
}

func (q *Deque) Size() int {
	return q.list.Size()
}

func (q *Deque) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *Deque) Clear() {
	q.list.Clear()
}

func (q *Deque) EnQueueFront(value interface{}) {
	q.list.AddForIndex(0, value)
}

func (q *Deque) EnQueueRear(value interface{}) {
	q.list.Add(value)
}

func (q *Deque) DeQueueFront() interface{} {
	return q.list.Remove(0)
}

func (q *Deque) DeQueueRear() interface{} {
	return q.list.Remove(q.list.Size() - 1)
}

func (q *Deque) Front() interface{} {
	if q.list.Size() == 0 {
		return nil
	}
	return q.list.Get(0)
}

func (q *Deque) Rear() interface{} {
	if q.list.Size() == 0 {
		return nil
	}
	return q.list.Get(q.list.Size() - 1)
}
