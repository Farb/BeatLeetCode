package main

// 单调栈:创建一个栈存储温度数组的下标，遍历温度数组，如果栈为空，下标i入栈；
// 如果栈不为空，下标i对应的温度小于栈顶值在温度数组下标中的值，则将i入栈；
// 如果下标i对应的温度>=栈中的值对应的温度，将栈中的所有数值依次进行计算差值即是相隔天数。
// 可以看出，该栈是一个温度递减的单调栈
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures))
	for i, temp := range temperatures {
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top] = i - top
		}
		stack = append(stack, i)
	}
	return res
}
