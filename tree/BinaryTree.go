package tree

import "strconv"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// 用 "#" 代表空节点
func InitBinaryTree(nums []string) *Node {
	var createTree func([]string, int) *Node
	createTree = func(nums []string, idx int) *Node {
		if len(nums) <= idx || nums[idx] == "#" {
			return nil
		}
		val, _ := strconv.Atoi(nums[idx])
		root := &Node{Val: val}
		l := 2*idx + 1
		r := 2*idx + 2
		root.Left = createTree(nums, l)
		root.Right = createTree(nums, r)
		return root

	}
	return createTree(nums, 0)
}
