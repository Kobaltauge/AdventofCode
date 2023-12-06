package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func valcheck(data [][]byte, row int, col int, hight int, width int) bool {
	m := map[string]func(int, int) (int, int){
		"n":  func(r int, c int) (int, int) { return r - 1, c },
		"ne": func(r int, c int) (int, int) { return r - 1, c + 1 },
		"e":  func(r int, c int) (int, int) { return r, c + 1 },
		"se": func(r int, c int) (int, int) { return r + 1, c + 1 },
		"s":  func(r int, c int) (int, int) { return r + 1, c },
		"sw": func(r int, c int) (int, int) { return r + 1, c - 1 },
		"w":  func(r int, c int) (int, int) { return r, c - 1 },
		"nw": func(r int, c int) (int, int) { return r - 1, c - 1 },
	}

	ways := []string{"n", "ne", "e", "se", "s", "sw", "w", "nw"}

	newrow, newcol := m["e"](row, col)
	if newrow >= 0 && newrow < hight && newcol >= 0 && newcol < width {
		if data[newrow][newcol] >= 48 && data[newrow][newcol] <= 57 {
			return valcheck(data, row, col+1, hight, width)
		} else {
			for _, way := range ways {
				newrow, newcol := m[way](row, col)
				if newrow >= 0 && newrow < hight && newcol >= 0 && newcol < width {
					numbers := []byte{46, 10, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
					if !(slices.Contains(numbers, data[newrow][newcol])) {
						return true
					}
				}

			}
			return false
		}
	}
	return false
}

func part1() {
	result := 0
	data, _ := os.ReadFile("sample.txt")
	grid := bytes.Fields(data)
	hight := len(grid)
	width := len(grid[0])
	res := make([][]byte, hight)
	for i := range res {
		res[i] = make([]byte, width)
	}

	for row := range grid {
		for col := range grid[row] {
			if !(valcheck(grid, row, col, hight, width)) {
				//res[idx] = byte(".")
				res[row][col] = 46
			} else {
				res[row][col] = grid[row][col]
			}
		}
	}
	for idx := 0; idx < len(res); idx++ {
		tmp := strings.Split(string(res[idx]), ".")
		fmt.Println(res)
		fmt.Println(tmp)
		for _, x := range tmp {
			r, _ := strconv.Atoi(x)
			result = result + r
		}
	}
	fmt.Println("Part 1:", result)
	//2375 to lows
}

func main() {
	part1()
}
