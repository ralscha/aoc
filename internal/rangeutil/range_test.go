package rangeutil

import (
	"testing"
)

func TestRange(t *testing.T) {
	r := NewRange(1, 5)

	// Test Length
	if r.Length() != 5 {
		t.Errorf("Expected length 5, got %d", r.Length())
	}

	// Test Contains
	tests := []struct {
		value    int
		expected bool
	}{
		{0, false},
		{1, true},
		{3, true},
		{5, true},
		{6, false},
	}

	for _, test := range tests {
		if got := r.Contains(test.value); got != test.expected {
			t.Errorf("Contains(%d) = %v, want %v", test.value, got, test.expected)
		}
	}
}

func TestRangeOverlaps(t *testing.T) {
	tests := []struct {
		r1       Range
		r2       Range
		expected bool
	}{
		{NewRange(1, 5), NewRange(4, 8), true},
		{NewRange(1, 5), NewRange(5, 8), true},
		{NewRange(1, 5), NewRange(6, 8), false},
		{NewRange(1, 5), NewRange(0, 2), true},
		{NewRange(1, 5), NewRange(0, 0), false},
	}

	for _, test := range tests {
		if got := test.r1.Overlaps(test.r2); got != test.expected {
			t.Errorf("Range%v.Overlaps(%v) = %v, want %v",
				test.r1, test.r2, got, test.expected)
		}
	}
}

func TestRangeSplit(t *testing.T) {
	r := NewRange(1, 5)

	tests := []struct {
		at          int
		wantBefore  *Range
		wantAfter   *Range
		description string
	}{
		{0, nil, &Range{1, 5}, "split before start"},
		{1, nil, &Range{1, 5}, "split at start"},
		{3, &Range{1, 2}, &Range{3, 5}, "split in middle"},
		{6, &Range{1, 5}, nil, "split after end"},
	}

	for _, test := range tests {
		before, after := r.Split(test.at)
		if !rangeEqual(before, test.wantBefore) || !rangeEqual(after, test.wantAfter) {
			t.Errorf("%s: Split(%d) = {%v, %v}, want {%v, %v}",
				test.description, test.at, before, after, test.wantBefore, test.wantAfter)
		}
	}
}

func TestRangeIntersection(t *testing.T) {
	tests := []struct {
		r1   Range
		r2   Range
		want *Range
	}{
		{NewRange(1, 5), NewRange(4, 8), &Range{4, 5}},
		{NewRange(1, 5), NewRange(5, 8), &Range{5, 5}},
		{NewRange(1, 5), NewRange(6, 8), nil},
		{NewRange(1, 5), NewRange(0, 2), &Range{1, 2}},
	}

	for _, test := range tests {
		if got := test.r1.Intersection(test.r2); !rangeEqual(got, test.want) {
			t.Errorf("Range%v.Intersection(%v) = %v, want %v",
				test.r1, test.r2, got, test.want)
		}
	}
}

func TestRangeUnion(t *testing.T) {
	tests := []struct {
		r1   Range
		r2   Range
		want *Range
	}{
		{NewRange(1, 5), NewRange(4, 8), &Range{1, 8}},
		{NewRange(1, 5), NewRange(5, 8), &Range{1, 8}},
		{NewRange(1, 5), NewRange(6, 8), &Range{1, 8}},
		{NewRange(1, 5), NewRange(7, 8), nil},
		{NewRange(1, 5), NewRange(0, 2), &Range{0, 5}},
	}

	for _, test := range tests {
		if got := test.r1.Union(test.r2); !rangeEqual(got, test.want) {
			t.Errorf("Range%v.Union(%v) = %v, want %v",
				test.r1, test.r2, got, test.want)
		}
	}
}

func TestRangeMap(t *testing.T) {
	r := NewRange(1, 5)
	double := func(x int) int { return x * 2 }
	got := r.Map(double)
	want := Range{2, 10}

	if got != want {
		t.Errorf("Range%v.Map(double) = %v, want %v", r, got, want)
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		ranges []Range
		want   []Range
	}{
		{
			ranges: []Range{{1, 3}, {2, 4}, {6, 8}},
			want:   []Range{{1, 4}, {6, 8}},
		},
		{
			ranges: []Range{{1, 3}, {4, 6}, {5, 8}},
			want:   []Range{{1, 8}},
		},
		{
			ranges: []Range{{1, 5}, {2, 3}, {4, 6}},
			want:   []Range{{1, 6}},
		},
	}

	for _, test := range tests {
		got := Merge(test.ranges)
		if !rangeSliceEqual(got, test.want) {
			t.Errorf("Merge(%v) = %v, want %v", test.ranges, got, test.want)
		}
	}
}

// Helper functions for comparing Range values
func rangeEqual(a, b *Range) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Start == b.Start && a.End == b.End
}

func rangeSliceEqual(a, b []Range) bool {
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
