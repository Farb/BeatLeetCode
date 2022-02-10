package main

//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		if pre == nil || pre.Val != cur.Val { // 当前节点的值不等于前一个节点的值
			pre = cur // 将当前节点赋值给pre
		} else {
			pre.Next = cur.Next // 当前节点的值等于前一个节点的值，说明当前节点重复，故删除当前节点，将pre指向当前节点的下一个
		}
		cur = cur.Next
	}
	return head
}
