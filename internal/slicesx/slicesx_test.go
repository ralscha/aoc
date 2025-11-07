package slicesx

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "single element",
			input:    []int{5},
			expected: 5,
		},
		{
			name:     "multiple positive elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "negative elements",
			input:    []int{-1, -2, -3},
			expected: -6,
		},
		{
			name:     "mixed positive and negative",
			input:    []int{10, -5, 3, -2},
			expected: 6,
		},
		{
			name:     "with zeros",
			input:    []int{0, 5, 0, 10},
			expected: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sum(tt.input)
			if result != tt.expected {
				t.Errorf("Sum(%v) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSumFloat(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{
			name:     "empty slice",
			input:    []float64{},
			expected: 0.0,
		},
		{
			name:     "single element",
			input:    []float64{5.5},
			expected: 5.5,
		},
		{
			name:     "multiple elements",
			input:    []float64{1.5, 2.5, 3.0},
			expected: 7.0,
		},
		{
			name:     "negative floats",
			input:    []float64{-1.5, -2.5},
			expected: -4.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sum(tt.input)
			if result != tt.expected {
				t.Errorf("Sum(%v) = %f, want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSumDifferentTypes(t *testing.T) {
	t.Run("uint8", func(t *testing.T) {
		input := []uint8{1, 2, 3, 4, 5}
		expected := uint8(15)
		result := Sum(input)
		if result != expected {
			t.Errorf("Sum(%v) = %d, want %d", input, result, expected)
		}
	})

	t.Run("int64", func(t *testing.T) {
		input := []int64{100, 200, 300}
		expected := int64(600)
		result := Sum(input)
		if result != expected {
			t.Errorf("Sum(%v) = %d, want %d", input, result, expected)
		}
	})

	t.Run("float32", func(t *testing.T) {
		input := []float32{1.5, 2.5, 3.0}
		expected := float32(7.0)
		result := Sum(input)
		if result != expected {
			t.Errorf("Sum(%v) = %f, want %f", input, result, expected)
		}
	})
}

func TestSumIter(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "empty iterator",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "single element",
			input:    []int{5},
			expected: 5,
		},
		{
			name:     "multiple positive elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "negative elements",
			input:    []int{-1, -2, -3},
			expected: -6,
		},
		{
			name:     "mixed positive and negative",
			input:    []int{10, -5, 3, -2},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumIter(slices.Values(tt.input))
			if result != tt.expected {
				t.Errorf("SumIter(%v) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSumIterFloat(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{
			name:     "empty iterator",
			input:    []float64{},
			expected: 0.0,
		},
		{
			name:     "multiple floats",
			input:    []float64{1.5, 2.5, 3.0},
			expected: 7.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumIter(slices.Values(tt.input))
			if result != tt.expected {
				t.Errorf("SumIter(%v) = %f, want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "with duplicates",
			input:    []int{1, 2, 2, 3, 1, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "all duplicates",
			input:    []int{5, 5, 5, 5},
			expected: []int{5},
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "multiple duplicates",
			input:    []int{1, 1, 2, 2, 3, 3, 1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Unique(tt.input)
			// Sort both slices for comparison since order is not guaranteed
			slices.Sort(result)
			slices.Sort(tt.expected)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("Unique(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestUniqueString(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "no duplicates",
			input:    []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "with duplicates",
			input:    []string{"apple", "banana", "apple", "cherry", "banana"},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name:     "all same",
			input:    []string{"test", "test", "test"},
			expected: []string{"test"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Unique(tt.input)
			// Sort both slices for comparison since order is not guaranteed
			slices.Sort(result)
			slices.Sort(tt.expected)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("Unique(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestUniqueInt64(t *testing.T) {
	input := []int64{100, 200, 100, 300, 200}
	result := Unique(input)
	expected := []int64{100, 200, 300}

	slices.Sort(result)
	slices.Sort(expected)

	if !slices.Equal(result, expected) {
		t.Errorf("Unique(%v) = %v, want %v", input, result, expected)
	}
}
