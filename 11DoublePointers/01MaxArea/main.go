package main

// https://leetcode-cn.com/problems/container-with-most-water/
func maxArea(height []int) int {
	max := 0
	for i, j := 0, len(height)-1; i < j; {
		area := area(height, i, j)
		if area > max {
			max = area
		}
		if height[i] >= height[j] {
			j--
		} else {
			i++
		}
	}
	return max
}
func area(height []int, i, j int) int {
	return (j - i) * min(height[i], height[j])
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
