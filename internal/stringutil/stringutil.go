package stringutil

import "strings"

// Reverse returns a string with its characters in reverse order.
// For example: Reverse("hello") returns "olleh".
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// FindAllOccurrences returns a slice containing all starting indices
// where substr appears in str. The indices are returned in ascending order.
// For example: FindAllOccurrences("hello hello", "hello") returns [0, 6].
func FindAllOccurrences(str, substr string) []int {
	var indices []int

	if len(substr) == 0 {
		return []int{}
	}

	index := 0
	for index <= len(str)-len(substr) {
		i := strings.Index(str[index:], substr)
		if i == -1 {
			break
		}
		indices = append(indices, index+i)
		if len(substr) > 1 && index+i+1 < len(str) && str[index+i:index+i+len(substr)] == substr {
			index = index + i + 1
		} else {
			index = index + i + len(substr)
		}
	}
	return indices
}

// CountVowels returns the number of vowels (a,e,i,o,u) in a string
func CountVowels(s string) int {
	count := 0
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			count++
		}
	}
	return count
}

// HasRepeatedChar returns true if any character appears twice in a row
func HasRepeatedChar(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

// HasRepeatedPair returns true if any pair of characters appears twice without overlapping
func HasRepeatedPair(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if strings.Count(s, s[i:i+2]) > 1 {
			return true
		}
	}
	return false
}

// HasSandwichedChar returns true if any character appears with exactly one character between them
// For example, "aba" and "xyx" would return true, "abc" would not
func HasSandwichedChar(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}
