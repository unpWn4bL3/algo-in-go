package main

import (
	"container/heap"
	"fmt"
)

type element struct {
	v   int
	idx int
}
type PriorityQueue []*element

func main() {
	nums := []int{1, -1}
	r := maxSlidingWindow(nums, 1)
	fmt.Println(r)
}

func maxSlidingWindow(nums []int, k int) []int {
	pq := PriorityQueue([]*element{})
	heap.Init(&pq)
	for i := 0; i < k; i++ {
		heap.Push(&pq, &element{v: nums[i], idx: i})
	}
	r := []int{}
	r = append(r, pq[0].v)
	for i := k; i < len(nums); i++ {
		for len(pq) > 0 && pq[0].idx <= i-k {
			heap.Pop(&pq)
		}
		heap.Push(&pq, &element{v: nums[i], idx: i})
		r = append(r, pq[0].v)
	}
	return r
}

// Len is the number of elements in the collection.
func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i int, j int) bool {
	return pq[i].v > pq[j].v // 大顶堆
}

// Swap swaps the elements with indexes i and j.
func (pq PriorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*element)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	r := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return r
}
