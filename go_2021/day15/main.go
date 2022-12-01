package main

import (
	"bytes"
	"fmt"
	"os"
)

func printdist(dist [][]int) {
	for i := range dist {
		fmt.Println(dist[i])
	}
	fmt.Println("#######")
}

func part1() {
	data, _ := os.ReadFile("sample.txt")
	grid := bytes.Fields(data)

	n := len(grid)
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		// fill fields with high number to compare with
		for j := range dist[i] {
			dist[i][j] = 1e9
		}
	}

	type point struct{ x, y int }
	// create work array to store the next fields to check
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

	// at start the first "distance" must be substracted,
	// because we add it in the function
	add(point{0, 0}, -int(grid[0][0]-'0'))

	visit := func(p point) {
		d := dist[p.x][p.y]
		if p.x == n-1 && p.y == n-1 {
			fmt.Println(d)
			return
		}

		add(point{p.x - 1, p.y}, d)
		add(point{p.x + 1, p.y}, d)
		add(point{p.x, p.y - 1}, d)
		add(point{p.x, p.y + 1}, d)
	}

	for _, w := range work {
		for _, p := range w {
			visit(p)
		}
	}
}

func part2() {
	data, _ := os.ReadFile("input.txt")
	gridraw := bytes.Fields(data)
	m := len(gridraw)
	n := m * 5

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
		if i < m {
			for j := 0; j < n; j++ {
				if j >= m {
					value := int(grid[i][j-m])
					value = value + 1
					if value > 9 {
						value -= 9
					}
					grid[i][j] = value
				} else {
					value := int(gridraw[i][j%m]) - '0'
					grid[i][j] = value
				}
			}
		} else {
			for j := 0; j < n; j++ {
				value := grid[i-m][j]
				value = value + 1
				if value > 9 {
					value -= 9
				}
				grid[i][j] = value
			}
		}
	}

	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = 1e9
		}
	}

	type point struct{ x, y int }
	work := make([][]point, 20*n)

	add := func(p point, d int) {
		// check if off grid
		if p.x < 0 || p.x >= n || p.y < 0 || p.y >= n {
			return
		}
		// add distance of point
		d += int(grid[p.x][p.y])
		// fmt.Println(p.x, p.y)
		// printdist(dist)
		if dist[p.x][p.y] <= d {
			return
		}
		dist[p.x][p.y] = d
		work[d] = append(work[d], p)
	}

	// at start the first "distance" must be substracted,
	// because we add it in the function
	add(point{0, 0}, -int(grid[0][0]))

	visit := func(p point) {
		d := dist[p.x][p.y]
		if p.x == n-1 && p.y == n-1 {
			fmt.Println(d)
			return
		}

		add(point{p.x - 1, p.y}, d)
		add(point{p.x + 1, p.y}, d)
		add(point{p.x, p.y - 1}, d)
		add(point{p.x, p.y + 1}, d)
	}

	for _, w := range work {
		for _, p := range w {
			visit(p)
		}
	}
}

func main() {
	part1()
	part2()
}
