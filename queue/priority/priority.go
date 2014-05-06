package priority

import (
	"container/heap"

	"github.com/joshuarubin/gil"
)

// A pqHeap implements heap.Interface and holds priorityQueueItems.
type pqHeap gil.Slice

// required for heap.Interface
func (pq pqHeap) Len() int {
	return len(pq)
}

// required for heap.Interface
func (pq pqHeap) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Less(pq[j])
}

// required for heap.Interface
func (pq pqHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// required for heap.Interface
func (pq *pqHeap) Push(x interface{}) {
	*pq = append(*pq, x.(gil.Interface))
}

// required for heap.Interface
func (pq *pqHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// PriorityQueue is a generic priority queue implementation.
// Items all implement gil.Interface and so are ordered according to their Less method
type PriorityQueue struct {
	heap *pqHeap
}

// New creates and initializes a PriorityQueue
func New() *PriorityQueue {
	q := &PriorityQueue{
		heap: &pqHeap{},
	}
	heap.Init(q.heap)
	return q
}

// Push an item into the queue
func (q *PriorityQueue) Push(value gil.Interface) gil.Queue {
	heap.Push(q.heap, value)
	return q
}

// Pop off the next item in the queue
func (q *PriorityQueue) Pop() gil.Interface {
	return heap.Pop(q.heap).(gil.Interface)
}

// Len returns the number of items in the queue
func (q *PriorityQueue) Len() int {
	return q.heap.Len()
}

// Peek at the next item in the queue without removing it
func (q *PriorityQueue) Peek() gil.Interface {
	return (*q.heap)[0]
}
