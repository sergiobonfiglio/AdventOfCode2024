package utils

import "container/heap"

type Item[T any] struct {
	Data  T
	Value int // The Value of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A MinHeap implements heap.Interface and holds Items.
// type MinHeap[T any] []*Item[T]
type MinHeap[T any] struct {
	heap heapData[T]
}

type heapData[T any] []*Item[T]

var _ heap.Interface = &heapData[int]{}

func NewMinHeap[T any]() *MinHeap[T] {
	h := &MinHeap[T]{
		heap: make([]*Item[T], 0),
	}
	//heap.Init(&h.heap)
	return h
}

func (mh *MinHeap[T]) Len() int {
	return len(mh.heap)
}

func (mh *MinHeap[T]) HeapPush(data T, value int) *Item[T] {
	item := &Item[T]{Data: data, Value: value}
	heap.Push(&mh.heap, item)
	return item
}

func (mh *MinHeap[T]) HeapPop() (T, int) {
	item := heap.Pop(&mh.heap).(*Item[T])
	return item.Data, item.Value
}

func (mh *MinHeap[T]) Update(item *Item[T], value int) {
	item.Value = value
	heap.Fix(&mh.heap, item.index)
}

func (mh *heapData[T]) Len() int { return len(*mh) }

func (mh *heapData[T]) Less(i, j int) bool {
	return (*mh)[i].Value < (*mh)[j].Value
}

func (mh *heapData[T]) Swap(i, j int) {
	(*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
	(*mh)[i].index = i
	(*mh)[j].index = j
}

func (mh *heapData[T]) Push(x any) {
	n := len(*mh)
	item := x.(*Item[T])
	item.index = n
	*mh = append(*mh, item)
}

func (mh *heapData[T]) Pop() any {
	old := *mh
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*mh = old[0 : n-1]
	return item
}
