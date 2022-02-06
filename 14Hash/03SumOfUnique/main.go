package main

// https://leetcode-cn.com/problems/sum-of-unique-elements/
//遍历一次：将每个元素出现的次数存入hash表，如果是第一次出现，则累加到答案，如果是第二次则从答案中减去，第2次以上不处理
func sumOfUnique_hash(nums []int) (res int) {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
		if m[v] == 1 {
			res += v
		} else if m[v] == 2 {
			res -= v
		}
	}
	return
}

// 状态机思想， 1 <= nums[i] <= 100；1 <= nums.length <= 100
func sumOfUnique(nums []int) (res int) {
	cnt := [101]int{} //因为数的范围是[1,100]，所以申请101个长度的数组（第0个不用）标记每个数出现的次数（下标对应元素，下标对应的值代表出现次数）
	states := [101]int{1, -1} //申请101个长度的数组（第0个不用）标记每个状态对应的策略，第一个是+n,第2个是-n,其余的都是+0
	for _, n := range nums {
		res += states[cnt[n]] * n //第1个状态是+n,第2个是-n,其余的都是+0
		cnt[n]++ //统计每个整数出现的次数
	}
	return
}
