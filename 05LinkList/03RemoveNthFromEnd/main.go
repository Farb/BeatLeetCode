package main

//https://leetcode-cn.com/leetbook/read/linked-list/jf1cc/
//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路：起始，快慢指针同时指向伪头节点，然后快指针先走n步，然后快慢指针同频走，直到快指针到达尾部,此时慢指针到达倒数第n+1个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Next: head}
	slow, fast := dummyNode, dummyNode
	for n > 0 {
		fast = fast.Next
		n--
	}

	for fast.Next != nil { //快慢指针同时走，直到快指针到达尾部
		fast = fast.Next
		slow = slow.Next
	}
	if slow.Next != nil { //这里要判空，否则1语法报错；2倒数第一个节点没有下一个节点
		slow.Next = slow.Next.Next
	} 
	return dummyNode.Next
}

// 栈思路：先将所有节点放入栈中，根据先进后出原则，弹栈，
func removeNthFromEnd_stack(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummyNode := &ListNode{0, head}
	for cur := dummyNode; cur.Next != nil; cur = cur.Next {
		nodes = append(nodes, cur)
	}
	preNode := nodes[len(nodes)-1-n]
	preNode.Next = preNode.Next.Next
	return dummyNode.Next
}
