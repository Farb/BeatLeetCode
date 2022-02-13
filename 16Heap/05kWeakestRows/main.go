package main

import "container/heap"

//https://leetcode-cn.com/problems/the-k-weakest-rows-in-a-matrix/
func kWeakestRows(mat [][]int, k int) []int {
	h := new(hp)
	for row, arr := range mat {
		rowForce := 0
		for _, force := range arr {
			if force == 0 {
				break
			}
			rowForce += force
		}
		heap.Push(h, &ForcePair{force: rowForce, rowIndex: row})
		if h.Len() > k { // 弹出堆顶的最大元素，保证堆中有k个最小元素
			_ = heap.Pop(h).(*ForcePair)
		}
	}
	res := make([]int, h.Len())
	for i := h.Len() - 1; i >= 0; i-- { // 因为是大顶堆，所以弹出的首先是最大的元素，因而放到答案数组后面
		item := heap.Pop(h).(*ForcePair)
		res[i] = item.rowIndex
	}
	return res
}

type ForcePair struct {
	force    int
	rowIndex int
}
type hp []*ForcePair

func (a hp) Len() int      { return len(a) }
func (a hp) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a hp) Less(i, j int) bool { // 大顶堆排序规则
	if a[i].force > a[j].force {
		return true
	}
	if a[i].force == a[j].force {
		return a[i].rowIndex > a[j].rowIndex
	}
	return false
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(*ForcePair))
}
func (h *hp) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
