package main

//https://leetcode-cn.com/problems/concatenation-of-array/

func getConcatenation(nums []int) []int {
	nums = append(nums, nums...)
	return nums
}