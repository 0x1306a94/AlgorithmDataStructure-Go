package queue

import (
	"fmt"
	"testing"
)

func TestNewQueue(t *testing.T) {

	q := NewQueue()
	q.EnQueue(11)
	q.EnQueue(22)
	q.EnQueue(33)
	q.EnQueue(44)

	// 11 22 33 44
	for !q.IsEmpty() {
		fmt.Println(q.DeQueue())
	}
}

func TestNewDeque(t *testing.T) {

	q := NewDeque()
	q.EnQueueFront(11)
	q.EnQueueFront(22)
	q.EnQueueRear(33)
	q.EnQueueRear(44)
	// 尾 44 33 11 22 头
	for !q.IsEmpty() {
		fmt.Println(q.DeQueueRear())
	}
}

func TestNewCircleQueue(t *testing.T) {

	q := NewCircleQueue(10)
	fmt.Println("capacity: ", cap(q.elements))
	for i := 0; i < 20; i++ {
		q.EnQueue(i + 1)
	}
	for i := 0; i < 4; i++ {
		q.DeQueue()
	}

	for i := 21; i <= 24; i++ {
		q.EnQueue(i)
	}
	// front:  4
	// capacity: 22
	// [23 24 <nil> <nil> 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22]
	fmt.Println("front: ", q.front)
	fmt.Println("capacity: ", cap(q.elements))
	fmt.Println(q.elements)

	for i := 0; i < 15; i++ {
		q.DeQueue()
	}

	// front:  0
	// capacity: 5
	// [20 21 22 23 24]
	fmt.Println("front: ", q.front)
	fmt.Println("capacity: ", cap(q.elements))
	fmt.Println(q.elements)

	q.Clear()
	// front:  0
	// capacity: 5
	// [<nil> <nil> <nil> <nil> <nil>]
	fmt.Println("front: ", q.front)
	fmt.Println("capacity: ", cap(q.elements))
	fmt.Println(q.elements)
}

func TestNewCircleDeque(t *testing.T) {

	q := NewCircleDeque(10)
	fmt.Println("capacity: ", cap(q.elements))
	for i := 0; i < 20; i++ {
		q.EnQueueFront(i + 1)
	}
	// front:  17
	// capacity: 22
	// [15 14 13 12 11 10 9 8 7 6 5 4 3 2 1 <nil> <nil> 20 19 18 17 16]
	fmt.Println("front: ", q.front)
	fmt.Println("capacity: ", cap(q.elements))
	fmt.Println(q.elements)
	q.EnQueueRear(21)
	// front:  17
	// capacity: 22
	// [15 14 13 12 11 10 9 8 7 6 5 4 3 2 1 21 <nil> 20 19 18 17 16]
	fmt.Println("front: ", q.front)
	fmt.Println("capacity: ", cap(q.elements))
	fmt.Println(q.elements)

	for i := 0; i < 4; i++ {
		q.DeQueueRear()
	}
	// front:  17
	// capacity: 22
	// [15 14 13 12 11 10 9 8 7 6 5 4 <nil> <nil> <nil> <nil> <nil> 20 19 18 17 16]
	fmt.Println("front: ", q.front)
	fmt.Println("capacity: ", cap(q.elements))
	fmt.Println(q.elements)

	for i := 21; i <= 24; i++ {
		q.EnQueueRear(i)
	}
	// front:  17
	// capacity: 22
	// [15 14 13 12 11 10 9 8 7 6 5 4 21 22 23 24 <nil> 20 19 18 17 16]
	fmt.Println("front: ", q.front)
	fmt.Println("capacity: ", cap(q.elements))
	fmt.Println(q.elements)

	for i := 0; i < 15; i++ {
		q.DeQueueFront()
	}

	// front:  5
	// capacity: 11
	// [<nil> <nil> <nil> <nil> <nil> 5 4 21 22 23 24]
	fmt.Println("front: ", q.front)
	fmt.Println("capacity: ", cap(q.elements))
	fmt.Println(q.elements)

	q.Clear()
	// front:  0
	// capacity: 10
	// [<nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>]
	fmt.Println("front: ", q.front)
	fmt.Println("capacity: ", cap(q.elements))
	fmt.Println(q.elements)
}
