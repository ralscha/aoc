package container

import (
	"container/heap"
)

type PriorityQueue[T any] struct {
	pq *priorityQueue[T]
}

func NewPriorityQueue[T any]() PriorityQueue[T] {
	var pq priorityQueue[T]
	heap.Init(&pq)
	return PriorityQueue[T]{
		pq: &pq,
	}
}

func (q PriorityQueue[T]) IsEmpty() bool {
	return q.pq.Len() == 0
}

func (pq PriorityQueue[T]) Len() int {
	return pq.pq.Len()
}

func (pq PriorityQueue[T]) Push(value T, priority int) {
	heap.Push(pq.pq, &item[T]{
		value:    value,
		priority: priority,
	})
}

func (pq PriorityQueue[T]) Pop() T {
	item := heap.Pop(pq.pq).(*item[T])
	return item.value
}

type item[T any] struct {
	value    T
	priority int
}

type priorityQueue[T any] []*item[T]

func (pq priorityQueue[T]) Len() int { return len(pq) }

func (pq priorityQueue[T]) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
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
