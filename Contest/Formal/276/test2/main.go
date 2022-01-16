package main

//https://leetcode-cn.com/contest/weekly-contest-276/problems/minimum-moves-to-reach-target-score/
func minMoves(target int, maxDoubles int) (res int) {
	for target > 1 {
		if maxDoubles <= 0 { //不能加倍的话，直接算出递增步数
			res += target - 1
			break
		}
		res++
		if target&1 == 1 { // 奇数
			target -= 1
			continue
		}
		maxDoubles--//偶数且可以加倍
		target >>= 1
	}
	return
}
