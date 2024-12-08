package container

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
// The slice is pre-allocated to the exact size needed to avoid reallocations.
func (s *Set[T]) Values() []T {
	values := make([]T, len(s.m))
	i := 0
	for k := range s.m {
		values[i] = k
		i++
	}
	return values
}

// Clear removes all elements from the set.
// This is more efficient than removing elements one by one.
func (s *Set[T]) Clear() {
	s.m = make(map[T]struct{})
}

// AddAll adds all elements from the given slice to the set.
// This is more efficient than adding elements one by one.
func (s *Set[T]) AddAll(data []T) {
	for _, v := range data {
		s.m[v] = struct{}{}
	}
}

// Union returns a new set containing all elements from both sets.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for k := range s.m {
		result.m[k] = struct{}{}
	}
	for k := range other.m {
		result.m[k] = struct{}{}
	}
	return result
}

// Intersection returns a new set containing elements present in both sets.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	// Iterate over the smaller set for efficiency
	if len(s.m) > len(other.m) {
		s, other = other, s
	}
	for k := range s.m {
		if other.Contains(k) {
			result.m[k] = struct{}{}
		}
	}
	return result
}

// Difference returns a new set containing elements present in this set but not in the other set.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for k := range s.m {
		if !other.Contains(k) {
			result.m[k] = struct{}{}
		}
	}
	return result
}
