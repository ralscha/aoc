// Package slicesutil provides utility functions for working with slices
package slicesutil

// Sum returns the sum of all elements in a slice of numbers.
// It works with any numeric type (int, float32, float64, etc.)
func Sum[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}
