package container

import "testing"

func TestBag(t *testing.T) {
	b := NewBag[int]()
	if b.Contains(1) {
		t.Errorf("Bag should not contain 1")
	}
	b.Add(1)
	if !b.Contains(1) {
		t.Errorf("Bag should contain 1")
	}
	b.Add(2)
	b.Add(3)
	b.Add(4)
	b.Add(5)
	b.Remove(3)
	if !b.Contains(1) {
		t.Errorf("Bag should contain 1")
	}
	if !b.Contains(2) {
		t.Errorf("Bag should contain 2")
	}
	if b.Contains(3) {
		t.Errorf("Bag should not contain 3")
	}
	if !b.Contains(4) {
		t.Errorf("Bag should contain 4")
	}
	if !b.Contains(5) {
		t.Errorf("Bag should contain 5")
	}

	if b.Count(1) != 1 {
		t.Errorf("Bag should contain 1 once")
	}
	if b.Count(2) != 1 {
		t.Errorf("Bag should contain 2 once")
	}
	if b.Count(3) != 0 {
		t.Errorf("Bag should not contain 3")
	}
	if b.Count(4) != 1 {
		t.Errorf("Bag should contain 4 once")
	}
	if b.Count(5) != 1 {
		t.Errorf("Bag should contain 5 once")
	}

	b.Add(1)
	if b.Count(1) != 2 {
		t.Errorf("Bag should contain 1 twice")
	}

	b.Remove(1)
	if b.Count(1) != 1 {
		t.Errorf("Bag should contain 1 once")
	}

	b.Remove(1)
	if b.Count(1) != 0 {
		t.Errorf("Bag should not contain 1")
	}

	m := b.Values()
	if len(m) != 3 {
		t.Errorf("Bag should have 3 distinct values")
	}
	if m[1] != 0 {
		t.Errorf("Bag should not contain 1")
	}
	if m[2] != 1 {
		t.Errorf("Bag should contain 2 once")
	}
	if m[3] != 0 {
		t.Errorf("Bag should not contain 3")
	}
	if m[4] != 1 {
		t.Errorf("Bag should contain 4 once")
	}
	if m[5] != 1 {
		t.Errorf("Bag should contain 5 once")
	}

}
