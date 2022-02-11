package main

import "container/heap"

//https://leetcode-cn.com/problems/merge-k-sorted-lists/solution/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 3- k个指针比较，每次取最小值
func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	dummyNode := new(ListNode)
	tail := dummyNode
	for {
		var minNode *ListNode // 每次寻找遍历K条链表时，必须重新初始化最小节点信息，否则可能受上次结果影响
		minIndex := -1
		for i := 0; i < k; i++ { // 遍历K条链表，找出最小的节点
			if lists[i] == nil { // 遍历某个链表时，如果为nil，说明链表已经合并完成
				continue
			}
			if minNode == nil || lists[i].Val < minNode.Val {
				minNode = lists[i]
				minIndex = i
			}
		}
		if minIndex == -1 { // 没找到最小节点
			break
		}
		tail.Next = minNode
		tail = tail.Next
		lists[minIndex] = lists[minIndex].Next // 最小值的那个链表删除最小节点
	}
	return dummyNode.Next
}

type SortedList []*ListNode

func (a SortedList) Len() int           { return len(a) }
func (a SortedList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortedList) Less(i, j int) bool { return a[i].Val < a[j].Val }
func (a *SortedList) Push(v interface{}) {
	*a = append(*a, v.(*ListNode))
}
func (a *SortedList) Pop() (v interface{}) {
	*a, v = (*a)[:a.Len()-1], (*a)[a.Len()-1]
	return
}

// 2-小根堆合并：每次取所有链表的头节点的最小值，依次合并，直到优先队列中没有节点
func mergeKLists_minHeap(lists []*ListNode) *ListNode {
	sortedList := new(SortedList)
	for _, list := range lists {
		if list != nil { // 判空处理，list可能是空链表
			heap.Push(sortedList, list) // 将所有链表的头节点加入链表
		}
	}
	dummyNode := &ListNode{}
	cur := dummyNode
	for sortedList.Len() > 0 {
		node := heap.Pop(sortedList).(*ListNode)
		cur.Next = node
		cur = cur.Next
		if node.Next != nil { // 当链表的节点后面没有下一个节点时，说明该节点所在的链表合并完成；否则添加下一个节点（未完成合并的头节点）到队列中
			heap.Push(sortedList, node.Next)
		}
	}
	return dummyNode.Next
}

// 1-分治思想：两两合并
func mergeKLists_dc(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 { // 如果没有链表，返回nil
		return nil
	}
	if n == 1 {
		return lists[0] // 如果只有一个链表，直接返回
	}
	mid := n >> 1
	lowList := mergeKLists(lists[:mid])     // 自上向下看，递归合并0到中间区间的链表；自下向上看，lowList=lists[0]
	highList := mergeKLists(lists[mid:])    // 自上向下看，递归合并中间到末尾的链表；自下向上看，highList=lists[1]
	return mergeTwoLists(lowList, highList) //自上向下看，合并两半链表组成的长链表；自下向上看，合并lists[0]，lists[1]
}

// 合并两个链表的递归写法，递归栈帧较占内存，空间复杂度O(logN)
// func mergeTwoLists(a, b *ListNode) *ListNode {
// 	if a == nil {
// 		return b
// 	}
// 	if b == nil {
// 		return a
// 	}
// 	if a.Val < b.Val {
// 		a.Next = mergeTwoLists(a.Next, b)
// 		return a
// 	}
// 	b.Next = mergeTwoLists(a, b.Next)
// 	return b
// }

// 合并两个链表的迭代写法
func mergeTwoLists(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	dummyNode := &ListNode{}
	cur := dummyNode
	for a != nil && b != nil {
		if a.Val < b.Val {
			cur.Next = a
			a = a.Next
		} else {
			cur.Next = b
			b = b.Next
		}
		cur = cur.Next
	}
	if a != nil {
		cur.Next = a
	} else {
		cur.Next = b
	}
	return dummyNode.Next
}
