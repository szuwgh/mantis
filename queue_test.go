package mantis

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := newQueue[int](3)

	if !q.isEmpty() {
		t.Errorf("Expected queue to be empty")
	}

	err := q.push(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = q.push(2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = q.push(3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !q.isFull() {
		t.Errorf("Expected queue to be full")
	}

	err = q.push(4)
	if err == nil {
		t.Errorf("Expected error when pushing to full queue")
	}

	item := q.pop()
	if item != 1 {
		t.Errorf("Expected 1, got %v", item)
	}

	item = q.pop()
	if item != 2 {
		t.Errorf("Expected 2, got %v", item)
	}

	item = q.pop()
	if item != 3 {
		t.Errorf("Expected 3, got %v", item)
	}

	if !q.isEmpty() {
		t.Errorf("Expected queue to be empty")
	}

	item = q.pop()
	if item != 0 {
		t.Errorf("Expected 0, got %v", item)
	}
}
