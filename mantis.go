package mantis

import (
	"errors"
	"time"
)

var (
	// Error types for the Ants API.
	//---------------------------------------------------------------------------
	// ErrInvalidPoolSize will be returned when setting a negative number as pool capacity.
	errQueueIsFull = errors.New("queue is full")
)

var DefaultPool *Pool

func init() {

}

type WorkerQueue[T interface{}] interface {
	pop() T
	push(T)
}

type Workder interface {
	run()
	take(func()) error
	lastUsedTime() time.Time
	setlastUsedTime(time.Time)
}

func Go(f func()) {
	DefaultPool.submit(f)
}

func Join(f1, f2 func()) {

}
