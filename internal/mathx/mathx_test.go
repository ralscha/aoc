package mathx

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{
			name: "abs 1",
			args: 1,
			want: 1,
		},
		{
			name: "abs -1",
			args: -1,
			want: 1,
		},
		{
			name: "abs 0",
			args: 0,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "min",
			args: []int{1, 2, 3, 4, 5},
			want: 1,
		},
		{
			name: "min negative",
			args: []int{-1, -2, -3, -4, -5},
			want: -5,
		},
		{
			name: "min negative and positive",
			args: []int{-1, -2, -3, -4, -5, 1, 2, 3, 4, 5},
			want: -5,
		},
		{
			name: "min 1",
			args: []int{1},
			want: 1,
		},
		{
			name: "min 2",
			args: []int{1, 2},
			want: 1,
		},
		{
			name: "min 2 reverse",
			args: []int{2, 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args[0], tt.args[1:]...); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "max",
			args: []int{1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "max negative",
			args: []int{-1, -2, -3, -4, -5},
			want: -1,
		},
		{
			name: "max negative and positive",
			args: []int{-1, -2, -3, -4, -5, 1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "max 1",
			args: []int{1},
			want: 1,
		},
		{
			name: "max 2",
			args: []int{1, 2},
			want: 2,
		},
		{
			name: "max 2 reverse",
			args: []int{2, 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args[0], tt.args[1:]...); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCombinations(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want [][]int
	}{
		{
			name: "combinations 1",
			args: []int{1},
			want: [][]int{{1}},
		},
		{
			name: "combinations 2",
			args: []int{1, 2},
			want: [][]int{{1}, {2}, {1, 2}},
		},
		{
			name: "combinations 3",
			args: []int{1, 2, 3},
			want: [][]int{{1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name: "combinations 4",
			args: []int{1, 2, 3, 4},
			want: [][]int{{1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}, {4}, {1, 4}, {2, 4}, {1, 2, 4}, {3, 4}, {1, 3, 4}, {2, 3, 4}, {1, 2, 3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Combinations(tt.args); !equal(got, tt.want) {
				t.Errorf("Combinations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestPermutations(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want [][]int
	}{
		{
			name: "permutations 1",
			args: []int{1},
			want: [][]int{{1}},
		},
		{
			name: "permutations 2",
			args: []int{1, 2},
			want: [][]int{{1, 2}, {2, 1}},
		},
		{
			name: "permutations 3",
			args: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name: "permutations 4",
			args: []int{1, 2, 3, 4},
			want: [][]int{{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4}, {1, 3, 4, 2}, {1, 4, 2, 3}, {1, 4, 3, 2}, {2, 1, 3, 4}, {2, 1, 4, 3}, {2, 3, 1, 4}, {2, 3, 4, 1}, {2, 4, 1, 3}, {2, 4, 3, 1}, {3, 1, 2, 4}, {3, 1, 4, 2}, {3, 2, 1, 4}, {3, 2, 4, 1}, {3, 4, 1, 2}, {3, 4, 2, 1}, {4, 1, 2, 3}, {4, 1, 3, 2}, {4, 2, 1, 3}, {4, 2, 3, 1}, {4, 3, 1, 2}, {4, 3, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Permutations(tt.args); !equal(got, tt.want) {
				t.Errorf("Permutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
