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

func TestCartesianProductSelf(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		values []int
		want   [][]int
	}{
		{
			name:   "empty input",
			n:      1,
			values: []int{},
			want:   nil,
		},
		{
			name:   "invalid n",
			n:      0,
			values: []int{1, 2},
			want:   nil,
		},
		{
			name:   "single value n=1",
			n:      1,
			values: []int{1},
			want:   [][]int{{1}},
		},
		{
			name:   "multiple values n=1",
			n:      1,
			values: []int{1, 2},
			want:   [][]int{{1}, {2}},
		},
		{
			name:   "multiple values n=2",
			n:      2,
			values: []int{1, 2},
			want:   [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}},
		},
		{
			name:   "multiple values n=3",
			n:      3,
			values: []int{1, 2},
			want: [][]int{
				{1, 1, 1}, {1, 1, 2}, {1, 2, 1}, {1, 2, 2},
				{2, 1, 1}, {2, 1, 2}, {2, 2, 1}, {2, 2, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CartesianProductSelf(tt.n, tt.values); !equal(got, tt.want) {
				t.Errorf("CartesianProductSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLcm(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "lcm 1",
			args: []int{1},
			want: 1,
		},
		{
			name: "lcm 2",
			args: []int{1, 2},
			want: 2,
		},
		{
			name: "lcm 3",
			args: []int{1, 2, 3},
			want: 6,
		},
		{
			name: "lcm 4",
			args: []int{1, 2, 3, 4},
			want: 12,
		},
		{
			name: "lcm 5",
			args: []int{1, 2, 3, 4, 5},
			want: 60,
		},
		{
			name: "lcm 6",
			args: []int{1, 2, 3, 4, 5, 6},
			want: 60,
		},
		{
			name: "lcm 7",
			args: []int{1, 2, 3, 4, 5, 6, 7},
			want: 420,
		},
		{
			name: "lcm 8",
			args: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: 840,
		},
		{
			name: "lcm 9",
			args: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 2520,
		},
		{
			name: "lcm 10",
			args: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: 2520,
		},
		{
			name: "lcm 11",
			args: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
				11},
			want: 27720,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lcm(tt.args); got != tt.want {
				t.Errorf("Lcm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGcd(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "gcd 2",
			args: []int{1, 2},
			want: 1,
		},
		{
			name: "gcd 17 34",
			args: []int{17, 34},
			want: 17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gcd(tt.args[0], tt.args[1]); got != tt.want {
				t.Errorf("Gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}
