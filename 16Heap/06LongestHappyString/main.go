package main

import (
	"container/heap"
)

func longestDiverseString(a int, b int, c int) string {
	queue := new(hp) // 初始化优先队列，以[字符，字符剩余数量] 存储 进行排序 使得字符剩余数量从大到小的（大根堆）排序
	if a > 0 {
		heap.Push(queue, &Pair{ch: 'a', count: a})
	}
	if b > 0 {
		heap.Push(queue, &Pair{'b', b})
	}
	if c > 0 {
		heap.Push(queue, &Pair{ch: 'c', count: c})
	}
	res := []rune{}
	for queue.Len() > 0 {
		curPair := heap.Pop(queue).(*Pair) // 取出当前字符剩余数量最大的元素
		n := len(res)
		if n >= 2 && res[n-2] == curPair.ch && res[n-1] == curPair.ch { // 如果长度大于等于2时，则需要对当前取出的剩余数量最大的元素与前两个元素是否一致且连续相同
			if queue.Len() <= 0 { // 如果是连续相同的 我们需要进行第一次判断当前队列中是否还有剩余数量的其他字符， 也就是说为了我们找到第二大剩余字符数量的字符
				break
			}
			nextPair := heap.Pop(queue).(*Pair) // 取出当前第二大的剩余数量的字符
			res = append(res, nextPair.ch)  // 将其追加到结果
			nextPair.count--                // 如果拼接字符的剩余数量不为0时咱们再把它添加到大根堆中再次排序
			if nextPair.count > 0 {
				heap.Push(queue, nextPair)
			}
		} else {
			res = append(res, curPair.ch) // 如果不是连续相同的咱们就把它 直接追加到结果
			curPair.count--               //当前数量自减
		}
		if curPair.count > 0 { // 如果当前字符的剩余数量>0,则加入队列继续下次排序
			heap.Push(queue, curPair)
		}
	}
	return string(res)
}

type Pair struct {
	ch    rune
	count int
}
type hp []*Pair

func (a hp) Len() int           { return len(a) }
func (a hp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a hp) Less(i, j int) bool { return a[i].count > a[j].count } // 大根堆，数量多的优先
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(*Pair))
}

func (h *hp) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
