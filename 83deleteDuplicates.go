package main

import "fmt"

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	deleteDuplicates(node)
}
func deleteDuplicates(head *ListNode) *ListNode {
	fmt.Println(head.Val)
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = deleteDuplicates(head.Next)
	if head.Val == head.Next.Val {
		if head.Next.Next == nil {
			head.Next = nil
		} else {
			head = head.Next
		}
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
