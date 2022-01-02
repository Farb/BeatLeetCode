package main

// https://leetcode-cn.com/problems/merge-sorted-array/
//思路：从后向前遍历数组nums1和nums2,谁大谁放到nums1的末尾，结束条件是nums2中的元素都放入nums1，
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j ,k:= m-1, n-1,m+n-1
	for  j>=0 {
		if i < 0 || nums1[i] <= nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
		k--
	}
}

//思路：两个指针分别指向两个数组的尾部,最后拷贝nums2数组剩余的部分到nums1
func merge_2(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, k := m-1, n-1, m+n-1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[k] = nums1[p1]
			p1--
		} else {
			nums1[k] = nums2[p2]
			p2--
		}
		k--
	}
	// 跳出for循环后，如果p2还没走到0，说明nums2数组还没合并完成
	for i := 0; i <= p2; i++ {
		nums1[i]=nums2[i]
	}
}
