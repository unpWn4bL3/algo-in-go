package main

import (
	"algo/sort"
	"fmt"
)

func main() {
	nums := []int{1, 2, 0}
	sort.InsertionSort(nums)
	fmt.Println(nums)
}
