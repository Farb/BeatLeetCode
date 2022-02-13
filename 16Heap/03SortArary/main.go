package main

import "container/heap"

// https://leetcode-cn.com/problems/sort-an-array/

func sortArray(nums []int) []int {
	n := len(nums)
	sortedNums := new(hp)
	for _, num := range nums {
		heap.Push(sortedNums, num)
	}
	i := n - 1
	for sortedNums.Len() > 0 {
		nums[i] = heap.Pop(sortedNums).(int)
		i--
	}
	return nums
}

type hp []int

func (a hp) Len() int           { return len(a) }
func (a hp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a hp) Less(i, j int) bool { return a[i] > a[j] } // 大顶堆
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *hp) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
