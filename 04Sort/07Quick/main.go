package main

// 快速排序：核心包括哨兵分割和递归哨兵两边的左右子数组
// 最差时间复杂度：O(n2)，平均时间复杂度O(NlogN)
func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	part := partition(nums, l, r)
	quickSort(nums, l, part-1)
	quickSort(nums, part+1, r)
}

func partition(nums []int, l, r int) int {
	i, j := l, r
	for i < j {
		for i < j && nums[j] >= nums[l] {
			j--
		}

		for i < j && nums[i] <= nums[l] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[l] = nums[l], nums[i]
	return i
}
