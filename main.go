package main

import (
	linkedlist "algo/linked-list"
	"fmt"
	"math"
)

func main() {
	ll := make([]*linkedlist.ListNode, 0)
	ll = append(ll, linkedlist.NewLinkedList([]int{1, 4, 5}))
	ll = append(ll, linkedlist.NewLinkedList([]int{1, 3, 4}))
	ll = append(ll, linkedlist.NewLinkedList([]int{2, 6}))
	r := mergeKLists(ll)
	for cur := r; cur != nil; cur = cur.Next {
		fmt.Println(cur.Val)
	}

}

func mergeKLists(lists []*linkedlist.ListNode) *linkedlist.ListNode {
	dummy := &linkedlist.ListNode{}
	cur := dummy
	for {
		allEmpty := true
		min := math.MaxInt
		minIdx := -1
		for i := 0; i < len(lists); i++ {
			head := lists[i]
			if head != nil {
				allEmpty = false
				if head.Val < min {
					minIdx = i
					min = head.Val
					fmt.Println(i, minIdx, min)
				}
			}
		}
		if allEmpty {
			break
		}
		cur.Next = lists[minIdx]
		lists[minIdx] = lists[minIdx].Next
		cur = cur.Next
		cur.Next = nil
		for cur := dummy.Next; cur != nil; cur = cur.Next {
			fmt.Println(cur.Val)
		}
		fmt.Println("----")

	}
	return dummy.Next
}
