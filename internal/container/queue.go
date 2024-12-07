package container

import "container/list"

// Queue implements a FIFO (First In First Out) queue data structure.
// Elements are added to the back and removed from the front of the queue.
type Queue[T any] struct {
	q *list.List
}

// NewQueue creates and returns a new empty queue.
// The type parameter T can be any type.
func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		q: list.New(),
	}
}

// IsEmpty returns true if the queue contains no elements.
func (q Queue[T]) IsEmpty() bool {
	return q.q.Len() == 0
}

// Len returns the number of elements in the queue.
func (q Queue[T]) Len() int {
	return q.q.Len()
}

// Push adds a new element to the back of the queue.
func (q Queue[T]) Push(data T) {
	q.q.PushBack(data)
}

// Pop removes and returns the element at the front of the queue.
// Panics if the queue is empty.
func (q Queue[T]) Pop() T {
	if q.IsEmpty() {
		panic("Queue is empty")
	}
	data := q.q.Front()
	q.q.Remove(data)
	return data.Value.(T)
}
