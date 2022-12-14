package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	cave := make([][]string, 1000)
	for i := 0; i < 1000; i++ {
		cave[i] = make([]string, 1000)
	}
	top := 0
	bottom := 0
	left := 100000
	right := 0

	for fileScanner.Scan() {
		raw := strings.Fields(fileScanner.Text())
		startx := 0
		starty := 0
		endx := 0
		endy := 0
		for idx, koor := range raw {
			if koor != "->" {
				x, _ := strconv.Atoi(strings.Split(koor, ",")[0])
				y, _ := strconv.Atoi(strings.Split(koor, ",")[1])
				if idx == 0 {
					endx = x
					endy = y
					continue
				}
				startx = endx
				starty = endy
				endx = x
				endy = y
				fx := 0
				tx := 0
				fy := 0
				ty := 0
				if startx < endx {
					fx = startx
					tx = endx
				} else {
					tx = startx
					fx = endx
				}
				if starty < endy {
					fy = starty
					ty = endy
				} else {
					fy = endy
					ty = starty
				}
				if startx < left {
					left = startx
				}
				if endx > right {
					right = endx
				}
				if starty < top {
					top = starty
				}
				if endy > bottom {
					bottom = endy
				}
				for i := fy; i <= ty; i++ {
					for j := fx; j <= tx; j++ {
						cave[i][j] = "#"
					}
				}
			}
		}

	}
	for i := top; i <= bottom; i++ {
		for j := left; j <= right; j++ {
			if cave[i][j] == "#" {
				fmt.Print(cave[i][j])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	// for i := range dist {
	// 	dist[i] = make([]int, n)
	// 	// fill fields with high number to compare with
	// 	for j := range dist[i] {
	// 		dist[i][j] = 1e9
	// 	}
	// }

	// type point struct{ x, y int }
	// // create work array to store the next fields to check
	// work := make([][]point, 20*n)

	// add := func(p point, d int) {
	// 	// check if off grid
	// 	if p.x < 0 || p.x >= n || p.y < 0 || p.y >= n {
	// 		return
	// 	}
	// 	// add distance of point
	// 	d += int(grid[p.x][p.y]) - '0'
	// 	// fmt.Println(p.x, p.y)
	// 	// printdist(dist)
	// 	// only store smallest distances
	// 	if dist[p.x][p.y] <= d {
	// 		return
	// 	}
	// 	dist[p.x][p.y] = d
	// 	// store end of path (next field to check)
	// 	work[d] = append(work[d], p)
	// }

	// // at start the first "distance" must be substracted,
	// // because we add it in the function
	// add(point{0, 0}, -int(grid[0][0]-'0'))

	// visit := func(p point) {
	// 	d := dist[p.x][p.y]
	// 	if p.x == n-1 && p.y == n-1 {
	// 		fmt.Println(d)
	// 		return
	// 	}

	// 	add(point{p.x - 1, p.y}, d)
	// 	add(point{p.x + 1, p.y}, d)
	// 	add(point{p.x, p.y - 1}, d)
	// 	add(point{p.x, p.y + 1}, d)
	// }

	// for _, w := range work {
	// 	for _, p := range w {
	// 		visit(p)
	// 	}
	// }
}

func main() {
	part1()
}
