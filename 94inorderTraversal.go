package main

import "test/model"

func main() {

}

func inorderTraversal(root *model.TreeNode) []int {
	res := []int{}
	inorder := func(node *model.TreeNode) {}
	inorder = func(node *model.TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
}
