package main

//https://leetcode-cn.com/problems/water-bottles/

func numWaterBottles(numBottles int, numExchange int) int {
	sum,left := numBottles,0
	for numBottles != 0 {
		numBottles,left = ( numBottles + left) / numExchange, ( numBottles + left) % numExchange
		sum += numBottles
	}
	return sum
}
