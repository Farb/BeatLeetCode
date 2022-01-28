package main

import (
	"sort"
)

// https://leetcode-cn.com/problems/the-number-of-weak-characters-in-the-game/

func numberOfWeakCharacters(props [][]int) (res int) {
	// 对由攻击值和防御值组成的二维数组进行排序，排序规则时优先按攻击值降序排序，攻击者相等时，防御值按升序排序
	sort.Slice(props, func(i, j int) bool {
		return props[i][0] > props[j][0] || props[i][0] == props[j][0] && props[i][1] < props[j][1]
	})

	maxDefense := 0 //1.因为攻击值降序排序，防御值升序排序，所以攻击值相等时，只有防御值是不能确定弱角色的（题意说攻击值和防御值都小于某个角色）
	for _, ps := range props { //2.每次遍历的攻击值不相等时，每次的遍历攻击值必然小于前次的攻击值，如果防御值也小于前面的最大防御值时，那么本次遍历的角色就是弱角色
		if ps[1] < maxDefense { //综上2点，只有当每次遍历的角色的防御值小于之前遍历的最大防御值时，结果+1
			res++
		} else {
			maxDefense = ps[1] // 最大防御值更新
		}
	}
	return
}
