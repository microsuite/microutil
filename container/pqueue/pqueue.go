package pqueue

import "container/heap"

// item respresents an item in the priority queue.
type Item struct {
	// the value of the item
	Value interface{}

	// the priority of the item in the queue
	Priority int64

	// the index of the item in the heap
	Index int
}

// priorityQueue implements a priority queue.
type PriorityQueue []*Item

// NewPriorityQueue returns a new priority queue with the given capacity.
func NewPriorityQueue(cap int) PriorityQueue {
	return make(PriorityQueue, 0, cap)
}

// Len returns the length of the priority queue.
func (pq PriorityQueue) Len() int { return len(pq) }

// Less returns whether the item at index i has a higher priority than the item at index j.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// Swap swaps the items at indices i and j.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push adds an item to the priority queue.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	c := cap(*pq)
	if n+1 > c {
		npq := make(PriorityQueue, n, c*2)
		copy(npq, *pq)
		*pq = npq
	}
	*pq = (*pq)[0 : n+1]
	item := x.(*Item)
	item.Index = n
	(*pq)[n] = item
}

// Pop removes the highest priority item from the priority queue.
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	c := cap(*pq)
	if n < (c/2) && c > 25 {
		npq := make(PriorityQueue, n, c/2)
		copy(npq, *pq)
		*pq = npq
	}
	item := (*pq)[n-1]
	item.Index = -1
	*pq = (*pq)[0 : n-1]
	return item
}

// Update updates the value and priority of an item in the priority queue.
func (pq *PriorityQueue) Update(item *Item, value string, priority int64) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

// PeekAndShift returns the highest priority item from the priority queue and
// shifts the priority of the remaining items.
func (pq *PriorityQueue) PeekAndShift(max int64) (*Item, int64) {
	if pq.Len() == 0 {
		return nil, 0
	}

	item := (*pq)[0]
	if item.Priority > max {
		return nil, item.Priority - max
	}
	heap.Remove(pq, 0)

	return item, 0
}
