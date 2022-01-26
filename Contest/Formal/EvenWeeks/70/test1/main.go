package main

import "sort"

func minimumCost(cost []int) (res int) {
	sort.Ints(cost)
	for len(cost)>=3 {
		n:=len(cost)
		low,high:=cost[n-2],cost[n-1]
		res+=low+high
		cost=cost[:n-3]
	}
	res+=getLessEqualTwoCost(cost)
	return
}

func getLessEqualTwoCost(cost []int) (res int)  {
	for _, v := range cost {
		res += v
	}
	return
}