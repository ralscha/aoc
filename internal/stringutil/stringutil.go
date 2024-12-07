package stringutil

import "strings"

// Reverse returns a string with its characters in reverse order.
// For example: Reverse("hello") returns "olleh".
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// FindAllOccurrences returns a slice containing all starting indices
// where substr appears in str. The indices are returned in ascending order.
// For example: FindAllOccurrences("hello hello", "hello") returns [0, 6].
func FindAllOccurrences(str, substr string) []int {
	var indices []int
	index := strings.Index(str, substr)
	for index != -1 {
		indices = append(indices, index)
		index = strings.Index(str[index+1:], substr)
		if index != -1 {
			index += indices[len(indices)-1] + 1
		}
	}
	return indices
}
