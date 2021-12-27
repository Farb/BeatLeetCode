package main

import "strings"

//https://leetcode-cn.com/problems/maximum-number-of-words-found-in-sentences/
//找出每个句子空格的个数，句子单词的个数等于空格数+1
func mostWordsFound_blank(sentences []string) int {
	max := 0
	for _, sentence := range sentences {
		count := strings.Count(sentence, " ")
		if count > max {
			max = count
		}
	}
	return max + 1
}

//计算切分后的切片长度
func mostWordsFound(sentences []string) int {
	max := 0
	for _, sentence := range sentences {
		count := len(strings.Split(sentence, " "))
		if count > max {
			max = count
		}
	}
	return max 
}
