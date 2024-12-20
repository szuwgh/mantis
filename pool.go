package mantis

import (
	"sync"
	"sync/atomic"
	"time"
)

type Pool struct {
	runningCount   int32
	mu             sync.Locker
	cap            int
	workers        WorkerQueue[Workder]
	cond           *sync.Cond
	expiryDuration time.Duration
	workerCache    sync.Pool
}

func NewPool(cap int) {

}

func (p *Pool) submit(f func()) {
	w := p.getFreeWorkder()
	w.take(f)
}

func (p *Pool) incRunning() {
	atomic.AddInt32(&p.runningCount, 1)
}

func (p *Pool) decRunning() {
	atomic.AddInt32(&p.runningCount, -1)
}

func (p *Pool) runingCount() int {
	return int(atomic.LoadInt32(&p.runningCount))
}

func (p *Pool) getFreeWorkder() (w Workder) {
	p.mu.Lock()
	for {
		if w = p.workers.pop(); w != nil {
			p.mu.Unlock()
			return
		}
		if p.runingCount() < p.cap {
			p.mu.Unlock()
			w = p.workerCache.Get().(Workder)
			w.run()
			return w
		}
		p.mu.Unlock()
		p.cond.Wait()
	}
}

// 回收workder
func (p *Pool) RecycleWorkder(worker Workder) {
	worker.setlastUsedTime(time.Now())
	p.mu.Lock()
	defer p.mu.Unlock()
	p.workers.push(worker)
	p.cond.Signal()
}
