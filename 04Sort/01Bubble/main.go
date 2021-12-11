package main

// 冒泡排序：n个数字需要n-1轮遍历，第一轮两两数字比较，如果前者比后者大，交换顺序，
// 一直遍历到第n-1-1个元素；第2轮遍历到第n-2-1个元素（因为第n-1个已经排好序），第n-1轮遍历到第1(n-1-(n-1)=0)个元素
func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
