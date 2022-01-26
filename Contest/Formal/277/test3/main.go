package main

func findLonely(nums []int) (res []int) {
	lonelyMap := map[int]bool{}
	for _, n := range nums {
		_, preOk := lonelyMap[n-1]
		_, curOk := lonelyMap[n]
		_, nextOk := lonelyMap[n+1]
		if curOk {
			lonelyMap[n] = false
		}
		if preOk {
			lonelyMap[n-1] = false
			lonelyMap[n] = false
		}
		if nextOk {
			lonelyMap[n+1] = false
			lonelyMap[n] = false
		}
		if isLonely, ok := lonelyMap[n]; ok && !isLonely {
			continue
		}
		lonelyMap[n] = true
	}

	for k, v := range lonelyMap {
		if v {
			res = append(res, k)
		}
	}
	return
}
