package main

import "math/rand"

//https://leetcode-cn.com/problems/linked-list-random-node/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}

//蓄水池抽样
func (this *Solution) GetRandom() int {
	i := 1
	reserve := 0
	for node := this.head; node != nil; node = node.Next {
		if 0 == rand.Intn(i) {
			reserve = node.Val
		}
		i++
	}
	return reserve
}

/* 空间复杂度O(n)
type Solution struct {
	arr []int
}

func Constructor(head *ListNode) Solution {
	var arr []int
	for head!=nil {
		arr = append(arr, head.Val)
		head=head.Next
	}
	return Solution{arr: arr}
}

func (this *Solution) GetRandom() int {
	size:=len(this.arr)
	randIdx:= rand.Intn(size)
	return this.arr[randIdx]
}
*/
/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
