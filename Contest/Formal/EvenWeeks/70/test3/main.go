package main

import "sort"

type Goods struct {
	distance int
	cost     int
	row      int
	col      int
}

//BFS :如果找到的货物等于k了，则停止搜索

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) (res [][]int) {
	goods := make([]Goods, 0, k)
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	queue := [][2]int{{start[0], start[1]}}
	distance := 0
bfs:
	for len(queue) > 0 {
		layerCount := len(queue)
		for i := 0; i < layerCount; i++ {
			curPoint := queue[0]
			queue = queue[1:]
			x, y := curPoint[0], curPoint[1]

			if !inGrid(grid, x, y) || grid[x][y] <= 0 {
				continue
			}
			if grid[x][y] != 1 && inPriceRange(pricing, grid[x][y]) {
				goods = append(goods, Goods{distance: distance, cost: grid[x][y], row: x, col: y})
				if len(goods) == k {
					break bfs
				}
			}

			grid[x][y] = -1
			//访问上下左右4个方格
			for i := 0; i < len(dx); i++ {
				newX, newY := x+dx[i], y+dy[i]
				newPoints := [2]int{newX, newY}
				queue = append(queue, newPoints)
			}
		}
		distance++
	}
	sort.Slice(goods, func(i, j int) bool {
		if goods[i].distance < goods[j].distance {
			return true
		}
		if goods[i].distance > goods[j].distance {
			return false
		}
		if goods[i].cost < goods[j].cost {
			return true
		}
		if goods[i].cost > goods[j].cost {
			return false
		}
		if goods[i].row < goods[j].row {
			return true
		}
		if goods[i].row > goods[j].row {
			return false
		}
		if goods[i].col < goods[j].col {
			return true
		}
		if goods[i].col > goods[j].col {
			return false
		}
		return false

	})
	for _, good := range goods {
		res = append(res, []int{good.row, good.col})
	}
	return
}

func inGrid(grid [][]int, row, col int) bool {
	return 0 <= row && row < len(grid) && 0 <= col && col < len(grid[0])
}

func inPriceRange(pricing []int, cost int) bool {
	return pricing[0] <= cost && cost <= pricing[1]
}
