package main

import (
	"fmt"
	"test/model"
)

func main() {
	node145 := &model.TreeNode{
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
	fmt.Printf("%v", postorderTraversal(node145))
}

func postorderTraversal(root *model.TreeNode) []int {
	var postorder func(*model.TreeNode)
	var values []int
	postorder = func(node *model.TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		values = append(values, node.Val)
	}
	postorder(root)
	return values
}
