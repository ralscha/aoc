package container

import (
	"slices"
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewSet[int]()
	if s == nil {
		t.Error("NewSet returned nil")
	}
	if s.Len() != 0 {
		t.Errorf("Expected empty set, got length %d", s.Len())
	}

	// Test with string type
	ss := NewSet[string]()
	if ss == nil {
		t.Error("NewSet[string] returned nil")
	}
	if ss.Len() != 0 {
		t.Errorf("Expected empty string set, got length %d", ss.Len())
	}
}

func TestSetBasicOperations(t *testing.T) {
	s := NewSet[int]()

	// Test Add and Contains
	s.Add(1)
	if !s.Contains(1) {
		t.Error("Set should contain 1")
	}
	if s.Contains(2) {
		t.Error("Set should not contain 2")
	}

	// Test duplicate adds
	s.Add(1)
	if s.Len() != 1 {
		t.Error("Adding duplicate should not increase size")
	}

	// Test Remove
	s.Remove(1)
	if s.Contains(1) {
		t.Error("Set should not contain 1 after removal")
	}

	// Test removing non-existent element
	s.Remove(2)
	if s.Len() != 0 {
		t.Error("Removing non-existent element should not affect size")
	}

	// Test with string type
	ss := NewSet[string]()
	ss.Add("hello")
	ss.Add("world")
	if !ss.Contains("hello") {
		t.Error("Set should contain 'hello'")
	}
	if !ss.Contains("world") {
		t.Error("Set should contain 'world'")
	}
	if ss.Contains("foo") {
		t.Error("Set should not contain 'foo'")
	}
	ss.Remove("hello")
	if ss.Contains("hello") {
		t.Error("Set should not contain 'hello' after removal")
	}
}

func TestSetValues(t *testing.T) {
	s := NewSet[int]()
	numbers := []int{1, 2, 3, 4, 5}

	for _, n := range numbers {
		s.Add(n)
	}

	values := s.Values()
	if len(values) != len(numbers) {
		t.Errorf("Expected %d values, got %d", len(numbers), len(values))
	}

	// Sort both slices for comparison since map iteration order is random
	slices.Sort(values)
	if !equalSlices(values, numbers) {
		t.Errorf("Expected values %v, got %v", numbers, values)
	}

	// Test with string type
	ss := NewSet[string]()
	strings := []string{"apple", "banana", "cherry"}
	for _, s := range strings {
		ss.Add(s)
	}

	strValues := ss.Values()
	if len(strValues) != len(strings) {
		t.Errorf("Expected %d values, got %d", len(strings), len(strValues))
	}

	slices.Sort(strValues)
	slices.Sort(strings)
	if !equalStringSlices(strValues, strings) {
		t.Errorf("Expected values %v, got %v", strings, strValues)
	}
}

// Helper function to compare two sorted slices
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Helper function to compare two sorted string slices
func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
