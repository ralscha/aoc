package geomutil

import (
	"aoc/internal/gridutil"
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		name string
		p1   gridutil.Coordinate
		p2   gridutil.Coordinate
		want float64
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
			want: 5, // 3-4-5 triangle
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Distance(tt.p1, tt.p2)
			if math.Abs(got-tt.want) > 1e-10 {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestArea(t *testing.T) {
	tests := []struct {
		name   string
		points []gridutil.Coordinate
		want   float64
	}{
		{
			name: "square",
			points: []gridutil.Coordinate{
				{Row: 0, Col: 0},
				{Row: 0, Col: 2},
				{Row: 2, Col: 2},
				{Row: 2, Col: 0},
			},
			want: 4,
		},
		{
			name: "triangle",
			points: []gridutil.Coordinate{
				{Row: 0, Col: 0},
				{Row: 0, Col: 2},
				{Row: 2, Col: 0},
			},
			want: 2,
		},
		{
			name:   "less than 3 points",
			points: []gridutil.Coordinate{{Row: 0, Col: 0}, {Row: 1, Col: 1}},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Area(tt.points)
			if math.Abs(got-tt.want) > 1e-10 {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPointInPolygon(t *testing.T) {
	square := []gridutil.Coordinate{
		{Row: 0, Col: 0},
		{Row: 0, Col: 2},
		{Row: 2, Col: 2},
		{Row: 2, Col: 0},
	}

	tests := []struct {
		name    string
		point   gridutil.Coordinate
		polygon []gridutil.Coordinate
		want    bool
	}{
		{
			name:    "point inside square",
			point:   gridutil.Coordinate{Row: 1, Col: 1},
			polygon: square,
			want:    true,
		},
		{
			name:    "point outside square",
			point:   gridutil.Coordinate{Row: 3, Col: 3},
			polygon: square,
			want:    false,
		},
		{
			name:    "point on edge",
			point:   gridutil.Coordinate{Row: 0, Col: 1},
			polygon: square,
			want:    true,
		},
		{
			name:    "point on vertex",
			point:   gridutil.Coordinate{Row: 0, Col: 0},
			polygon: square,
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointInPolygon(tt.point, tt.polygon); got != tt.want {
				t.Errorf("PointInPolygon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvexHull(t *testing.T) {
	tests := []struct {
		name   string
		points []gridutil.Coordinate
		want   []gridutil.Coordinate
	}{
		{
			name: "square with point inside",
			points: []gridutil.Coordinate{
				{Row: 0, Col: 0},
				{Row: 0, Col: 2},
				{Row: 2, Col: 2},
				{Row: 2, Col: 0},
				{Row: 1, Col: 1}, // inside point
			},
			want: []gridutil.Coordinate{
				{Row: 0, Col: 0},
				{Row: 0, Col: 2},
				{Row: 2, Col: 2},
				{Row: 2, Col: 0},
			},
		},
		{
			name: "triangle",
			points: []gridutil.Coordinate{
				{Row: 0, Col: 0},
				{Row: 0, Col: 2},
				{Row: 2, Col: 0},
			},
			want: []gridutil.Coordinate{
				{Row: 0, Col: 0},
				{Row: 0, Col: 2},
				{Row: 2, Col: 0},
			},
		},
		{
			name: "less than 3 points",
			points: []gridutil.Coordinate{
				{Row: 0, Col: 0},
				{Row: 1, Col: 1},
			},
			want: []gridutil.Coordinate{
				{Row: 0, Col: 0},
				{Row: 1, Col: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvexHull(tt.points)
			if len(got) != len(tt.want) {
				t.Errorf("ConvexHull() length = %v, want %v", len(got), len(tt.want))
				return
			}

			// Check area is the same (allows for different vertex order)
			gotArea := Area(got)
			wantArea := Area(tt.want)
			if math.Abs(gotArea-wantArea) > 1e-10 {
				t.Errorf("ConvexHull() area = %v, want %v", gotArea, wantArea)
			}
		})
	}
}
