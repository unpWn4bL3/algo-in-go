package queue

import "fmt"

// 基于最大二叉堆实现的优先队列

type PriorityQueue struct {
	Heap []int //最大二叉堆 0不用，这样left为2i，right为2i+1，parent为i/2
	N    int
}

func NewPriorityQueue(nums []int) *PriorityQueue {
	pq := &PriorityQueue{
		Heap: make([]int, 1),
		N:    0,
	}
	for _, v := range nums {
		pq.Insert(v)
	}
	return pq
}

func (q *PriorityQueue) Max() int {
	return q.Heap[1]
}

func (q *PriorityQueue) Min() int {
	return q.Heap[q.N]
}

// 上浮
func (q *PriorityQueue) Swim(i int) {
	for i > 1 && q.Heap[i] > q.Heap[i/2] {
		q.Swap(i, i/2)
		i = i / 2
	}
}

// 下沉
func (q *PriorityQueue) Sink(i int) {
	for i <= q.N {
		// 与较大的子节点交换
		var bigger int
		if q.Heap[2*i] > q.Heap[2*i+1] {
			bigger = 2 * i
		} else {
			bigger = 2*i + 1
		}
		if q.Heap[i] < q.Heap[bigger] {
			q.Swap(i, bigger)
			i = bigger
		} else {
			break
		}
	}
}

// 将新的值插入heap尾部并swim至堆平衡
func (q *PriorityQueue) Insert(e int) {
	q.N++
	q.Heap = append(q.Heap, e)
	q.Swim(q.N)
}

// 将heap头尾交换，删除尾部，然后sink头部至堆平衡
func (q *PriorityQueue) Del(i int) {
	q.Swap(1, q.N)
	q.Heap = q.Heap[0:q.N]
	q.N--
	q.Sink(1)
}

func (q *PriorityQueue) Swap(i, j int) {
	q.Heap[i], q.Heap[j] = q.Heap[j], q.Heap[i]
}

func main() {
	// nums := []int{1, 3, 5, 7, 9, 2, 4, 6, 8}
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pq := NewPriorityQueue(nums)
	fmt.Println(pq.Heap)
}
