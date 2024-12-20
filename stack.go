package mantis

import (
	"time"
)

type Stack[T Workder] struct {
	items  []T
	expiry []T
}

func newStack[T Workder](size int) *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0, size),
	}
}

func (wq *Stack[T]) len() int {
	return len(wq.items)
}

func (wq *Stack[T]) isEmpty() bool {
	return len(wq.items) == 0
}

func (wq *Stack[T]) insert(w T) error {
	wq.items = append(wq.items, w)
	return nil
}

func (wq *Stack[T]) detach() T {
	l := wq.len()
	if l == 0 {
		var zero T
		return zero
	}

	w := wq.items[l-1]
	var zero T
	wq.items[l-1] = zero // avoid memory leaks
	wq.items = wq.items[:l-1]

	return w
}

func (wq *Stack[T]) refresh(duration time.Duration) []T {
	n := wq.len()
	if n == 0 {
		return nil
	}
	var zero T
	expiryTime := time.Now().Add(-duration)
	index := wq.binarySearch(0, n-1, expiryTime)

	wq.expiry = wq.expiry[:0]
	if index != -1 {
		wq.expiry = append(wq.expiry, wq.items[:index+1]...)
		m := copy(wq.items, wq.items[index+1:])
		for i := m; i < n; i++ {
			wq.items[i] = zero
		}
		wq.items = wq.items[:m]
	}
	return wq.expiry
}

func (wq *Stack[T]) binarySearch(l, r int, expiryTime time.Time) int {
	for l <= r {
		mid := l + ((r - l) >> 1) // avoid overflow when computing mid
		if expiryTime.Before(wq.items[mid].lastUsedTime()) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}
