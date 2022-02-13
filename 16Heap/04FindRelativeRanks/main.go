package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode-cn.com/problems/relative-ranks/
func findRelativeRanks(score []int) []string {
	n := len(score)
	sorted := new(hp)
	for i, v := range score {
		heap.Push(sorted, &ScorePair{v, i})
	}
	// 逐一取出分数最大的，从第一名依次取
	ans := make([]string, n)
	for i := 1; i <= n; i++ {
		item := heap.Pop(sorted).(*ScorePair)
		rank := ""
		switch i {
		case 1:
			rank = "Gold Medal"
		case 2:
			rank = "Silver Medal"
		case 3:
			rank = "Bronze Medal"
		default:
			rank = fmt.Sprintf("%d", i)
		}
		ans[item.index] = rank
	}
	return ans
}

type ScorePair struct {
	score int
	index int
}
type hp []*ScorePair

func (a hp) Len() int           { return len(a) }
func (a hp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a hp) Less(i, j int) bool { return a[i].score > a[j].score } // 大根堆
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(*ScorePair))
}

func (h *hp) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
