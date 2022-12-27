package grid

import "testing"

func TestGrid2D(t *testing.T) {
	g := NewGrid2D[int](3, 3, false)
	if g.Width() != 3 {
		t.Errorf("Grid should have width 3")
	}
	if g.Height() != 3 {
		t.Errorf("Grid should have height 3")
	}
	if g.Get(Coordinates{0, 0}) != 0 {
		t.Errorf("Grid should have nil at (0, 0)")
	}
	if g.Get(Coordinates{1, 1}) != 0 {
		t.Errorf("Grid should have nil at (1, 1)")
	}
	if g.Get(Coordinates{2, 2}) != 0 {
		t.Errorf("Grid should have nil at (2, 2)")
	}
	g.Set(Coordinates{0, 0}, 1)
	g.Set(Coordinates{1, 1}, 2)
	g.Set(Coordinates{2, 2}, 3)
	if g.Get(Coordinates{0, 0}) != 1 {
		t.Errorf("Grid should have 1 at (0, 0)")
	}
	if g.Get(Coordinates{1, 1}) != 2 {
		t.Errorf("Grid should have 2 at (1, 1)")
	}
	if g.Get(Coordinates{2, 2}) != 3 {
		t.Errorf("Grid should have 3 at (2, 2)")
	}
}

func TestGrid2DNeighbours(t *testing.T) {
	g := NewGrid2D[int](3, 3, false)
	g.Set(Coordinates{0, 0}, 1)
	g.Set(Coordinates{1, 1}, 2)
	g.Set(Coordinates{2, 2}, 3)
	neighbours := g.GetNeighbours8(Coordinates{1, 1})
	if len(neighbours) != 2 {
		t.Errorf("Grid should have 2 neighbours at (1, 1)")
	}
	if neighbours[0] != 1 {
		t.Errorf("Grid should have 1 as neighbour at (0, 0)")
	}
	if neighbours[1] != 3 {
		t.Errorf("Grid should have 0 as neighbour at (2, 2)")
	}
}

func TestGrid2DNeighboursWrap(t *testing.T) {
	g := NewGrid2D[int](3, 3, true)
	g.Set(Coordinates{0, 0}, 1)
	g.Set(Coordinates{1, 1}, 2)
	g.Set(Coordinates{2, 2}, 3)
	neighbours := g.GetNeighbours4(Coordinates{1, 1})
	if len(neighbours) != 0 {
		t.Errorf("Grid should have 0 neighbours at (1, 1)")
	}
}

func TestGrid2DMove(t *testing.T) {
	g := NewGrid2D[int](3, 3, false)
	g.Set(Coordinates{1, 1}, 2)
	g.Move(Coordinates{1, 1}, Coordinates{0, 1})
	if g.Get(Coordinates{1, 1}) != 0 {
		t.Errorf("Grid should have nil at (1, 1)")
	}
	if g.Get(Coordinates{1, 2}) != 2 {
		t.Errorf("Grid should have 2 at (1, 2)")
	}
	g.Move(Coordinates{1, 2}, Coordinates{0, 1})
	if g.Get(Coordinates{1, 2}) != 2 {
		t.Errorf("Grid should have 2 at (1, 2)")
	}
}

func TestGrid2DMoveWrap(t *testing.T) {
	g := NewGrid2D[int](3, 3, true)
	g.Set(Coordinates{1, 1}, 2)
	g.Move(Coordinates{1, 1}, Coordinates{0, 1})
	g.Move(Coordinates{1, 2}, Coordinates{0, 1})
	if g.Get(Coordinates{1, 1}) != 0 {
		t.Errorf("Grid should have nil at (1, 1)")
	}
	if g.Get(Coordinates{1, 2}) != 0 {
		t.Errorf("Grid should have nil at (1, 2)")
	}
	if g.Get(Coordinates{1, 0}) != 2 {
		t.Errorf("Grid should have 2 at (1, 0)")
	}
}
