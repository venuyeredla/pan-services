package stack_queue

import (
	"fmt"

	"github.com/venuyeredla/pan-services/pkg/dsa/errors"
	"github.com/venuyeredla/pan-services/pkg/dsa/types"
)

type PqEntry struct {
	Item     any
	Priority int
}

type PriorityQueue struct {
	min  bool
	list []PqEntry
}

func NewPQ(size int, is_min bool) *PriorityQueue {
	return &PriorityQueue{
		min:  is_min,
		list: make([]PqEntry, 0, size),
	}
}

func (pq *PriorityQueue) Push(entry PqEntry) {
	pq.list = append(pq.list, entry)
	pq.fixUp(len(pq.list) - 1)
}

// Notes:
// Parent of i : (i+1)/2 - 1
// Left Child of i : (i+1)*2 - 1
// Right Child of i : (i+1)*2
func (h *PriorityQueue) fixUp(k int) {
	parent := (k+1)/2 - 1
	for k > 0 {
		if h.gte(parent, k) {
			return
		}
		h.list[parent], h.list[k] = h.list[k], h.list[parent]
		k = parent
		parent = (k+1)/2 - 1
	}
}

// Pop the highest (or lowest) priority item
func (pq *PriorityQueue) Pop() PqEntry {
	var entry PqEntry
	if len(pq.list) == 0 {
		return entry
	}
	entry = pq.list[0]
	pq.list[0] = pq.list[len(pq.list)-1]
	pq.list = pq.list[:len(pq.list)-1]
	pq.fixDown(0)
	return entry
}

func (h *PriorityQueue) fixDown(k int) {
	kid := (k+1)*2 - 1
	for kid < len(h.list) {
		if kid+1 < len(h.list) && h.lt(kid, kid+1) {
			kid++
		}
		if h.gte(k, kid) {
			break
		}
		h.list[kid], h.list[k] = h.list[k], h.list[kid]
		k = kid
		kid = (k+1)*2 - 1
	}
}

// Peek at the highest (or lowest) priority item
func (pq *PriorityQueue) Peek() PqEntry {
	var entry PqEntry
	if len(pq.list) == 0 {
		return entry
	}
	return pq.list[0]
}

// How many items in the heap?
func (pq *PriorityQueue) Size() int {
	return len(pq.list)
}

// Is this a min heap?
func (pq *PriorityQueue) MinHeap() bool {
	return pq.min
}

// Is this a max heap?
func (pq *PriorityQueue) MaxHeap() bool {
	return !pq.min
}

func (h *PriorityQueue) gte(i, j int) bool {
	if h.min {
		return h.list[i].Priority <= h.list[j].Priority
	} else {
		return h.list[i].Priority >= h.list[j].Priority
	}
}

func (h *PriorityQueue) lt(i, j int) bool {
	if h.min {
		return h.list[i].Priority > h.list[j].Priority
	} else {
		return h.list[i].Priority < h.list[j].Priority
	}
}

func (h *PriorityQueue) Items() (it types.Iterator) {
	i := 0
	return func() (item interface{}, next types.Iterator) {
		var e PqEntry
		if i < len(h.list) {
			e = h.list[i]
			i++
			return e.Item, it
		}
		return nil, nil
	}
}

// Verify the heap is properly structured.
func (h *PriorityQueue) Verify() error {
	for i := 1; i < len(h.list); i++ {
		parent := (i+1)/2 - 1
		if h.lt(parent, i) {
			return errors.Errorf("parent %v '<' kid %v", h.list[parent], h.list[i])
		}
	}
	return nil
}

func HeapSort(arr []int) {
	pqMin := NewPQ(len(arr), true)
	for _, ele := range arr {
		entry := PqEntry{Priority: ele, Item: ele}
		pqMin.Push(entry)
	}
	for i := 0; i < len(arr); i++ {
		entry := pqMin.Pop().Priority
		fmt.Printf("%v ,", entry)
	}
}

/*
Approach 1 : Frequency map => Sort By Value => then pick top K elements.
Approach 2 : using Priority queue.
*/
func KMostOccurance(a []int, k int) []int {
	freqMap := make(map[int]int)
	for _, v := range a {
		_, present := freqMap[v]
		if present {
			freqMap[v]++
		} else {
			freqMap[v] = 1
		}
	}
	pq := NewPQ(10, true)
	for key, preq := range freqMap {
		entry := PqEntry{Priority: preq, Item: key}
		if pq.Size() < k {
			pq.Push(entry)
		} else {
			peekValue := types.Int32(pq.Peek().Priority)
			if peekValue < types.Int32(preq) {
				pq.Pop()
				pq.Push(entry)
			}
		}
	}
	/*
		sort.Ints(keys)
		sort.SliceStable(keys, func(i, j int) bool {
			return freqMap[keys[i]] > freqMap[keys[j]]
		})
	*/
	result := make([]int, 0)
	for pq.Size() > 0 {
		item := pq.Pop().Item
		result = append(result, item.(int))
	}
	return result
}

func KthLargestSumSubArray(a []int, k int) int {
	prefixSum := make([]int, len(a))
	prefixSum[0] = a[0]
	for i := 1; i < len(a); i++ {
		prefixSum[i] = prefixSum[i-1] + a[i]
	}
	pq := NewPQ(10, true)
	for i := len(a) - 1; i >= 0; i-- {
		for j := i - 1; j >= -1; j-- {
			var sum int
			if j == -1 {
				sum = prefixSum[i]
			} else {
				sum = prefixSum[i] - prefixSum[j]
			}
			entry := PqEntry{Priority: sum, Item: sum}
			if pq.Size() < k {
				pq.Push(entry)
			} else {
				if pq.Peek().Priority < sum {
					pq.Pop()
					pq.Push(entry)
				}
			}
		}
	}
	return pq.Pop().Priority
}
