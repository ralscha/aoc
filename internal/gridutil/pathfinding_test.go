package gridutil

import (
	"testing"
)

func TestBFS(t *testing.T) {
	tests := []struct {
		name      string
		grid      []string
		start     Coordinate
		dirs      []Direction
		wantPath  []Coordinate
		wantFound bool
	}{
		{
			name: "simple path",
			grid: []string{
				"S...",
				"....",
				"...G",
			},
			start: Coordinate{0, 0},
			dirs:  Get4Directions(),
			wantPath: []Coordinate{
				{0, 0},
				{1, 0},
				{2, 0},
				{2, 1},
				{2, 2},
				{2, 3},
			},
			wantFound: true,
		},
		{
			name: "path with obstacles",
			grid: []string{
				"S.##.",
				"#.#..",
				"....G",
			},
			start: Coordinate{0, 0},
			dirs:  Get4Directions(),
			wantPath: []Coordinate{
				{0, 0},
				{0, 1},
				{1, 1},
				{2, 1},
				{2, 2},
				{2, 3},
				{2, 4},
			},
			wantFound: true,
		},
		{
			name: "no path possible",
			grid: []string{
				"S#G",
				"###",
				"...",
			},
			start:     Coordinate{0, 0},
			dirs:      Get4Directions(),
			wantPath:  nil,
			wantFound: false,
		},
		{
			name: "start is goal",
			grid: []string{
				"G",
			},
			start: Coordinate{0, 0},
			dirs:  Get4Directions(),
			wantPath: []Coordinate{
				{0, 0},
			},
			wantFound: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create grid from string representation
			grid := NewGrid2D[rune](false)
			for row, line := range tt.grid {
				for col, ch := range line {
					grid.Set(row, col, ch)
				}
			}

			// Define goal checker function
			isGoal := func(pos Coordinate, val rune) bool {
				return val == 'G'
			}

			// Define obstacle checker function
			isObstacle := func(pos Coordinate, val rune) bool {
				return val == '#'
			}

			// Run ShortestPathWithBFS
			result, found := grid.ShortestPathWithBFS(tt.start, isGoal, isObstacle, tt.dirs)

			// Check if path was found when expected
			if found != tt.wantFound {
				t.Errorf("ShortestPathWithBFS() found = %v, want %v", found, tt.wantFound)
				return
			}

			// If path should be found, verify it
			if tt.wantFound {
				if len(result.Path) != len(tt.wantPath) {
					t.Errorf("ShortestPathWithBFS() path length = %v, want %v", len(result.Path), len(tt.wantPath))
					return
				}

				// Check each coordinate in path
				for i := range result.Path {
					if result.Path[i] != tt.wantPath[i] {
						t.Errorf("ShortestPathWithBFS() path[%d] = %v, want %v", i, result.Path[i], tt.wantPath[i])
					}
				}

				// Verify path cost matches path length
				if result.Cost != float64(len(result.Path)-1) {
					t.Errorf("ShortestPathWithBFS() cost = %v, want %v", result.Cost, float64(len(result.Path)-1))
				}
			}

			// Verify path is valid by checking each step is adjacent
			if found && len(result.Path) > 1 {
				for i := 1; i < len(result.Path); i++ {
					curr := result.Path[i]
					prev := result.Path[i-1]
					rowDiff := abs(curr.Row - prev.Row)
					colDiff := abs(curr.Col - prev.Col)

					// Check if step is valid according to allowed directions
					validStep := false
					for _, dir := range tt.dirs {
						if curr.Row == prev.Row+dir.Row && curr.Col == prev.Col+dir.Col {
							validStep = true
							break
						}
					}
					if !validStep {
						t.Errorf("Invalid path step from %v to %v", prev, curr)
					}

					// For diagonal movement, both row and col can change by 1
					maxDiff := 1
					if rowDiff > maxDiff || colDiff > maxDiff {
						t.Errorf("Invalid step size from %v to %v", prev, curr)
					}
				}
			}
		})
	}
}

