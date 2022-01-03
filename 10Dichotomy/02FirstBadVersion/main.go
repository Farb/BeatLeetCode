package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
//递归法
func firstBadVersion(n int) int {
	return process(1, n)
}

func process(start, end int) int {
	if start == end {
		return start
	}
	mid := start + (end-start)>>1
	if !isBadVersion(mid) {
		return process(mid+1, end)
	}
	return process(start, mid)
}
func isBadVersion(version int) bool {
	return false
}

// 循环迭代法模拟二分法
func firstBadVersion_iteration(n int) int {
	i, j, mid := 1, n, 0
	for i < j {
		mid = i + (j-i)>>1
		if isBadVersion(mid) {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return i
}
