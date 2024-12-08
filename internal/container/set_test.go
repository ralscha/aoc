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

func TestSetClear(t *testing.T) {
	s := NewSet[int]()
	numbers := []int{1, 2, 3, 4, 5}

	for _, n := range numbers {
		s.Add(n)
	}

	s.Clear()
	if s.Len() != 0 {
		t.Error("Set should be empty after clear")
	}
	for _, n := range numbers {
		if s.Contains(n) {
			t.Errorf("Set should not contain %d after clear", n)
		}
	}

	// Test clearing an empty set
	s.Clear()
	if s.Len() != 0 {
		t.Error("Clearing empty set should keep it empty")
	}
}

func TestSetAddAll(t *testing.T) {
	s := NewSet[int]()
	numbers := []int{1, 2, 3, 4, 5}

	s.AddAll(numbers)
	if s.Len() != len(numbers) {
		t.Errorf("Expected length %d, got %d", len(numbers), s.Len())
	}

	// Test adding duplicates
	s.AddAll(numbers)
	if s.Len() != len(numbers) {
		t.Error("Adding duplicates should not increase size")
	}

	// Test adding empty slice
	s.Clear()
	s.AddAll([]int{})
	if s.Len() != 0 {
		t.Error("Adding empty slice should result in empty set")
	}

	// Test with larger dataset
	largeNumbers := make([]int, 1000)
	for i := range largeNumbers {
		largeNumbers[i] = i
	}
	s.AddAll(largeNumbers)
	if s.Len() != len(largeNumbers) {
		t.Errorf("Expected length %d, got %d", len(largeNumbers), s.Len())
	}
}

func TestSetUnion(t *testing.T) {
	tests := []struct {
		name     string
		set1     []int
		set2     []int
		expected []int
	}{
		{
			name:     "Disjoint sets",
			set1:     []int{1, 2, 3},
			set2:     []int{4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Overlapping sets",
			set1:     []int{1, 2, 3},
			set2:     []int{2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Empty first set",
			set1:     []int{},
			set2:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Empty second set",
			set1:     []int{1, 2, 3},
			set2:     []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Both empty sets",
			set1:     []int{},
			set2:     []int{},
			expected: []int{},
		},
		{
			name:     "Large overlapping sets",
			set1:     generateSequence(0, 1000),
			set2:     generateSequence(500, 1500),
			expected: generateSequence(0, 1500),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s1 := NewSet[int]()
			s2 := NewSet[int]()
			s1.AddAll(tt.set1)
			s2.AddAll(tt.set2)

			result := s1.Union(s2)
			values := result.Values()
			slices.Sort(values)
			slices.Sort(tt.expected)

			if !equalSlices(values, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, values)
			}
		})
	}

	// Test with string type
	t.Run("String type union", func(t *testing.T) {
		s1 := NewSet[string]()
		s2 := NewSet[string]()
		s1.AddAll([]string{"a", "b", "c"})
		s2.AddAll([]string{"b", "c", "d"})

		result := s1.Union(s2)
		values := result.Values()
		slices.Sort(values)
		expected := []string{"a", "b", "c", "d"}

		if !equalStringSlices(values, expected) {
			t.Errorf("Expected %v, got %v", expected, values)
		}
	})
}

