package main

// 归并排序：分治法，先分（分割）再治（合并），每次从中点分割，当仅剩一个元素时，结束
// 时间复杂度：O(NlogN),需要开辟额外辅助空间O(N)，分割数组需要logN的递归栈
func mergeSort(nums []int) []int {
	if len(nums)<=1 {
		return nums
	}
	mid:=len(nums)/2
	left:=nums[:mid]
	right:=nums[mid:]
	return merge(mergeSort(left),mergeSort(right))
}

func merge(left,right []int) []int  {
	tmp:=make([]int,0,len(left)+len(right))
	for len(left)>0&&len(right)>0 {
		l,r:=left[0],right[0]
		if l<=r {
			tmp = append(tmp, l)
			left=left[1:]
		}else{
			tmp=append(tmp, r)
			right=right[1:]
		}
	}
	if len(left)>0 {
		tmp = append(tmp, left...)
	}
	if len(right)>0 {
		tmp = append(tmp, right...)
	}
	return tmp
}