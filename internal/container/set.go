package container

import (
	"cmp"
)

type Set[T cmp.Ordered] struct {
	m map[T]struct{}
}

func NewSet[T cmp.Ordered]() Set[T] {
	return Set[T]{
		m: make(map[T]struct{}),
	}
}

func (s Set[T]) Add(data T) {
	s.m[data] = struct{}{}
}

func (s Set[T]) Remove(data T) {
	delete(s.m, data)
}

func (s Set[T]) Contains(data T) bool {
	_, ok := s.m[data]
	return ok
}

func (s Set[T]) Len() int {
	return len(s.m)
}
