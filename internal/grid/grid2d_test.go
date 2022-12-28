package grid

import "testing"

func TestGrid2D(t *testing.T) {
	g := NewGrid2D[int](false)
	_, ok := g.Get(0, 0)
	if ok {
		t.Errorf("Grid should have nil at (0, 0)")
	}
	_, ok = g.Get(1, 1)
	if ok {
		t.Errorf("Grid should have nil at (1, 1)")
	}
	_, ok = g.Get(2, 2)
	if ok {
		t.Errorf("Grid should have nil at (2, 2)")
	}
	g.Set(0, 0, 1)
	g.Set(1, 1, 2)
	g.Set(2, 2, 3)
	v, ok := g.Get(0, 0)
	if v != 1 {
		t.Errorf("Grid should have 1 at (0, 0)")
	}
	v, ok = g.Get(1, 1)
	if v != 2 {
		t.Errorf("Grid should have 2 at (1, 1)")
	}
	v, ok = g.Get(2, 2)
	if v != 3 {
		t.Errorf("Grid should have 3 at (2, 2)")
	}
	if g.Width() != 3 {
		t.Errorf("Grid should have Width 3")
	}
	if g.Height() != 3 {
		t.Errorf("Grid should have Height 3")
	}
}

func TestGrid2DNeighbours(t *testing.T) {
	g := NewGrid2D[int](false)
	g.Set(0, 0, 1)
	g.Set(1, 1, 2)
	g.Set(2, 2, 3)
	neighbours := g.GetNeighbours8(1, 1)
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
	g := NewGrid2D[int](true)
	g.Set(0, 0, 1)
	g.Set(1, 1, 2)
	g.Set(2, 2, 3)
	neighbours := g.GetNeighbours4(1, 1)
	if len(neighbours) != 0 {
		t.Errorf("Grid should have 0 neighbours at (1, 1)")
	}
}

func TestGrid2DMove(t *testing.T) {
	g := NewGrid2D[int](false)
	g.SetMaxRowCol(1, 2)
	g.Set(1, 1, 2)
	g.Move(1, 1, DirectionE)
	_, ok := g.Get(1, 1)
	if ok {
		t.Errorf("Grid should have nil at (1, 1)")
	}
	v, ok := g.Get(1, 2)
	if v != 2 {
		t.Errorf("Grid should have 2 at (1, 2)")
	}
	g.Move(1, 2, DirectionE)
	v, ok = g.Get(1, 2)
	if v != 2 {
		t.Errorf("Grid should have 2 at (1, 2)")
	}
}

func TestGrid2DMoveWrap(t *testing.T) {
	g := NewGrid2D[int](true)
	g.SetMinRowCol(0, 0)
	g.SetMaxRowCol(2, 2)
	g.Set(1, 1, 2)
	g.Move(1, 1, DirectionE)
	g.Move(1, 2, DirectionE)
	v, ok := g.Get(1, 1)
	if ok {
		t.Errorf("Grid should have nil at (1, 1)")
	}
	v, ok = g.Get(1, 2)
	if ok {
		t.Errorf("Grid should have nil at (1, 2)")
	}
	v, ok = g.Get(1, 0)
	if v != 2 {
		t.Errorf("Grid should have 2 at (1, 0)")
	}
}
