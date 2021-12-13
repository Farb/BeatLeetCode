package main

// https://leetcode-cn.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	return process(nums, target, 0, len(nums))
}

func process(nums []int, target, left, right int) int {
	if right == left {
		return right
	}
	mid := left + (right-left)>>1
	if nums[mid] == target {
		return mid
	}
	if target < nums[mid] {
		return process(nums, target, left, mid)
	}
	return process(nums, target, mid+1, right)
}
