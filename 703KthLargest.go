package main

import (
	"container/heap"
	"sort"
)

func main() {
	//nums := []int{4, 5, 8, 2}
	nums1 := []int{}

	k := Constructor(3, nums1)
	println(k.Add(3))
	println(k.Add(5))
	println(k.Add(10))
	println(k.Add(9))
	println(k.Add(4))
}

type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{
		k: k,
	}
	for _, num := range nums {
		kth.Add(num)
	}
	return kth
}

func (this *KthLargest) Push(val interface{}) {
	this.IntSlice = append(this.IntSlice, val.(int))
}

func (this *KthLargest) Pop() interface{} {
	slice := this.IntSlice
	i := slice[slice.Len()-1]
	this.IntSlice = slice[:len(slice)-1]
	return i
}
func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.k {
		heap.Pop(this)
	}
	return this.IntSlice[0]
}
