package delayqueue

import (
	"container/heap"
	"sync"
	"sync/atomic"
	"time"

	"github.com/microsuite/microutil/container/pqueue"
)

// DelayQueue is a priority queue that implements a delay queue.
type DelayQueue struct {
	Chan chan interface{}

	mu sync.Mutex
	pq pqueue.PriorityQueue

	sleeping int32
	wakeupC  chan struct{}
}

// NewDelayQueue creates a new delay queue.
func NewDelayQueue(size int) *DelayQueue {
	return &DelayQueue{
		Chan:    make(chan interface{}),
		pq:      pqueue.NewPriorityQueue(size),
		wakeupC: make(chan struct{}),
	}
}

// Add adds an element to the delay queue.
func (dq *DelayQueue) Add(elem interface{}, expiration int64) {
	item := &pqueue.Item{
		Value:    elem,
		Priority: expiration,
	}

	dq.mu.Lock()
	heap.Push(&dq.pq, item)
	index := item.Index
	dq.mu.Unlock()

	// If the element is the first one in the queue, wake up the Poll method.
	if index == 0 {
		if atomic.CompareAndSwapInt32(&dq.sleeping, 1, 0) {
			dq.wakeupC <- struct{}{}
		}
	}
}

// Poll polls an element from the delay queue.
func (dq *DelayQueue) Poll(quit chan struct{}, nowF func() int64) {
	for {
		// Obtains the current time.
		now := nowF()

		dq.mu.Lock()
		// Peeks the first element in the queue.
		item, delta := dq.pq.PeekAndShift(now)
		if item == nil {
			// If there is no expired element, reset the delay queue sleep state.
			// Ensure atomicity between Poll and Add operations.
			atomic.StoreInt32(&dq.sleeping, 1)
		}
		dq.mu.Unlock()

		if item == nil {
			if delta == 0 {
				// No element is expired, wait for the next element.
				select {
				case <-dq.wakeupC:
					// Wait for the first element to be written.
					continue
				case <-quit:
					goto exit
				}
			} else if delta > 0 {
				// delta > 0 , There is at least one item in the delay queue waiting to be processed.
				select {
				case <-dq.wakeupC:
					continue
				case <-time.After(time.Duration(delta) * time.Second):
					// Waiting for the "earliest" element in the queue to expire
					if atomic.SwapInt32(&dq.sleeping, 0) == 0 {
						<-dq.wakeupC
					}
					continue
				case <-quit:
					goto exit
				}
			}
		}

		select {
		case dq.Chan <- item.Value: // Fetch the expiring element and send the expiring element through the delay queue C channel
		case <-quit:
			goto exit
		}
	}

exit:
	// Reset the states
	atomic.StoreInt32(&dq.sleeping, 0)
}
