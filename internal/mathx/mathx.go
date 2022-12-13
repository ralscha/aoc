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
		for _, p := range Permutations(remaining) {
			result = append(result, append([]T{v}, p...))
		}
	}
	return result
}

func Min[E constraints.Ordered](inputOne E, rest ...E) E {
	min := inputOne
	for i := 0; i < len(rest); i++ {
		if rest[i] < min {
			min = rest[i]
		}
	}
	return min
}

func Max[E constraints.Ordered](inputOne E, rest ...E) E {
	max := inputOne
	for i := 0; i < len(rest); i++ {
		if rest[i] > max {
			max = rest[i]
		}
	}
	return max
}
