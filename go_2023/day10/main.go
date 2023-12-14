package main

import (
	"bytes"
	"fmt"
	"os"
)

func contains(ints []int, num int) bool {
	for _, v := range ints {
		if v == num {
			return true
		}
	}
	return false
}

func part1() {
	data, _ := os.ReadFile("input.txt")
	grid := bytes.Fields(data)

	n := len(grid)
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
	}

	type point struct{ x, y int }

	start := point{}
	for i := 0; i <= len(grid)-1; i++ {
		for j := 0; j <= len(grid[i])-1; j++ {
			if grid[i][j] == 83 {
				start = point{j, i}
			}
		}
	}

	work := make([][]point, 20*n)

	add := func(p point, d int) {
		// check if off grid
		if p.x < 0 || p.x >= n || p.y < 0 || p.y >= n {
			return
		}
		// add distance of point
		d += int(grid[p.x][p.y]) - '0'
		// fmt.Println(p.x, p.y)
		// printdist(dist)
		// only store smallest distances
		if dist[p.x][p.y] <= d {
			return
		}
		dist[p.x][p.y] = d
		// store end of path (next field to check)
		work[d] = append(work[d], p)
	}

	add(start, 0)

	visit := func(p point) {
		d := dist[p.x][p.y]
		if grid[p.y][p.x] == 69 {
			fmt.Println(d)
			return
		}

		fmt.Println(p)
		fmt.Println(p)
		// if p == (p[p.x - 1][p.y] - 1) {
		// 	add(point{p.x - 1, p.y}, d)
		// }
		// if (point{p.x + 1, p.y} - 1) == p {
		// 	add(point{p.x + 1, p.y}, d)
		// }
		// if (point{p.x, p.y - 1} - 1) == p {
		// 	add(point{p.x, p.y - 1}, d)
		// }
		// if (point{p.x, p.y + 1} - 1) == p {
		// 	add(point{p.x, p.y + 1}, d)
		// }
	}

	for _, w := range work {
		for _, p := range w {
			visit(p)
		}
	}

	// fmt.Println("Part 1:", result)
}

func main() {
	part1()
}
