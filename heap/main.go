package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Heap struct {
	IsMin bool
	Arr   []int
	K     int
}

func NewHeap(isMin bool, k int) *Heap {
	return &Heap{
		IsMin: isMin,
		Arr:   []int{},
		K:     k,
	}
}

func (h *Heap) String() string {
	res := ""
	if h.IsMin == true {
		res += "{Type: Min, "
	} else {
		res += "{Type: Max, "
	}

	res += fmt.Sprintf("K: %v, Arr: %v}", h.K, h.Arr)

	return res
}

func (h *Heap) Insert(newValue int) {
	h.Arr = append(h.Arr, newValue)

	curIdx := len(h.Arr) - 1
	curVal := h.Arr[curIdx]
	parentIdx := (curIdx - curIdx%h.K) / h.K
	parentVal := h.Arr[parentIdx]

	for curVal != parentVal && ((curVal < parentVal) == h.IsMin) {
		h.Arr[curIdx], h.Arr[parentIdx] = h.Arr[parentIdx], h.Arr[curIdx]
		curIdx = parentIdx
		curVal = h.Arr[curIdx]
		parentIdx = (curIdx - curIdx%h.K) / h.K
		parentVal = h.Arr[parentIdx]
	}
}

func (h *Heap) Pop() int {
	res := h.Arr[0]

	h.Arr[0] = h.Arr[len(h.Arr)-1]
	h.Arr = h.Arr[:len(h.Arr)-1]

	h.Heapify()

	return res
}

func (h *Heap) heapify(curIdx int) {
	next := curIdx

	for i := 1; i <= h.K; i++ {
		nextIdx := curIdx*h.K + i
		if nextIdx < len(h.Arr) && ((h.Arr[next] > h.Arr[nextIdx]) == h.IsMin) {
			next = nextIdx
		}
	}

	if next != curIdx {
		h.Arr[curIdx], h.Arr[next] = h.Arr[next], h.Arr[curIdx]
		h.heapify(next)
	}
}

func (h *Heap) Heapify() {
	h.heapify(0)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	minHeap := NewHeap(true, 2)

	for i := 0; i < 30; i++ {
		r := rand.Intn(42)
		minHeap.Insert(r)
	}

	fmt.Println(minHeap)
	for len(minHeap.Arr) > 0 {
		fmt.Println(minHeap.Pop())
	}

	fmt.Println(minHeap)
}
