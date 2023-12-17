package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func printGrid(grid [][]byte) {
	height := len(grid)
	for y := 0; y <= height-1; y++ {
		fmt.Println(string(grid[y][:]))
	}
	fmt.Println()
}

func stringGrid(grid [][]byte) string {
	res := ""
	for _, y := range grid {
		res = res + string(y)
	}
	return res
}

func moveStone(grid [][]byte, x int, y int, way string) [][]byte {
	height := len(grid)
	width := len(grid[0])
	switch way {
	case "n":
		if y-1 < 0 {
			return grid
		}
		if grid[y-1][x] != 46 {
			return grid
		}
		grid[y][x] = 46
		grid[y-1][x] = 79
		grid = moveStone(grid, x, y-1, "n")
	case "e":
		if x+1 >= width {
			return grid
		}
		if grid[y][x+1] != 46 {
			return grid
		}
		grid[y][x] = 46
		grid[y][x+1] = 79
		grid = moveStone(grid, x+1, y, "e")
	case "s":
		if y+1 >= height {
			return grid
		}
		if grid[y+1][x] != 46 {
			return grid
		}
		grid[y][x] = 46
		grid[y+1][x] = 79
		grid = moveStone(grid, x, y+1, "s")
	case "w":
		if x-1 < 0 {
			return grid
		}
		if grid[y][x-1] != 46 {
			return grid
		}
		grid[y][x] = 46
		grid[y][x-1] = 79
		grid = moveStone(grid, x-1, y, "w")
	}
	return grid
}

func calcScore(grid [][]byte) int {
	mul := 1
	result := 0
	for i := len(grid) - 1; i >= 0; i-- {
		count := 0
		for _, s := range grid[i] {
			if s == 79 {
				count++
			}
		}
		result = result + (mul * count)
		mul++
	}
	return result
}

func part1() {
	data, _ := os.ReadFile("input.txt")
	grid := bytes.Fields(data)

	height := len(grid)
	width := len(grid[0])

	for y := 0; y <= height-1; y++ {
		for x := 0; x <= width-1; x++ {
			if grid[y][x] == 79 {
				grid = moveStone(grid, x, y, "n")
			}
		}
	}
	fmt.Println()
	mul := 1
	result := 0
	for i := len(grid) - 1; i >= 0; i-- {
		count := 0
		for _, s := range grid[i] {
			if s == 79 {
				count++
			}
		}
		result = result + (mul * count)
		mul++
	}
	fmt.Println("Part 1: ", result)
}

func part2() {
	data, _ := os.ReadFile("input.txt")
	grid := bytes.Fields(data)

	height := len(grid)
	width := len(grid[0])

	resultarray := make([]string, 0)
	resultmap := make(map[string][][]byte)
	loop := false
	startloop := -1
	endloop := -1
	lasti := -1
	idx := 0
	for cycle := 0; cycle <= 1000000000; cycle++ {
		if loop {
			idx++
			if idx > endloop {
				idx = startloop
			}
			continue
		} else {
			idx++
		}
		for i, _ := range resultarray {
			if strings.Contains(stringGrid(grid), resultarray[i]) {
				if i == startloop {
					endloop = lasti
					loop = true
				}
				if startloop == -1 {
					startloop = i
				}
				lasti = i
				break
			}
			if loop {
				break
			}
		}
		// if cycle%1000000 == 0 {
		// 	fmt.Println(time.Now(), " ", cycle)
		// }
		chkstr := stringGrid(grid)
		resultarray = append(resultarray, chkstr)
		resultmap[chkstr] = make([][]byte, len(grid))
		for c := range grid {
			resultmap[chkstr][c] = make([]byte, len(grid[c]))
			copy(resultmap[chkstr][c], grid[c])
		}
		for y := 0; y <= height-1; y++ {
			for x := 0; x <= width-1; x++ {
				if grid[y][x] == 79 {
					grid = moveStone(grid, x, y, "n")
				}
			}
		}
		// fmt.Println("n")
		// printGrid(grid)
		for x := 0; x <= width-1; x++ {
			for y := height - 1; y >= 0; y-- {
				if grid[y][x] == 79 {
					grid = moveStone(grid, x, y, "w")
				}
			}
		}
		// fmt.Println("w")
		// printGrid(grid)
		for y := height - 1; y >= 0; y-- {
			for x := 0; x <= width-1; x++ {
				if grid[y][x] == 79 {
					grid = moveStone(grid, x, y, "s")
				}
			}
		}
		// fmt.Println("s")
		// printGrid(grid)
		for y := 0; y <= height-1; y++ {
			for x := width - 1; x >= 0; x-- {
				if grid[y][x] == 79 {
					grid = moveStone(grid, x, y, "e")
				}
			}
		}
		// fmt.Println("e")
		// printGrid(grid)
		// if cycle%1000000 == 0 {
		// 	fmt.Println(time.Now(), " ", cycle)
		// }
	}
	gridstring := resultarray[idx+1]
	grid = resultmap[gridstring]
	fmt.Println("Part 2: ", calcScore(grid))
}

func main() {
	part1()
	part2()
}
