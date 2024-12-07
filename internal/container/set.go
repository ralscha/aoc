package container

import (
	"cmp"
)

// Set implements a collection of unique ordered elements.
// Each element can only appear once in the set.
type Set[T cmp.Ordered] struct {
	m map[T]struct{}
}

// NewSet creates and returns a new empty set.
// The type parameter T must satisfy the cmp.Ordered constraint.
func NewSet[T cmp.Ordered]() Set[T] {
	return Set[T]{
		m: make(map[T]struct{}),
	}
}

// Add inserts an element into the set if it's not already present.
// If the element already exists, the set remains unchanged.
func (s Set[T]) Add(data T) {
	s.m[data] = struct{}{}
}

// Remove deletes an element from the set if it exists.
// If the element doesn't exist, the set remains unchanged.
func (s Set[T]) Remove(data T) {
	delete(s.m, data)
}

// Contains checks if an element exists in the set.
// Returns true if the element is present, false otherwise.
func (s Set[T]) Contains(data T) bool {
	_, ok := s.m[data]
	return ok
}

// Len returns the number of elements in the set.
func (s Set[T]) Len() int {
	return len(s.m)
}
