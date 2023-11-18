package container

import (
	"cmp"
)

type Bag[T cmp.Ordered] struct {
	m map[T]int
}

func NewBag[T cmp.Ordered]() Bag[T] {
	return Bag[T]{
		m: make(map[T]int),
	}
}

func (b Bag[T]) Add(data T) {
	b.m[data]++
}

func (b Bag[T]) Remove(data T) {
	if b.m[data] == 1 {
		delete(b.m, data)
	} else {
		b.m[data]--
	}
}

func (b Bag[T]) Contains(data T) bool {
	_, ok := b.m[data]
	return ok
}

func (b Bag[T]) Count(data T) int {
	return b.m[data]
}

func (b Bag[T]) Values() map[T]int {
	m := make(map[T]int)
	for k, v := range b.m {
		m[k] = v
	}
	return m
}
