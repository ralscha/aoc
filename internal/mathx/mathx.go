package mathx

import (
	"math"
)

// Combinations generates all possible non-empty combinations of elements from the input slice.
// It returns a slice containing all possible subsets of the input slice, excluding the empty set.
// The function uses a bit manipulation technique to generate combinations efficiently.
// For example, given [1,2], it returns [[1], [2], [1,2]].
func Combinations[T any](input []T) [][]T {
	var results [][]T
	for i := range 1 << uint(len(input)) {
		var combination []T
		for ix, in := range input {
			if i&(1<<uint(ix)) > 0 {
				combination = append(combination, in)
			}
		}
		if len(combination) > 0 {
			results = append(results, combination)
		}
	}
	return results
}

// Permutations generates all possible arrangements of elements from the input slice where order matters.
// It returns a slice containing all possible permutations of the input slice.
// For example, given [1,2], it returns [[1,2], [2,1]].
// Returns an empty slice if input is empty, or a slice with one element if input has only one element.
func Permutations[T any](input []T) [][]T {
	if len(input) == 0 {
		return [][]T{}
	}
	if len(input) == 1 {
		return [][]T{input}
	}
	var result [][]T
	for i, v := range input {
		remaining := make([]T, 0, len(input)-1)
		remaining = append(remaining, input[:i]...)
		remaining = append(remaining, input[i+1:]...)
		for _, p := range Permutations[T](remaining) {
			result = append(result, append([]T{v}, p...))
		}
	}
	return result
}

// CartesianProductSelf generates the cartesian product of the given values, repeated n times.
// This includes all possible arrangements with repetition, where order matters.
// Parameters:
//   - n: The number of positions to generate (the length of each resulting arrangement)
//   - values: The slice of values to use for generating combinations
//
// Returns nil if n <= 0 or if values is empty.
// For example, with n=2 and values=[1,2], it returns [[1,1], [1,2], [2,1], [2,2]].
func CartesianProductSelf[T any](n int, values []T) [][]T {
	if n == 0 {
		return [][]T{{}}
	}
	result := make([][]T, 0)
	for _, v := range values {
		for _, p := range CartesianProductSelf[T](n-1, values) {
			result = append(result, append([]T{v}, p...))
		}
	}
	return result
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// Abs returns the absolute value of the input number.
// It works with both integer and floating-point types using Go's constraints package.
// For example: Abs(-5) returns 5, Abs(3.14) returns 3.14.
func Abs[E Number](input E) E {
	if input < 0 {
		return -input
	}
	return input
}

// Round rounds a float64 to the nearest integer
func Round(x float64) int {
	if x < 0 {
		return int(x - 0.5)
	}
	return int(x + 0.5)
}

// Lcm calculates the Least Common Multiple (LCM) of a slice of integers.
// It uses the property that LCM(a,b) = (a * b) / GCD(a,b) and extends it to multiple numbers.
// The function assumes the input slice has at least one element.
// For example: Lcm([4,6]) returns 12.
func Lcm(n []int) int {
	if len(n) == 0 {
		return 0
	}
	result := n[0]
	for i := 1; i < len(n); i++ {
		result = result * n[i] / Gcd(result, n[i])
	}
	return result
}

// Gcd calculates the Greatest Common Divisor (GCD) of two integers using Euclidean algorithm.
// For example: Gcd(48,18) returns 6.
func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// MannhattanDistance calculates the Manhattan distance (also known as L1 distance or taxicab distance)
// between two points (x1,y1) and (x2,y2) in a 2D grid.
// The Manhattan distance is the sum of the absolute differences of their coordinates.
// For example: MannhattanDistance(1,1,4,5) returns 7.
func MannhattanDistance(x1, y1, x2, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
}

// Factors returns all factors of a number n
func Factors(n int) []int {
	if n == 1 {
		return []int{1}
	}
	factors := []int{1, n}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			factors = append(factors, i)
			if i != n/i {
				factors = append(factors, n/i)
			}
		}
	}
	return factors
}
