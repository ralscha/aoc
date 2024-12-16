package container

import (
	"container/heap"
)

// PriorityQueue implements a priority queue data structure where elements are dequeued
// based on their priority value. Lower priority values are dequeued first.
type PriorityQueue[T any] struct {
	pq *priorityQueue[T]
}

// NewPriorityQueue creates and returns a new empty priority queue.
// The type parameter T can be any type.
func NewPriorityQueue[T any]() PriorityQueue[T] {
	var pq priorityQueue[T]
	heap.Init(&pq)
	return PriorityQueue[T]{
		pq: &pq,
	}
}

// IsEmpty returns true if the priority queue contains no elements.
func (pq PriorityQueue[T]) IsEmpty() bool {
	return pq.pq.Len() == 0
}

// Len returns the number of elements in the priority queue.
func (pq PriorityQueue[T]) Len() int {
	return pq.pq.Len()
}

// Push adds a new element to the priority queue with the given priority.
// Lower priority values are dequeued first.
func (pq PriorityQueue[T]) Push(value T, priority int) {
	heap.Push(pq.pq, &item[T]{
		value:    value,
		priority: priority,
	})
}

// Pop removes and returns the element with the lowest priority value from the queue.
// Panics if the queue is empty.
func (pq PriorityQueue[T]) Pop() T {
	item := heap.Pop(pq.pq).(*item[T])
	return item.value
}

// Peek returns the element with the lowest priority value from the queue without removing it.
func (pq PriorityQueue[T]) Peek() T {
	item := (*pq.pq)[0]
	return item.value
}

type item[T any] struct {
	value    T
	priority int
}

type priorityQueue[T any] []*item[T]

func (pq priorityQueue[T]) Len() int { return len(pq) }

func (pq priorityQueue[T]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq priorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue[T]) Push(x any) {
	item := x.(*item[T])
	*pq = append(*pq, item)
}

func (pq *priorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
