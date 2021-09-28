package main

import "fmt"

//时间和空间复杂度均为 O(n)，思路是遍历每个数组元素x，将target-x=y 的差值存入字典，
// 再遍历之后的元素，如果发现元素已存在于字典中，则找到了配对
func twoSum(nums []int, target int) []int {
	m := make(map[int]int,len(nums))
	for i := 0; i < len(nums); i++ {
		if value, exist := m[nums[i]]; exist {
			return []int{value, i}
		}
		m[target-nums[i]] = i
	}
	return []int{0, 0}
}
func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))

	nums = []int{3, 3}
	target = 6
	fmt.Println(twoSum(nums, target))

	nums = []int{3, 3}
	target = 5
	fmt.Println(twoSum(nums, target))
}
