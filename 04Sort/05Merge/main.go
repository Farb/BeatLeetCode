package main

// 归并排序：分治法，先分（分割）再治（合并），每次从中点分割，当仅剩一个元素时，结束
// 时间复杂度：O(NlogN),需要开辟额外辅助空间O(N)，分割数组需要logN的递归栈
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left, right := nums[:mid], nums[mid:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	tmp := make([]int, 0, len(left)+len(right)) //申请一个临时数组，长度为两个子数组的长度之和
	for len(left) > 0 && len(right) > 0 {       //合并两个子数组
		l, r := left[0], right[0]
		if l <= r {
			tmp = append(tmp, l)
			left = left[1:]
		} else {
			tmp = append(tmp, r)
			right = right[1:]
		}
	}
	if len(left) > 0 { //如果左子数组有元素，右子数组无元素，则直接将左子数组追加到临时数组
		tmp = append(tmp, left...)
	}
	if len(right) > 0 { //如果右子数组有元素，左子数组无元素，则直接将右子数组追加到临时数组
		tmp = append(tmp, right...)
	}
	return tmp
}
