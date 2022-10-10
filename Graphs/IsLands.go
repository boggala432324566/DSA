package main

import "fmt"

var queue []Direction

type Direction struct {
	r int
	c int
}

func EnQueue(ele Direction) {
	queue = append(queue, ele)
}
func DeQueue() Direction {
	ele := queue[0]
	queue = queue[1:]
	return ele
}
func isEmpty() bool {
	return len(queue) == 0
}

func (g *Graph) IsLands() {
	grid := [4][4]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {0, 0, 1, 1}}
	isLands := g.NumIsLands(grid)
	fmt.Printf("Number of IsLands:%d ", isLands)

}

func (g *Graph) NumIsLands(grid [4][4]int) int {
	row := len(grid)
	col := len(grid[0])
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				count++
				EnQueue(Direction{r: i, c: j})
				grid[i][j] = 2
				for !isEmpty() {
					ele := DeQueue()
					r := ele.r + 1
					c := ele.c
					if r >= 0 && r < row && c >= 0 && c < col && grid[r][c] == 1 {
						EnQueue(Direction{r: r, c: c})
						grid[r][c] = 2
					}
					r = ele.r - 1
					c = ele.c
					if r >= 0 && r < row && c >= 0 && c < col && grid[r][c] == 1 {
						EnQueue(Direction{r: r, c: c})
						grid[r][c] = 2
					}
					r = ele.r
					c = ele.c + 1
					if r >= 0 && r < row && c >= 0 && c < col && grid[r][c] == 1 {
						EnQueue(Direction{r: r, c: c})
						grid[r][c] = 2
					}
					r = ele.r
					c = ele.c - 1
					if r >= 0 && r < row && c >= 0 && c < col && grid[r][c] == 1 {
						EnQueue(Direction{r: r, c: c})
						grid[r][c] = 2
					}
				}

			}
		}
	}
	return count
}
