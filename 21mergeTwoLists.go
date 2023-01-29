package main

import "fmt"

func main() {

	list1 := &ListNode21{
		Val: 1,
		Next: &ListNode21{
			Val: 2,
			Next: &ListNode21{
				Val:  4,
				Next: nil,
			},
		},
	}

	list2 := &ListNode21{
		Val: 1,
		Next: &ListNode21{
			Val: 3,
			Next: &ListNode21{
				Val:  4,
				Next: nil,
			},
		},
	}
	fmt.Printf("%+v", mergeTwoLists(list1, list2))
}
func mergeTwoLists(list1, list2 *ListNode21) *ListNode21 {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

type ListNode21 struct {
	Val  int
	Next *ListNode21
}
