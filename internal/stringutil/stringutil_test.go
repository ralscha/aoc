package stringutil

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "reverse abc",
			args: "abc",
			want: "cba",
		},
		{
			name: "reverse empty",
			args: "",
			want: "",
		},
		{
			name: "reverse abcd",
			args: "abcd",
			want: "dcba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

// test stringutil.FindAllOccurrences
func TestFindAllOccurrences(t *testing.T) {
	tests := []struct {
		name   string
		args   string
		substr string
		want   []int
	}{
		{
			name:   "find all occurrences 1",
			args:   "abababab",
			substr: "ab",
			want:   []int{0, 2, 4, 6},
		},
		{
			name:   "find all occurrences 2",
			args:   "abababab",
			substr: "ba",
			want:   []int{1, 3, 5},
		},
		{
			name:   "find all occurrences 3",
			args:   "abababab",
			substr: "c",
			want:   []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAllOccurrences(tt.args, tt.substr); !equal(got, tt.want) {
				t.Errorf("FindAllOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b []int) bool {
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