func TestDijkstra(t *testing.T) {
	tests := []struct {
		name      string
		grid      []string
		start     Coordinate
		dirs      []Direction
		wantPath  []Coordinate
		wantCost  float64
		wantFound bool
	}{
		{
			name: "simple weighted path",
			grid: []string{
				"S123",
				"4567",
				"89AG",
			},
			start: Coordinate{0, 0},
			dirs:  Get4Directions(),
			wantPath: []Coordinate{
				{0, 0}, // S
				{0, 1}, // 1
				{0, 2}, // 2
				{0, 3}, // 3
				{1, 3}, // 7
				{2, 3}, // G
			},
			wantCost:  13.0, // Sum of path costs excluding start and goal
			wantFound: true,
		},
		{
			name: "path with high cost barrier",
			grid: []string{
				"S123",
				"9991",
				"123G",
			},
			start: Coordinate{0, 0},
			dirs:  Get4Directions(),
			wantPath: []Coordinate{
				{0, 0}, // S
				{0, 1}, // 1
				{0, 2}, // 2
				{0, 3}, // 3
				{1, 3}, // 1
				{2, 3}, // G
			},
			wantCost:  7.0, // Sum of path costs excluding start and goal
			wantFound: true,
		},
		{
			name: "path blocked by walls",
			grid: []string{
				"S#G",
				"###",
				"123",
			},
			start:     Coordinate{0, 0},
			dirs:      Get4Directions(),
			wantPath:  nil,
			wantCost:  0,
			wantFound: false,
		},
		{
			name: "start is goal",
			grid: []string{
				"G",
			},
			start: Coordinate{0, 0},
			dirs:  Get4Directions(),
			wantPath: []Coordinate{
				{0, 0},
			},
			wantCost:  0,
			wantFound: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create grid from string representation
			grid := NewGrid2D[rune](false)
			for row, line := range tt.grid {
				for col, ch := range line {
					grid.Set(row, col, ch)
				}
			}

			// Define goal checker function
			isGoal := func(pos Coordinate, val rune) bool {
				return val == 'G'
			}

			// Define obstacle checker function
			isObstacle := func(pos Coordinate, val rune) bool {
				return val == '#'
			}

			// Define cost function - convert rune to numeric cost
			getCost := func(from, to Coordinate, fromVal, toVal rune) float64 {
				if toVal >= '1' && toVal <= '9' {
					return float64(toVal - '0')
				}
				if toVal == 'A' {
					return 10.0
				}
				return 0.0 // Start (S) and Goal (G) have no cost
			}

			// Run ShortestPathWithDijkstra
			result, found := grid.ShortestPathWithDijkstra(tt.start, isGoal, isObstacle, getCost, tt.dirs)

			// Check if path was found when expected
			if found != tt.wantFound {
				t.Errorf("ShortestPathWithDijkstra() found = %v, want %v", found, tt.wantFound)
				return
			}

			// If path should be found, verify it
			if tt.wantFound {
				if len(result.Path) != len(tt.wantPath) {
					t.Errorf("ShortestPathWithDijkstra() path length = %v, want %v", len(result.Path), len(tt.wantPath))
					return
				}

				// Check each coordinate in path
				for i := range result.Path {
					if result.Path[i] != tt.wantPath[i] {
						t.Errorf("ShortestPathWithDijkstra() path[%d] = %v, want %v", i, result.Path[i], tt.wantPath[i])
					}
				}

				// Verify total cost
				if result.Cost != tt.wantCost {
					t.Errorf("ShortestPathWithDijkstra() cost = %v, want %v", result.Cost, tt.wantCost)
				}
			}

			// Verify path is valid by checking each step is adjacent
			if found && len(result.Path) > 1 {
				for i := 1; i < len(result.Path); i++ {
					curr := result.Path[i]
					prev := result.Path[i-1]
					rowDiff := abs(curr.Row - prev.Row)
					colDiff := abs(curr.Col - prev.Col)

					// Check if step is valid according to allowed directions
					validStep := false
					for _, dir := range tt.dirs {
						if curr.Row == prev.Row+dir.Row && curr.Col == prev.Col+dir.Col {
							validStep = true
							break
						}
					}
					if !validStep {
						t.Errorf("Invalid path step from %v to %v", prev, curr)
					}

					// For diagonal movement, both row and col can change by 1
					maxDiff := 1
					if rowDiff > maxDiff || colDiff > maxDiff {
						t.Errorf("Invalid step size from %v to %v", prev, curr)
					}
				}
			}
		})
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
