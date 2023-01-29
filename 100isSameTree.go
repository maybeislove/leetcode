package main

import (
	"fmt"
	"test/model"
)

func main() {
	p := &model.TreeNode{
		Val: 1,
		Left: &model.TreeNode{
			Val: 2,
		},
		Right: nil,
	}
	q := &model.TreeNode{
		Val:  1,
		Left: nil,
		Right: &model.TreeNode{
			Val: 2,
		},
	}
	fmt.Println(isSameTree(p, q))
}

func isSameTree(p *model.TreeNode, q *model.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}
