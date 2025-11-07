package slicesx

import (
	"iter"

	"aoc/internal/container"
)

// Number is a constraint that permits any numeric type that supports addition.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// Sum returns the sum of all elements in the slice.
// It works with any numeric type (integers and floats).
// Returns zero for an empty slice.
// For example: Sum([]int{1, 2, 3, 4, 5}) returns 15.
func Sum[T Number](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// SumIter returns the sum of all elements from an iterator.
// It works with any numeric type (integers and floats).
// Returns zero for an empty iterator.
// For example: SumIter(slices.Values([]int{1, 2, 3, 4, 5})) returns 15.
func SumIter[T Number](seq iter.Seq[T]) T {
	var sum T
	for v := range seq {
		sum += v
	}
	return sum
}

// Unique returns a new slice with duplicate values removed.
// The order of elements is not guaranteed to be preserved.
// It works with any comparable type.
// For example: Unique([]int{1, 2, 2, 3, 1, 4}) returns []int{1, 2, 3, 4} (in some order).
func Unique[T comparable](slice []T) []T {
	set := container.NewSet[T]()
	set.AddAll(slice)
	return set.Values()
}
