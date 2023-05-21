package container

import "testing"

func TestSet(t *testing.T) {
	s := NewSet[int]()
	if s.Contains(1) {
		t.Errorf("Set should not contain 1")
	}
	s.Add(1)
	if !s.Contains(1) {
		t.Errorf("Set should contain 1")
	}
	s.Add(2)
	s.Add(3)
	s.Add(4)
	s.Add(5)
	s.Remove(3)
	if !s.Contains(1) {
		t.Errorf("Set should contain 1")
	}
	if !s.Contains(2) {
		t.Errorf("Set should contain 2")
	}
	if s.Contains(3) {
		t.Errorf("Set should not contain 3")
	}
	if !s.Contains(4) {
		t.Errorf("Set should contain 4")
	}
	if !s.Contains(5) {
		t.Errorf("Set should contain 5")
	}
}
