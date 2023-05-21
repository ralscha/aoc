package container

import "container/list"

type Queue[T any] struct {
	q *list.List
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		q: list.New(),
	}
}

func (q Queue[T]) IsEmpty() bool {
	return q.q.Len() == 0
}

func (q Queue[T]) Len() int {
	return q.q.Len()
}

func (q Queue[T]) Push(data T) {
	q.q.PushBack(data)
}

func (q Queue[T]) Pop() T {
	if q.IsEmpty() {
		panic("Queue is empty")
	}
	data := q.q.Front()
	q.q.Remove(data)
	return data.Value
}
