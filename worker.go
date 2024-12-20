package mantis

import (
	"time"
)

type WorkerFunc struct {
	pool        *Pool
	tasks       chan func()
	recycleTime time.Time
}

func newWorkerFunc() {

}

func (w *WorkerFunc) lastUsedTime() time.Time {
	return w.recycleTime
}
func (w *WorkerFunc) setlastUsedTime(t time.Time) {
	w.recycleTime = t

}

func (w *WorkerFunc) take(task func()) error {
	w.tasks <- task
	return nil
}

func (w *WorkerFunc) run() {
	w.pool.incRunning()
	go func() {
		defer func() {
			if p := recover(); p != nil {
				w.pool.decRunning()
				w.pool.workerCache.Put(w)
			}
		}()
		for {
			f, ok := <-w.tasks
			if !ok {
				return
			}
			if f == nil {
				return
			}
			f()
			w.pool.RecycleWorkder(w)
		}
	}()
}
