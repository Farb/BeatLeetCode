package main

import "math"

//https://leetcode-cn.com/problems/detect-squares/
type DetectSquares struct {
	pointsMap map[[2]int]int //一维数组作为key,出现的次数作为value
}

func Constructor() DetectSquares {
	return DetectSquares{pointsMap: make(map[[2]int]int)}
}

func (this *DetectSquares) Add(point []int) {
	newPoint := [2]int{point[0], point[1]}
	this.pointsMap[newPoint]++
}

func (this *DetectSquares) Count(point []int) (res int) {
	x, y := point[0], point[1]
	for p, cnt := range this.pointsMap {
		//找纵坐标y相同的点,但x不同，这两点可以确定一条水平线
		newX, newY := p[0], p[1]
		if newY == y && newX != x {
			res += cnt * this.findPoints([2]int{newX, y + abs(x-newX)}, [2]int{x, y + abs(x-newX)}) //找水平线以上的两点存在的个数
			res += cnt * this.findPoints([2]int{newX, y - abs(x-newX)}, [2]int{x, y - abs(x-newX)}) //找水平线以下的两点存在的个数
		}
	}
	return
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

// 另外两点存在的点数之积，即为水平线上方或下方可以组成正方形的点数
func (this *DetectSquares) findPoints(point1, point2 [2]int) (cnt int) {
	p1Cnt:= this.pointsMap[point1]
	p2Cnt := this.pointsMap[point2]
	return p1Cnt * p2Cnt
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
