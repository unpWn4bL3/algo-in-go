package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkedList(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range nums {
		cur.Next = &ListNode{Val: v, Next: nil}
		cur = cur.Next
	}
	return dummy.Next
}
