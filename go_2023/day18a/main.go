package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var dirs = map[string][]int{
	"U": {-1, 0},
	"D": []int{1, 0},
	"L": []int{0, -1},
	"R": []int{0, 1},
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	points := [][]int{{0, 0}}
	for fileScanner.Scan() {
		x := fileScanner.Text()

		line := strings.FieldsFunc(x, func(r rune) bool {
			if r == '(' || r == ')' || r == '#' || r == ' ' {
				return true
			}
			return false
		})
		dir := dirs[line[0]]
		moves, _ := strconv.Atoi(line[1])
		row := points[len(points)-1][0]
		column := points[len(points)-1][1]

		points = append(points, []int{row + (dir[0] * moves), column + (dir[1] * moves)})

		// fmt.Println(points)

	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	miny := points[0][0]
	// maxy := points[0][len(points[0]-1)]
	fmt.Println("end", miny)
}

func main() {
	fmt.Println("Part 1")
	part1()
	// fmt.Println("Part 2")
	// part2()
}
