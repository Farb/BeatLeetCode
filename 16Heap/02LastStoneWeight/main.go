package main

import "container/heap"

// https://leetcode-cn.com/problems/last-stone-weight/

type IntHeap []int

func (a IntHeap) Len() int           { return len(a) }
func (a IntHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntHeap) Less(i, j int) bool { return a[i] > a[j] } // 大根堆
func (a *IntHeap) Push(v interface{}) {
	*a = append(*a, v.(int))
}
func (a *IntHeap) Pop() (v interface{}) {
	*a, v = (*a)[:a.Len()-1], (*a)[a.Len()-1]
	return v
}

// 大根堆思路：每次从堆中取最重的2个石头（如果有），取差值，如果差值大于0，将撞碎的小石头放入优先队列，继续循环
func lastStoneWeight(stones []int) int {
	n := len(stones)
	h := make(IntHeap, n)
	for i, stone := range stones {
		h[i] = stone
	}
	heap.Init(&h)  // 初始化堆
	w1, w2 := 0, 0 // 每次取最重的两块石头
	for h.Len() > 1 {
		w1 = heap.Pop(&h).(int)
		w2 = heap.Pop(&h).(int)
		diff := w1 - w2 // 取差值
		if diff > 0 {
			heap.Push(&h, diff) // 如果不为0，则将小石头入队
		}
	}
	if h.Len() == 1 { // 如果最后只剩一块石头，则返回最后一块石头的重量
		return heap.Pop(&h).(int)
	}
	return 0 // 没剩下石头，返回0
}
