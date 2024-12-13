package geomutil

import (
	"aoc/internal/gridutil"
	"testing"
)

func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		name string
		p1   gridutil.Coordinate
		p2   gridutil.Coordinate
		want int
	}{
		{
			name: "same point",
			p1:   gridutil.Coordinate{Row: 0, Col: 0},
			p2:   gridutil.Coordinate{Row: 0, Col: 0},
			want: 0,
		},
		{
			name: "horizontal",
			p1:   gridutil.Coordinate{Row: 0, Col: 0},
			p2:   gridutil.Coordinate{Row: 0, Col: 3},
			want: 3,
		},
		{
			name: "vertical",
			p1:   gridutil.Coordinate{Row: 0, Col: 0},
			p2:   gridutil.Coordinate{Row: 4, Col: 0},
			want: 4,
		},
		{
			name: "diagonal",
			p1:   gridutil.Coordinate{Row: 0, Col: 0},
			p2:   gridutil.Coordinate{Row: 3, Col: 4},
			want: 7, // |3| + |4|
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ManhattanDistance(tt.p1, tt.p2); got != tt.want {
				t.Errorf("ManhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
