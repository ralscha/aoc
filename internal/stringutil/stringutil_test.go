package stringutil

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"12345", "54321"},
		{"Hello, 世界", "界世 ,olleH"},
	}

	for _, test := range tests {
		if got := Reverse(test.input); got != test.expected {
			t.Errorf("Reverse(%q) = %q; want %q", test.input, got, test.expected)
		}
	}
}

func TestFindAllOccurrences(t *testing.T) {
	var empty []int
	tests := []struct {
		str      string
		substr   string
		expected []int
	}{
		{"hello hello", "hello", []int{0, 6}},
		{"aaa", "aa", []int{0, 1}},
		{"test", "x", empty},
		{"", "test", empty},
		{"test", "", []int{}},
	}

	for _, test := range tests {
		got := FindAllOccurrences(test.str, test.substr)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("FindAllOccurrences(%q, %q) = %v; want %v",
				test.str, test.substr, got, test.expected)
		}
	}
}

func TestCountVowels(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello", 2},
		{"aeiou", 5},
		{"xyz", 0},
		{"", 0},
		{"AEIOU", 0}, // only counts lowercase vowels
		{"beautiful", 5},
	}

	for _, test := range tests {
		if got := CountVowels(test.input); got != test.expected {
			t.Errorf("CountVowels(%q) = %d; want %d", test.input, got, test.expected)
		}
	}
}

func TestHasRepeatedChar(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"hello", true},  // 'll'
		{"world", false}, // no repeats
		{"", false},
		{"a", false},
		{"aaa", true},
		{"test", false},
	}

	for _, test := range tests {
		if got := HasRepeatedChar(test.input); got != test.expected {
			t.Errorf("HasRepeatedChar(%q) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func TestHasRepeatedPair(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"xyxy", true},       // 'xy' appears twice
		{"aabcdefgaa", true}, // 'aa' appears twice
		{"aaa", false},       // overlapping 'aa' doesn't count
		{"", false},
		{"abcabc", true}, // 'abc' appears twice
		{"test", false},
	}

	for _, test := range tests {
		if got := HasRepeatedPair(test.input); got != test.expected {
			t.Errorf("HasRepeatedPair(%q) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func TestHasSandwichedChar(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"xyx", true},        // 'x_x'
		{"abcdefeghi", true}, // 'e_e'
		{"aaa", true},        // 'a_a'
		{"", false},
		{"ab", false},
		{"abc", false},
		{"test", false},
	}

	for _, test := range tests {
		if got := HasSandwichedChar(test.input); got != test.expected {
			t.Errorf("HasSandwichedChar(%q) = %v; want %v", test.input, got, test.expected)
		}
	}
}
