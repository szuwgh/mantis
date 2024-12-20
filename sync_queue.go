package mantis

import (
	"sync/atomic"
)

type SyncQueue[T any] struct {
	items []T
	head  int32
	tail  int32
	sz    int32
}

func newSyncQueue[T interface{}](size int) SyncQueue[T] {
	return SyncQueue[T]{
		items: make([]T, size+1),
		head:  0,
		tail:  0,
		sz:    int32(size + 1),
	}
}

func (q *SyncQueue[T]) size() int32 {
	return q.sz - 1
}

func (q *SyncQueue[T]) isEmpty() bool {
	return atomic.LoadInt32(&q.tail) == atomic.LoadInt32(&q.head)
}

func (q *SyncQueue[T]) isFull() bool {
	return (atomic.LoadInt32(&q.tail)+1)%q.sz == atomic.LoadInt32(&q.head)
}

func (q *SyncQueue[T]) pop() T {
	for {
		head := atomic.LoadInt32(&q.head)
		if head == atomic.LoadInt32(&q.tail) {
			var zero T
			return zero
		}
		// 尝试写入数据
		if atomic.CompareAndSwapInt32(&q.head, head, (head+1)%q.sz) {
			w := q.items[head]
			var zero T
			q.items[head] = zero
			return w
		}
	}
}

func (q *SyncQueue[T]) push(t T) error {
	for {
		tail := atomic.LoadInt32(&q.tail)
		next_tail := (tail + 1) % q.sz
		if next_tail == atomic.LoadInt32(&q.head) {
			return errQueueIsFull
		}
		if atomic.CompareAndSwapInt32(&q.tail, tail, next_tail) {
			q.items[tail] = t
			return nil
		}
	}
}
