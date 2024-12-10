package slicesutil

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected any
	}{
		{
			name:     "integers",
			input:    []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "empty integers",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "negative integers",
			input:    []int{-1, -2, 3, -4, 5},
			expected: 1,
		},
		{
			name:     "float64",
			input:    []float64{1.5, 2.5, 3.5},
			expected: 7.5,
		},
		{
			name:     "uint8",
			input:    []uint8{1, 2, 3},
			expected: uint8(6),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch input := tt.input.(type) {
			case []int:
				if got := Sum(input); got != tt.expected.(int) {
					t.Errorf("Sum() = %v, want %v", got, tt.expected)
				}
			case []float64:
				if got := Sum(input); got != tt.expected.(float64) {
					t.Errorf("Sum() = %v, want %v", got, tt.expected)
				}
			case []uint8:
				if got := Sum(input); got != tt.expected.(uint8) {
					t.Errorf("Sum() = %v, want %v", got, tt.expected)
				}
			}
		})
	}
}
