package main

import (
	"algo/tree"
	"fmt"
)

func main() {
	var root *tree.Node
	nums := []string{"1", "2", "3", "4", "5", "#", "6", "#", "#", "7", "8"}
	root = tree.InitBinaryTree(nums)
	foo(root)
}

func foo(root *tree.Node) {
	if root == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(root.Val)
	foo(root.Left)
	foo(root.Right)
}
