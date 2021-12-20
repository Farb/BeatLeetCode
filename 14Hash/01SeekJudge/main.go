package main

//https://leetcode-cn.com/problems/find-the-town-judge/
// 自己根据哈希表实现，效率有点低
func findJudge_hashTable(n int, trust [][]int) int {
	if len(trust) == 0 { //特殊情况：当信任数组为0时，说明没有信任关系，若n=1,则这个人肯定时法官；若n>1,则找不到法官
		if n == 1 {
			return 1
		}
		return -1
	}
	beTrustedMap := make(map[int]int, n) //记录每个人被信任的次数
	trustMap := make(map[int]int, n)     //记录每个人信任别人的次数
	for i := 0; i < len(trust); i++ {
		trust, beTrusted := trust[i][0], trust[i][1]
		trustMap[trust] += 1         //记录trust这个人信任了多少次人
		beTrustedMap[beTrusted] += 1 // 记录beTrusted这个人被信任的次数
	}
	possibleJudges := make([]int, 0, n/3)
	for k, v := range beTrustedMap { //找到可能的法官
		if v == n-1 {
			possibleJudges = append(possibleJudges, k)
		}
	}
	for _, no := range possibleJudges {
		if _, ok := trustMap[no]; !ok { //根据题意，只有1个人既满足被每个人信任，又满足不信任任何人，这个人就是法官
			return no
		}
	}
	return -1
}

// 有向图思路，每个人看作一个结点，trust中的每个元素看作一条边
// 每个信赖关系的第一个结点看作出度，第二个结点看作入度；仅当出度=0，入读=n-1的结点才是法官
func findJudge(n int, trust [][]int) int {
	if len(trust) == 0 { //特殊情况：当信任数组为0时，说明没有信任关系，若n=1,则这个人肯定时法官；若n>1,则找不到法官
		if n == 1 {
			return 1
		}
		return -1
	}

	in, out := make([]int, n+1), make([]int, n+1)
	for _, t := range trust {
		out[t[0]]++
		in[t[1]]++
	}

	for i := 1; i <= n; i++ {
		if out[i] == 0 && in[i] == n-1 {
			return i
		}
	}
	return -1
}
