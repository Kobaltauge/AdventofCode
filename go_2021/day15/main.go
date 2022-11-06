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

func main() {
	data, _ := os.ReadFile("sample.txt")
	grid := bytes.Fields(data)

	n := len(grid)
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
		d += int(grid[p.x][p.y]) - '0'
		// fmt.Println(p.x, p.y)
		// fmt.Println(dist[p.x][p.y], d)
		// printdist(dist)
		if dist[p.x][p.y] <= d {
			return
		}
		dist[p.x][p.y] = d
		work[d] = append(work[d], p)
		// fmt.Println(d, work[d])
	}

	// at start the first "distance" must be substracted,
	// because we add it in the function
	add(point{0, 0}, -int(grid[0][0]-'0'))

	visit := func(p point) {
		d := dist[p.x][p.y]
		if p.x == n-1 && p.y == n-1 {
			fmt.Println(d)
			os.Exit(0)
		}

		add(point{p.x - 1, p.y}, d)
		add(point{p.x + 1, p.y}, d)
		add(point{p.x, p.y - 1}, d)
		add(point{p.x, p.y + 1}, d)
	}

	for idx, w := range work {
		fmt.Println(idx, w)
		for _, p := range w {
			visit(p)
		}
	}
}
