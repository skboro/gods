package gods

// Ref https://golang.org/pkg/container/heap

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type Item struct {
	Value    interface{}
	Priority int64
	index    int // The index of the item in the heap. It is needed by update and is maintained by the heap.Interface methods.
}

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push pushes the element x onto the heap.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) Update(item *Item, value interface{}, priority int64) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.index)
}
