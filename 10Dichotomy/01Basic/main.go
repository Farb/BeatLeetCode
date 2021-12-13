package main

func search(nums []int, target int) int {
	return process(nums, target, 0, len(nums)-1)
}

func process(nums []int, target, left, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)>>1
	if nums[mid] == target {
		return mid
	}
	if nums[mid] < target {
		return process(nums, target, mid+1, right)
	}
	return process(nums, target, left, mid-1)
}
