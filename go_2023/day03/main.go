package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
)

func valcheck(data [][]byte, row int, col int, hight int, width int) bool {
	m := map[string]func([][]byte, int, int) byte{
		"n":  func(data [][]byte, row int, col int) byte { return data[row-1][col] },
		"ne": func(data [][]byte, row int, col int) byte { return data[row-1][col+1] },
		"e":  func(data [][]byte, row int, col int) byte { return data[row][col+1] },
		"se": func(data [][]byte, row int, col int) byte { return data[row+1][col+1] },
		"s":  func(data [][]byte, row int, col int) byte { return data[row+1][col] },
		"sw": func(data [][]byte, row int, col int) byte { return data[row+1][col-1] },
		"w":  func(data [][]byte, row int, col int) byte { return data[row][col-1] },
		"nw": func(data [][]byte, row int, col int) byte { return data[row-1][col-1] },
	}

	ways := []string{"n", "ne", "e", "se", "s", "sw", "w", "nw"}

	if ((m["e"](data, row, col)) >= 48) && (m["e"](data, row, col) <= 57) {
		return valcheck(data, row, col+1, hight, width)
	} else {
		for _, way := range ways {
			if !(row < 0 || row > hight || col < 0 || col > width) {
				numbers := []byte{46, 10, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
				if !(slices.Contains(numbers, m[way](data, row, col))) {
					return true
				}
			}

		}
		return false
	}
}

func part1() {
	data, _ := os.ReadFile("sample.txt")
	grid := bytes.Fields(data)
	hight := len(grid)
	width := len(grid[0])
	res := make([][]byte, hight)
	for i := range res {
		res[i] = make([]byte, width)
	}

	for row, _ := range grid {
		for col, _ := range grid[row] {
			if !(valcheck(grid, row, col, hight, width)) {
				//res[idx] = byte(".")
				res[row][col] = 46
			} else {
				res[row][col] = grid[row][col]
			}
		}
		fmt.Println(grid)
		fmt.Println(res)
	}
	fmt.Println("Part 1:", grid)
}

func main() {
	part1()
}
