package container

import "testing"

func TestPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	if q.Len() != 0 {
		t.Errorf("PriorityQueue should be empty")
	}
	if !q.IsEmpty() {
		t.Errorf("PriorityQueue should be empty")
	}
	q.Push(1, 5)
	q.Push(2, 4)
	q.Push(3, 6)
	q.Push(4, 1)
	q.Push(5, 0)
	if q.Len() != 5 {
		t.Errorf("PriorityQueue should have size 5")
	}
	if q.Pop() != 3 {
		t.Errorf("PriorityQueue should pop 3")
	}
	if q.Len() != 4 {
		t.Errorf("PriorityQueue should have size 4")
	}
	if q.Pop() != 1 {
		t.Errorf("PriorityQueue should pop 1")
	}
	if q.Len() != 3 {
		t.Errorf("PriorityQueue should have size 3")
	}
	if q.Pop() != 2 {
		t.Errorf("PriorityQueue should pop 2")
	}
	if q.Len() != 2 {
		t.Errorf("PriorityQueue should have size 2")
	}
	if q.Pop() != 4 {
		t.Errorf("PriorityQueue should pop 4")
	}
	if q.Len() != 1 {
		t.Errorf("PriorityQueue should have size 1")
	}
	if q.IsEmpty() {
		t.Errorf("PriorityQueue should not be empty")
	}
	if q.Pop() != 5 {
		t.Errorf("PriorityQueue should pop 5")
	}
	if q.Len() != 0 {
		t.Errorf("PriorityQueue should be empty")
	}
	if !q.IsEmpty() {
		t.Errorf("PriorityQueue should be empty")
	}
}
