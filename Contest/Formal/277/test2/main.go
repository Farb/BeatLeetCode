package main

func rearrangeArray(nums []int) (res []int) {
	n := len(nums)
	pos, neg := 0, 0
	isPosTurn := true
	for len(res) < n {
		if isPosTurn {
			if nums[pos] > 0 {
				res = append(res, nums[pos])
				isPosTurn = !isPosTurn
			}
			pos++
		} else {
			if nums[neg] < 0 {
				res = append(res, nums[neg])
				isPosTurn = !isPosTurn
			}
			neg++
		}
	}
	return
}
