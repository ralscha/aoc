package container

import (
	"cmp"
)

// Bag represents a multiset data structure that can store multiple occurrences of ordered elements.
// It maintains a count for each unique element, allowing duplicates to be tracked efficiently.
type Bag[T cmp.Ordered] struct {
	m map[T]int
}

// NewBag creates and returns a new empty Bag instance.
// The type parameter T must satisfy the cmp.Ordered constraint.
func NewBag[T cmp.Ordered]() Bag[T] {
	return Bag[T]{
		m: make(map[T]int),
	}
}

// Add inserts an element into the bag, incrementing its count if it already exists.
func (b Bag[T]) Add(data T) {
	b.m[data]++
}

// Remove decrements the count of an element in the bag.
// If the count reaches zero, the element is removed from the bag entirely.
func (b Bag[T]) Remove(data T) {
	if b.m[data] == 1 {
		delete(b.m, data)
	} else {
		b.m[data]--
	}
}

// Contains checks if an element exists in the bag.
// Returns true if the element is present, false otherwise.
func (b Bag[T]) Contains(data T) bool {
	_, ok := b.m[data]
	return ok
}

// Count returns the number of occurrences of a given element in the bag.
// Returns 0 if the element is not present.
func (b Bag[T]) Count(data T) int {
	return b.m[data]
}

// Values returns a copy of the internal map representing the bag's contents.
// The map keys are the elements and the values are their counts.
func (b Bag[T]) Values() map[T]int {
	m := make(map[T]int)
	for k, v := range b.m {
		m[k] = v
	}
	return m
}
