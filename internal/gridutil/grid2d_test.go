package gridutil

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

func TestGrid2DCoordinateBased(t *testing.T) {
	g := NewGrid2D[int](false)
	coord := Coordinate{Row: 1, Col: 1}

	// Test SetMaxRowColC
	g.SetMaxRowColC(coord)
	if g.maxRow != 1 || g.maxCol != 1 {
		t.Errorf("SetMaxRowColC failed, expected maxRow=1, maxCol=1, got maxRow=%d, maxCol=%d", g.maxRow, g.maxCol)
	}

	// Setup grid for neighbour tests
	g.Set(0, 0, 1)
	g.Set(1, 1, 2)
	g.Set(2, 2, 3)

	// Test PeekC
	val, isZero := g.PeekC(coord, DirectionNW)
	if isZero || val != 1 {
		t.Errorf("PeekC failed, expected value=1, isZero=false, got value=%d, isZero=%v", val, isZero)
	}

	// Test GetNeighbours8C
	neighbours8 := g.GetNeighbours8C(coord)
	if len(neighbours8) != 2 {
		t.Errorf("GetNeighbours8C should return 2 neighbours, got %d", len(neighbours8))
	}
	if neighbours8[0] != 1 || neighbours8[1] != 3 {
		t.Errorf("GetNeighbours8C returned wrong values")
	}

	// Test GetNeighbours4C
	neighbours4 := g.GetNeighbours4C(coord)
	if len(neighbours4) != 0 {
		t.Errorf("GetNeighbours4C should return 0 neighbours for this test case, got %d", len(neighbours4))
	}

	// Test GetNeighboursC with custom directions
	customDirections := []Direction{DirectionNW, DirectionSE}
	neighboursCustom := g.GetNeighboursC(coord, customDirections)
	if len(neighboursCustom) != 2 {
		t.Errorf("GetNeighboursC should return 2 neighbours, got %d", len(neighboursCustom))
	}
	if neighboursCustom[0] != 1 || neighboursCustom[1] != 3 {
		t.Errorf("GetNeighboursC returned wrong values")
	}
}
