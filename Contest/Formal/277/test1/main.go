package main

import "sort"

func countElements(nums []int) int {
	sort.Ints(nums)
	first, last := nums[0], nums[len(nums)-1]
	cnt := 0
	for _, v := range nums {
		if v == first || v == last {
			cnt++
		}
	}
	return len(nums) - cnt
}
