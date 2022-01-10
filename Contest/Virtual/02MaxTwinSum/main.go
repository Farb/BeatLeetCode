package main

//https://leetcode-cn.com/contest/biweekly-contest-69/problems/maximum-twin-sum-of-a-linked-list/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	arr := make([]int, 0, 16)

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	max := 0
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i]+arr[j] > max {
			max = arr[i] + arr[j]
		}
	}
	return max
}
