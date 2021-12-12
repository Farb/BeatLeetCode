package main

// 选择排序：每次从未排序的数组中，选择最小的数字，与第一个元素进行交换。不稳定排序。
// 需要n-1轮遍历，每轮从未排序的数组中选择最小的数字与第一个元素交换。
// 第1轮，需从n个元素中选择最小的元素与之交换；第2轮，需要从n-1个元素中选择最小元素；
// 第n-1轮，需从n-(n-1-1)=2个元素中选择最小元素
// 第n轮，需从n-(n-1)=1个元素中选择最小元素（因为只有一个元素了，肯定是最大的，所以这步可省略）
func selectSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}
	return nums
}
