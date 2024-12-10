package rangeutil

import "slices"

// Range represents a range of integers from start to end (inclusive)
type Range struct {
	Start int
	End   int
}

// NewRange creates a new Range with the given start and end values
func NewRange(start, end int) Range {
	return Range{Start: start, End: end}
}

// Length returns the length of the range
func (r Range) Length() int {
	return r.End - r.Start + 1
}

// Contains returns true if the given value is within the range
func (r Range) Contains(value int) bool {
	return value >= r.Start && value <= r.End
}

// Overlaps returns true if this range overlaps with another range
func (r Range) Overlaps(other Range) bool {
	return r.Start <= other.End && other.Start <= r.End
}

// Split splits this range at the given point, returning the ranges before and after
// If the split point is outside the range, returns nil for the appropriate part
func (r Range) Split(at int) (before, after *Range) {
	if at <= r.Start {
		return nil, &Range{Start: r.Start, End: r.End}
	}
	if at > r.End {
		return &Range{Start: r.Start, End: r.End}, nil
	}
	return &Range{Start: r.Start, End: at - 1}, &Range{Start: at, End: r.End}
}

// Intersection returns the intersection of this range with another range
// Returns nil if there is no intersection
func (r Range) Intersection(other Range) *Range {
	if !r.Overlaps(other) {
		return nil
	}
	start := max(r.Start, other.Start)
	end := min(r.End, other.End)
	return &Range{Start: start, End: end}
}

// Union returns the union of this range with another range if they overlap or are adjacent
// Returns nil if the ranges are not overlapping or adjacent
func (r Range) Union(other Range) *Range {
	if !r.Overlaps(other) && !r.Adjacent(other) {
		return nil
	}
	return &Range{
		Start: min(r.Start, other.Start),
		End:   max(r.End, other.End),
	}
}

// Adjacent returns true if this range is immediately adjacent to another range
func (r Range) Adjacent(other Range) bool {
	return r.Start == other.End+1 || other.Start == r.End+1
}

// Map applies a transformation function to the range and returns a new range
func (r Range) Map(transform func(int) int) Range {
	return Range{
		Start: transform(r.Start),
		End:   transform(r.End),
	}
}

// Merge merges a slice of ranges into the minimal set of non-overlapping ranges
func Merge(ranges []Range) []Range {
	if len(ranges) == 0 {
		return nil
	}

	// Sort ranges by start value
	sorted := make([]Range, len(ranges))
	copy(sorted, ranges)
	slices.SortFunc(sorted, func(a, b Range) int {
		return a.Start - b.Start
	})

	result := make([]Range, 0)
	current := sorted[0]

	for i := 1; i < len(sorted); i++ {
		if union := current.Union(sorted[i]); union != nil {
			current = *union
		} else {
			result = append(result, current)
			current = sorted[i]
		}
	}
	result = append(result, current)

	return result
}
