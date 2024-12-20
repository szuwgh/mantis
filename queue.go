package mantis

type Queue[T any] struct {
	items []T
	head  int32
	tail  int32
	sz    int32
}

func newQueue[T interface{}](size int) Queue[T] {
	return Queue[T]{
		items: make([]T, size+1),
		head:  0,
		tail:  0,
		sz:    int32(size + 1),
	}
}

func (q *Queue[T]) size() int32 {
	return q.sz - 1
}

func (q *Queue[T]) isEmpty() bool {
	return q.tail == q.head
}

func (q *Queue[T]) isFull() bool {
	return (q.tail+1)%q.sz == q.head
}

func (q *Queue[T]) pop() T {

	head := q.head
	if head == q.tail {
		var zero T
		return zero
	}
	// 尝试写入数据
	q.head = (head + 1) % q.sz
	w := q.items[head]
	var zero T
	q.items[head] = zero
	return w

}

func (q *Queue[T]) push(t T) error {
	tail := q.tail
	next_tail := (tail + 1) % q.sz
	if next_tail == q.head {
		return errQueueIsFull
	}
	q.items[tail] = t
	q.tail = next_tail
	return nil
}
