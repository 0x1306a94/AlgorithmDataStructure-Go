package queue

type CircleDeque struct {
	size, front int
	elements    []interface{}
}

func NewCircleDeque(capacity int) *CircleDeque {
	if capacity < defaultCapacity {
		capacity = defaultCapacity
	}
	return &CircleDeque{
		size:     0,
		front:    0,
		elements: make([]interface{}, capacity),
	}
}

func (q *CircleDeque) Size() int {
	return q.size
}

func (q *CircleDeque) IsEmpty() bool {
	return q.size == 0
}

func (q *CircleDeque) Clear() {
	q.front = 0
	q.size = 0
	if cap(q.elements) > defaultCapacity {
		q.elements = make([]interface{}, defaultCapacity)
	} else {
		size := len(q.elements)
		for i := 0; i < size; i++ {
			q.elements[q.index(i)] = nil
		}
	}
}

func (q *CircleDeque) EnQueueFront(value interface{}) {
	q.ensureCapacity(q.size + 1)
	q.front = q.index(-1)
	q.elements[q.front] = value
	q.size++
}

func (q *CircleDeque) EnQueueRear(value interface{}) {
	q.ensureCapacity(q.size + 1)
	q.elements[q.index(q.size)] = value
	q.size++
}

func (q *CircleDeque) DeQueueFront() interface{} {
	frontElement := q.elements[q.front]
	q.elements[q.front] = nil
	q.front = q.index(1)
	q.size--
	q.trim()
	return frontElement
}

func (q *CircleDeque) DeQueueRear() interface{} {
	rearIndex := q.index(q.size - 1)
	rear := q.elements[rearIndex]
	q.elements[rearIndex] = nil
	q.size--
	q.trim()
	return rear
}

func (q *CircleDeque) Front() interface{} {
	if q.size == 0 {
		return nil
	}
	return q.elements[q.front]
}

func (q *CircleDeque) Rear() interface{} {
	if q.size == 0 {
		return nil
	}
	return q.elements[q.index(q.size-1)]
}

func (q *CircleDeque) index(index int) int {
	index += q.front
	if index < 0 {
		return index + len(q.elements)
	}
	if index >= len(q.elements) {
		return index - len(q.elements)
	}
	return index
}

func (q *CircleDeque) ensureCapacity(capacity int) {
	oldCapacity := cap(q.elements)
	if oldCapacity >= capacity {
		return
	}
	// 新容量为旧容量的1.5倍
	newCapacity := oldCapacity + (oldCapacity >> 1)
	newElements := make([]interface{}, newCapacity)
	for i := 0; i < q.size; i++ {
		newElements[i] = q.elements[q.index(i)]
	}
	q.elements = newElements
	// 重置front
	q.front = 0
}

func (q *CircleDeque) trim() {
	oldCapacity := cap(q.elements)
	newCapacity := oldCapacity >> 1
	if q.size > newCapacity || oldCapacity <= defaultCapacity {
		return
	}
	newElements := make([]interface{}, newCapacity)
	size := q.size
	for i := 0; i < size; i++ {
		newElements[i] = q.elements[q.index(i)]
	}
	q.elements = newElements
	q.front = 0
}
