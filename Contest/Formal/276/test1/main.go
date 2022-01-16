package main

//https://leetcode-cn.com/contest/weekly-contest-276/problems/divide-a-string-into-groups-of-size-k/
func divideString(s string, k int, fill byte) []string {
	arr := []byte(s)
	res := make([]string, 0, len(arr)+k)
	for len(arr) >= k {
		res = append(res, string(arr[:k]))
		arr = arr[k:]
	}
	lastLen := len(arr)
	if lastLen <= 0 {
		return res
	}
	for i := 0; i < k-lastLen; i++ {
		arr = append(arr, fill)
	}
	res = append(res, string(arr))
	return res
}
