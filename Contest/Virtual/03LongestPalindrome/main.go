package main

//https://leetcode-cn.com/contest/biweekly-contest-69/problems/longest-palindrome-by-concatenating-two-letter-words/
func longestPalindrome(words []string) int {
	m := make(map[string]int) //如果一个ab字符串的出现的次数为m,另一个ba出现的次数n,且m>n，则能组成的最大回文长度为2*n*2
	for _, word := range words {
		m[word]++
	}
	palindromeStrCnt := 0
	sameCnt := 0
	for word := range m {
		if m[word] == 0 {
			continue
		}
		if word[0] == word[1] {
			sameCnt = max(sameCnt, m[word])
			continue
		}
		minCnt := min(m[word], m[reverse(word)])
		m[reverse(word)] = 0
		palindromeStrCnt += minCnt * 2
	}
	res := palindromeStrCnt*2 + sameCnt*2
	return res
}

func reverse(str string) string {
	arr := []byte(str)
	arr[0], arr[1] = arr[1], arr[0]
	return string(arr)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
