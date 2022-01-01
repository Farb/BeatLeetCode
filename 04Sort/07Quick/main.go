package main

// 快速排序：核心包括哨兵分割和递归哨兵两边的左右子数组
// 最差时间复杂度：O(n2)，平均时间复杂度O(NlogN)
// 快速排序是不稳定排序，体现在，比如，比哨兵小的右侧的相等的2个数，先是交换后面的一个数，再交换前一个数，导致顺序变了。
func quickSort(nums []int, l, r int) {
	if l >= r { //basecase:当左边界>=右边界时，说明只有一个元素，直接返回
		return
	}
	part := partition(nums, l, r) //哨兵分割
	quickSort(nums, l, part-1) //对哨兵左侧进行快排
	quickSort(nums, part+1, r) //对哨兵右侧快排
}

func partition(nums []int, l, r int) int {
	i, j := l, r //定义2个指针指向左右边界，并已左边界的元素作为哨兵
	for i < j {
		for i < j && nums[j] >= nums[l] { //先从右边界开始遍历，如果右指针j指向的元素>=哨兵，继续遍历
			j--
		}
		// 此时，右指针指向的值小于哨兵
		for i < j && nums[i] <= nums[l] { //再从左边界开始遍历，如果左指针i指向的元素<=哨兵，继续遍历
			i++
		}
		// 此时，左指针指向的值大于哨兵
		nums[i], nums[j] = nums[j], nums[i] //交换左右指针的值，使得较小值在左侧，较大值在右侧
	}
	nums[i], nums[l] = nums[l], nums[i]
	return i
}
