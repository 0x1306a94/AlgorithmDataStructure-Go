package queue

import "github.com/0x1306a94/AlgorithmDataStructure-Go/datastructure/list"

type Queue struct {
	list *list.LinkedList
}

func NewQueue() *Queue {
	return &Queue{
		list: list.NewLinkedList(),
	}
}

func (q *Queue) Size() int {
	return q.list.Size()
}

func (q *Queue) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *Queue) Clear() {
	q.list.Clear()
}

func (q *Queue) EnQueue(value interface{}) {
	q.list.Add(value)
}

func (q *Queue) DeQueue() interface{} {
	return q.list.Remove(0)
}

func (q *Queue) Front() interface{} {
	if q.list.Size() == 0 {
		return nil
	}
	return q.list.Get(0)
}
