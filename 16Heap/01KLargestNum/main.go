package main

import "container/heap"

// https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/
type IntHeap []int

func (a IntHeap) Len() int           { return len(a) }
func (a IntHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntHeap) Less(i, j int) bool { return a[i] < a[j] } // 小顶堆
func (a *IntHeap) Push(v interface{}) {
	*a = append(*a, v.(int))
}
func (a *IntHeap) Pop() (v interface{}) {
	*a, v = (*a)[:a.Len()-1], (*a)[a.Len()-1]
	return v
}

type KthLargest struct {
	sortedNums *IntHeap
	k          int
}

func Constructor(k int, nums []int) KthLargest {
	sortNums := make(IntHeap, len(nums))
	for i, num := range nums {
		sortNums[i] = num
	}
	heap.Init(&sortNums)       // 初始化堆
	for sortNums.Len()-k > 0 { // 确保堆中只有k个元素，这样，根元素就是第k大元素
		heap.Pop(&sortNums) // 其他元素出队列
	}
	return KthLargest{k: k, sortedNums: &sortNums}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.sortedNums, val)        // 入队
	for this.sortedNums.Len()-this.k > 0 { // 确保堆中只有k个元素
		heap.Pop(this.sortedNums)
	}
	v := heap.Pop(this.sortedNums).(int) // 弹出堆顶元素，即答案
	heap.Push(this.sortedNums, v)        // 再加入堆
	return v
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