func TestSetIntersection(t *testing.T) {
	tests := []struct {
		name     string
		set1     []int
		set2     []int
		expected []int
	}{
		{
			name:     "Disjoint sets",
			set1:     []int{1, 2, 3},
			set2:     []int{4, 5, 6},
			expected: []int{},
		},
		{
			name:     "Overlapping sets",
			set1:     []int{1, 2, 3},
			set2:     []int{2, 3, 4},
			expected: []int{2, 3},
		},
		{
			name:     "Empty first set",
			set1:     []int{},
			set2:     []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "Empty second set",
			set1:     []int{1, 2, 3},
			set2:     []int{},
			expected: []int{},
		},
		{
			name:     "Both empty sets",
			set1:     []int{},
			set2:     []int{},
			expected: []int{},
		},
		{
			name:     "Large overlapping sets",
			set1:     generateSequence(0, 1000),
			set2:     generateSequence(500, 1500),
			expected: generateSequence(500, 1000),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s1 := NewSet[int]()
			s2 := NewSet[int]()
			s1.AddAll(tt.set1)
			s2.AddAll(tt.set2)

			result := s1.Intersection(s2)
			values := result.Values()
			slices.Sort(values)
			slices.Sort(tt.expected)

			if !equalSlices(values, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, values)
			}
		})
	}

	// Test with string type
	t.Run("String type intersection", func(t *testing.T) {
		s1 := NewSet[string]()
		s2 := NewSet[string]()
		s1.AddAll([]string{"a", "b", "c"})
		s2.AddAll([]string{"b", "c", "d"})

		result := s1.Intersection(s2)
		values := result.Values()
		slices.Sort(values)
		expected := []string{"b", "c"}

		if !equalStringSlices(values, expected) {
			t.Errorf("Expected %v, got %v", expected, values)
		}
	})
}

func TestSetDifference(t *testing.T) {
	tests := []struct {
		name     string
		set1     []int
		set2     []int
		expected []int
	}{
		{
			name:     "Disjoint sets",
			set1:     []int{1, 2, 3},
			set2:     []int{4, 5, 6},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Overlapping sets",
			set1:     []int{1, 2, 3},
			set2:     []int{2, 3, 4},
			expected: []int{1},
		},
		{
			name:     "Empty first set",
			set1:     []int{},
			set2:     []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "Empty second set",
			set1:     []int{1, 2, 3},
			set2:     []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Both empty sets",
			set1:     []int{},
			set2:     []int{},
			expected: []int{},
		},
		{
			name:     "Subset",
			set1:     []int{1, 2, 3, 4},
			set2:     []int{2, 4},
			expected: []int{1, 3},
		},
		{
			name:     "Large sets difference",
			set1:     generateSequence(0, 1000),
			set2:     generateSequence(500, 1500),
			expected: generateSequence(0, 500),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s1 := NewSet[int]()
			s2 := NewSet[int]()
			s1.AddAll(tt.set1)
			s2.AddAll(tt.set2)

			result := s1.Difference(s2)
			values := result.Values()
			slices.Sort(values)
			slices.Sort(tt.expected)

			if !equalSlices(values, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, values)
			}
		})
	}

	// Test with string type
	t.Run("String type difference", func(t *testing.T) {
		s1 := NewSet[string]()
		s2 := NewSet[string]()
		s1.AddAll([]string{"a", "b", "c"})
		s2.AddAll([]string{"b", "c", "d"})

		result := s1.Difference(s2)
		values := result.Values()
		slices.Sort(values)
		expected := []string{"a"}

		if !equalStringSlices(values, expected) {
			t.Errorf("Expected %v, got %v", expected, values)
		}
	})
}

func TestChainedOperations(t *testing.T) {
	s1 := NewSet[int]()
	s2 := NewSet[int]()
	s3 := NewSet[int]()

	s1.AddAll([]int{1, 2, 3, 4})
	s2.AddAll([]int{3, 4, 5, 6})
	s3.AddAll([]int{5, 6, 7, 8})

	// Test chaining Union and Intersection
	result := s1.Union(s2).Intersection(s3)
	values := result.Values()
	slices.Sort(values)
	expected := []int{5, 6}
	if !equalSlices(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
	}

	// Test chaining Union and Difference
	result = s1.Union(s2).Difference(s3)
	values = result.Values()
	slices.Sort(values)
	expected = []int{1, 2, 3, 4}
	if !equalSlices(values, expected) {
		t.Errorf("Expected %v, got %v", expected, values)
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

// Helper function to generate a sequence of integers
func generateSequence(start, end int) []int {
	result := make([]int, end-start)
	for i := range result {
		result[i] = start + i
	}
	return result
}
