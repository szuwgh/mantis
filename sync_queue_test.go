package mantis

import (
	"testing"
)

func TestSyncQueue(t *testing.T) {
	q := newSyncQueue[int](5)

	if !q.isEmpty() {
		t.Errorf("Expected queue to be empty")
	}

	err := q.push(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if q.isEmpty() {
		t.Errorf("Expected queue to not be empty")
	}

	val := q.pop()
	if val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}

	if !q.isEmpty() {
		t.Errorf("Expected queue to be empty after pop")
	}

	// Test queue full condition
	for i := 0; i < 5; i++ {
		err := q.push(i)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}

	if !q.isFull() {
		t.Errorf("Expected queue to be full")
	}

	err = q.push(6)
	if err == nil {
		t.Errorf("Expected error when pushing to full queue")
	}

	// Test queue empty condition after popping all elements
	for i := 0; i < 5; i++ {
		val := q.pop()
		if val != i {
			t.Errorf("Expected %v, got %v", i, val)
		}
	}

	if !q.isEmpty() {
		t.Errorf("Expected queue to be empty after popping all elements")
	}
}
