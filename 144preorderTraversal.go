package main

import (
	"fmt"
	"test/model"
)

func main() {
	node144 := &model.TreeNode{
		Val: 1,
		Left: &model.TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &model.TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Printf("%v", preorderTraversal(node144))
}
func preorderTraversal(root *model.TreeNode) []int {
	var preorder func(*model.TreeNode)
	var values []int
	preorder = func(node *model.TreeNode) {
		if node == nil {
			return
		}
		values = append(values, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return values
}
