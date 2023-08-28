package queue

import "fmt"

// 基于最大二叉堆实现的优先队列

type PriorityQueue struct {
	heap []int //最大二叉堆 0不用，这样left为2i，right为2i+1，parent为i/2
	n    int
}

func NewPriorityQueue(nums []int) *PriorityQueue {
	pq := &PriorityQueue{
		heap: make([]int, 1),
		n:    0,
	}
	for _, v := range nums {
		pq.Insert(v)
	}
	return pq
}

func (q *PriorityQueue) Max() int {
	return q.heap[1]
}

func (q *PriorityQueue) Min() int {
	return q.heap[q.n]
}

// 上浮
func (q *PriorityQueue) Swim(i int) {
	for i > 1 && q.heap[i] > q.heap[i/2] {
		q.Swap(i, i/2)
		i = i / 2
	}
}

// 下沉
func (q *PriorityQueue) Sink(i int) {
	for i <= q.n {
		if q.heap[i] < q.heap[2*i] {
			q.Swap(i, 2*i)
			i = 2 * i
		} else if q.heap[i] < q.heap[2*i+1] {
			q.Swap(i, 2*i+1)
			i = 2*i + 1
		} else {
			break
		}
	}
}

// 将新的值插入heap尾部并swim至堆平衡
func (q *PriorityQueue) Insert(e int) {
	q.n++
	q.heap = append(q.heap, e)
	q.Swim(q.n)
}

// 将heap头尾交换，删除尾部，然后sink头部至堆平衡
func (q *PriorityQueue) Del(i int) {
	q.Swap(1, q.n)
	q.heap = q.heap[0:q.n]
	q.n--
	q.Sink(1)
}

func (q *PriorityQueue) Swap(i, j int) {
	q.heap[i], q.heap[j] = q.heap[j], q.heap[i]
}

func main() {
	// nums := []int{1, 3, 5, 7, 9, 2, 4, 6, 8}
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pq := NewPriorityQueue(nums)
	fmt.Println(pq.heap)
}
