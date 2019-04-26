package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Heap struct {
	isMin bool
	k     int
	arr   []float64
}

func New(isMin bool, k int) *Heap {
	return &Heap{
		isMin: isMin,
		k:     k,
		arr:   []float64{},
	}
}

func (h *Heap) String() string {
	return fmt.Sprintf("%#v", h)
}

func (h *Heap) Insert(newVal float64) {
	h.arr = append(h.arr, newVal)

	curIdx := len(h.arr) - 1
	parentIdx := (curIdx - h.k + 1) / h.k
	for h.arr[curIdx] != h.arr[parentIdx] && (h.arr[curIdx] < h.arr[parentIdx]) == h.isMin {

		h.arr[curIdx], h.arr[parentIdx] = h.arr[parentIdx], h.arr[curIdx]

		curIdx = parentIdx
		parentIdx = (curIdx - h.k + 1) / h.k
	}
}

func (h *Heap) Pop() float64 {
	res := h.arr[0]

	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]

	h.Heapify()

	return res
}

func (h *Heap) heapify(curIdx int) {

	targetIdx := curIdx

	for i := 1; i <= h.k; i++ {
		nextIdx := curIdx*h.k + i

		if nextIdx < len(h.arr) && ((h.arr[nextIdx] < h.arr[targetIdx]) == h.isMin) {
			targetIdx = nextIdx
		}
	}

	if targetIdx != curIdx {
		h.arr[curIdx], h.arr[targetIdx] = h.arr[targetIdx], h.arr[curIdx]
		h.heapify(targetIdx)
	}
}

func (h *Heap) Heapify() {
	h.heapify(0)
}

func (h *Heap) IsEmpty() bool {
	return len(h.arr) == 0
}

func main() {
	rand.Seed(time.Now().UnixNano())

	minHeap := New(false, 2)
	fmt.Println(minHeap)

	for i := 0; i < 100; i++ {
		minHeap.Insert(float64(rand.Intn(100)))
	}
	fmt.Println(minHeap)

	for minHeap.IsEmpty() == false {
		fmt.Printf("%v ", minHeap.Pop())
	}
}
