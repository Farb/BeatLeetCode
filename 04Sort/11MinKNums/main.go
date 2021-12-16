package main

// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
// 快速排序思路：如果k==哨兵的索引-1,即认为前k个小数已过滤出来，不需要继续排序了
func getLeastNumbers(arr []int, k int) []int {
	quickSort(arr, 0, len(arr)-1, k-1)
	return arr[:k]
}

func quickSort(nums []int, l, r, k int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r)
	if p == k {
		return
	}
	if p < k { //如果哨兵小于k，则需要继续分割右边的子数组
		quickSort(nums, p+1, r, k)
	} else {
		quickSort(nums, l, p-1, k)//如果哨兵>=k，则只需要继续分割左边的子数组
	}
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
