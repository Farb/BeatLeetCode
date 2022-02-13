package main

import "container/heap"

// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
func findKthLargest(nums []int, k int) int {
	h := new(hp)
	for _, v := range nums {
		heap.Push(h, v)
		if h.Len() > k { // 保证小根堆中有k+1个元素
			_ = heap.Pop(h).(int) // 移除根元素（最小的元素），还剩k个元素，其他元素都是最大的元素，则根元素就是第k大的元素
		}
	}
	return heap.Pop(h).(int)
}

type hp []int

func (a hp) Len() int           { return len(a) }
func (a hp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a hp) Less(i, j int) bool { return a[i] < a[j] } // 小顶堆
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func (h *hp) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
