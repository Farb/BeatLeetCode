package main

// 希尔排序：平均时间复杂度O(NlogN),最差时间复杂度O(N^2)
// 核心思路：每次对数组按照不同间隔（N/2,N/4，...，1或1，3，5，...,2N-1）进行分组，对每个分组进行直接插入排序
// 不稳定排序，因为按照不同间隔进行分组，所以相同的两个连续数会被分到不同组，导致交换之后，顺序发生变化
func ShellSort(nums []int) {
	// 每次循环将数组分成不同的间隔小组
	for step := len(nums) / 2; step >= 1; step-- {
		// 遍历每个间隔小组的中的所有元素（除首元素外）
		for i := step; i < len(nums); i += step {
			// 遍历当前元素之前的所有元素形成的有序数组，寻找位置
			for j := i - step; j >= 0; j -= step {
				if nums[j] > nums[j+step] { //如果前面的数比后面的数大，则交换位置
					nums[j], nums[j+step] = nums[j+step], nums[j]
				}
			}
		}
	}
}
