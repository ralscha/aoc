package container

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	if q.Len() != 0 {
		t.Errorf("Queue should be empty")
	}
	if !q.IsEmpty() {
		t.Errorf("Queue should be empty")
	}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	if q.Len() != 5 {
		t.Errorf("Queue should have size 5")
	}
	if q.Pop() != 1 {
		t.Errorf("Queue should pop 1")
	}
	if q.Len() != 4 {
		t.Errorf("Queue should have size 4")
	}
	if q.Pop() != 2 {
		t.Errorf("Queue should pop 2")
	}
	if q.Len() != 3 {
		t.Errorf("Queue should have size 3")
	}
	if q.Pop() != 3 {
		t.Errorf("Queue should pop 3")
	}
	if q.Len() != 2 {
		t.Errorf("Queue should have size 2")
	}
	if q.Pop() != 4 {
		t.Errorf("Queue should pop 4")
	}
	if q.Len() != 1 {
		t.Errorf("Queue should have size 1")
	}
	if q.IsEmpty() {
		t.Errorf("Queue should not be empty")
	}
	if q.Pop() != 5 {
		t.Errorf("Queue should pop 5")
	}
	if q.Len() != 0 {
		t.Errorf("Queue should be empty")
	}
	if !q.IsEmpty() {
		t.Errorf("Queue should be empty")
	}
}
