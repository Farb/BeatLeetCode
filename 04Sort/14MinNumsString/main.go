package main

import (
	"fmt"
	"strings"
)

//https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/
// 思路：快速排序
// 字符串比较规则：如果x+y>y+x,说明xy的字典序大于yx，所以yx更小
func minNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = fmt.Sprint(num)
	}
	quickSort(strs, 0, len(strs)-1)
	return strings.Join(strs, "")
}

func quickSort(strs []string, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	for i < j {
		//第一个元素是基准元素，如果基准元素+右侧元素<=右侧元素+基准元素,说明右侧元素大，基准元素字典序更小，符合从小到大的要求，不用交换顺序，故右指针自减
		for i < j && strings.Compare(strs[left]+strs[j], strs[j]+strs[left]) <= 0 {
			j--
		}
		//此时，右指针指向的元素字典序小于基准元素

		// 如果基准元素+左侧元素>=左侧元素+基准元素,说明左侧元素更小，基准元素字典序更大，符合从小到大的要求，不用交换顺序，故左指针自增
		for i < j && strings.Compare(strs[left]+strs[i], strs[i]+strs[left]) >=0 {
			i++
		}
		//此时，左指针指向的元素字典序比基准元素更大

		//交换左右指针指向的元素，即字典序大的元素靠后移动，小的靠前移动
		strs[i], strs[j] = strs[j], strs[i]
	}
	// 将基准元素放到合适的位置上
	strs[left], strs[i] = strs[i], strs[left]
	quickSort(strs, left, i-1)
	quickSort(strs, i+1, right)
}
