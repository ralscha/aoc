package container

import (
	"maps"
	"slices"
)

// Set implements a collection of unique ordered elements.
// Each element can only appear once in the set.
type Set[T comparable] struct {
	m map[T]struct{}
}

// NewSet creates and returns a new empty set.
// The type parameter T must satisfy the comparable constraint.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

// Add inserts an element into the set if it's not already present.
// If the element already exists, the set remains unchanged.
func (s *Set[T]) Add(data T) {
	s.m[data] = struct{}{}
}

// Remove deletes an element from the set if it exists.
// If the element doesn't exist, the set remains unchanged.
func (s *Set[T]) Remove(data T) {
	delete(s.m, data)
}

// Contains checks if an element exists in the set.
// Returns true if the element is present, false otherwise.
func (s *Set[T]) Contains(data T) bool {
	_, ok := s.m[data]
	return ok
}

// Len returns the number of elements in the set.
func (s *Set[T]) Len() int {
	return len(s.m)
}

// Values returns a slice of all elements in the set.
// Uses maps.Keys and slices.Collect for efficient key collection.
func (s *Set[T]) Values() []T {
	return slices.Collect(maps.Keys(s.m))
}

// Copy returns a new set with the same elements as the original set.
func (s *Set[T]) Copy() *Set[T] {
	newSet := NewSet[T]()
	maps.Copy(newSet.m, s.m)
	return newSet
}

// Intersection returns a new set with elements that are present in both sets.
func (s *Set[T]) Intersection(otherSet *Set[T]) *Set[T] {
	newSet := NewSet[T]()
	for k := range s.m {
		if otherSet.Contains(k) {
			newSet.Add(k)
		}
	}
	return newSet
}
