package main

import (
	"algo/queue"
	"fmt"
)

func main() {
	pq := queue.NewPriorityQueue([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(pq.Heap)
}
