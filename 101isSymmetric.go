package main

import "test/model"

func main() {

}

func isSymmetric(root *model.TreeNode) bool {
	return check(root, root)
}
func check(left, right *model.TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && check(left.Left, right.Right) && check(left.Right, right.Left)
}
