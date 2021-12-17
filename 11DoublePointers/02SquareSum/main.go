package main

import "math"

// https://leetcode-cn.com/problems/sum-of-square-numbers/

func judgeSquareSum(c int) bool {
	i, j := 0, int(math.Sqrt(float64(c)))
	if j*j == c {
		return true
	}
	for i <= j {
		sum := i*i + j*j
		if sum == c {
			return true
		}
		if sum > c {
			j--
		} else {
			i++
		}
	}
	return false
}
