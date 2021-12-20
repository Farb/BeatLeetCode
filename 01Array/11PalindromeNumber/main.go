package main
//https://leetcode-cn.com/problems/palindrome-number/
// 解法一：借助额外数组
func isPalindrome_arr(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) { //排除负数和尾数有0的数，除0之外
		return false
	}
	arr := []byte{}
	for x != 0 {
		arr = append(arr, byte(x%10))
		x /= 10
	}
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}
	return true
}

// 解法二：反转数字的后一半，判断是否和前一半相等即可
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) { //排除负数和尾数有0的数，除0之外
		return false
	}
	revertedNum := 0 //反转的数每次*10,并加上x%10的数。到一半的判断逻辑是：如果数字的长度是偶数，则当反转数和x相等时，达到一半的位置。
	for x > revertedNum { // 如果是奇数，则当反转数>x时。
		revertedNum = revertedNum*10 + x%10 
		x /= 10
	}

	return revertedNum == x || revertedNum/10 == x
}
