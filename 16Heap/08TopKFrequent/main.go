package main

import "container/heap"

// https://leetcode-cn.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) (res []int) {
	cntMap := map[int]int{}
	for _, v := range nums { // 统计每个元素出现的次数
		cntMap[v]++
	}
	h := new(hp)
	for num, cnt := range cntMap { // 将每个元素根据出现次数排序，放入小根堆
		heap.Push(h, &Pair{Num: num, Cnt: cnt})
		if h.Len() > k { // 如果超过k个元素，移除根元素，保证最多有k个元素，根元素为第k大元素
			_ = heap.Pop(h)
		}
	}
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(*Pair).Num)
	}
	return
}

type Pair struct {
	Num int
	Cnt int
}

type hp []*Pair

func (a hp) Len() int           { return len(a) }
func (a hp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a hp) Less(i, j int) bool { return a[i].Cnt < a[j].Cnt } // 小根堆
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(*Pair))
}
func (h *hp) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
