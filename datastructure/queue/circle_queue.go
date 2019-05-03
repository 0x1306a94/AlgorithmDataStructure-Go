package queue

const (
	defaultCapacity int = 10
)

type CircleQueue struct {
	size, front int
	elements    []interface{}
}

func NewCircleQueue(capacity int) *CircleQueue {
	if capacity < defaultCapacity {
		capacity = defaultCapacity
	}
	return &CircleQueue{
		size:     0,
		front:    0,
		elements: make([]interface{}, capacity),
	}
}

func (q *CircleQueue) Size() int {
	return q.size
}

func (q *CircleQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *CircleQueue) Clear() {
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

func (q *CircleQueue) EnQueue(value interface{}) {
	q.ensureCapacity(q.size + 1)
	q.elements[q.index(q.size)] = value
	q.size++
}

func (q *CircleQueue) DeQueue() interface{} {
	frontElement := q.elements[q.front]
	q.elements[q.front] = nil
	q.front = q.index(1)
	q.size--
	q.trim()
	return frontElement
}

func (q *CircleQueue) Front() interface{} {
	if q.size == 0 {
		return nil
	}
	return q.elements[q.front]
}

func (q *CircleQueue) index(index int) int {
	index += q.front
	if index >= len(q.elements) {
		return index - len(q.elements)
	}
	return index
}

func (q *CircleQueue) ensureCapacity(capacity int) {
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

func (q *CircleQueue) trim() {
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
