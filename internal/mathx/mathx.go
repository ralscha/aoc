package mathx

import "golang.org/x/exp/constraints"

func Combinations[T any](input []T) [][]T {
	var results [][]T
	for i := 0; i < 1<<uint(len(input)); i++ {
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

func Abs[E constraints.Float | constraints.Integer](input E) E {
	if input < 0 {
		return -input
	}
	return input
}

func Lcm(n []int) int {
	result := n[0]
	for i := 1; i < len(n); i++ {
		result = (result * n[i]) / Gcd(result, n[i])
	}
	return result
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func MannhattanDistance(x1, y1, x2, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
}
