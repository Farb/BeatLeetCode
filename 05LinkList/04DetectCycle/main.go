package main

//https://leetcode-cn.com/leetbook/read/linked-list/jjhf6/


//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

 func detectCycle(head *ListNode) *ListNode {
    fast,slow:=head,head
	for fast!=nil&&fast.Next!=nil {
		fast=fast.Next.Next
		slow=slow.Next
		if fast==slow {
			break
		}
	}
	if fast==nil||fast.Next==nil { //判断是否有环
		return nil
	}

	fast=head
	for fast!=slow {
		fast=fast.Next
		slow=slow.Next
	}
	return fast
}