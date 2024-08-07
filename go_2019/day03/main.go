package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct{ x, y int }

func points(path []string, pos point) []point {
	var points []point
	for _, way := range path {
		dir := way[:1]
		temp := strings.Replace(way[1:], "\r", "", -1)
		length, err := strconv.Atoi(temp)
		if err != nil {
			fmt.Println(err)
			break
		}
		// if dir == "U" {
		// 	fmt.Println(pos.y + length)
		// } else if dir == "R" {
		// 	fmt.Println(pos.x + length)
		// } else if dir == "D" {
		// 	fmt.Println(pos.y - length)
		// } else if dir == "L" {
		// 	fmt.Println(pos.x - length)
		// }
		for p := 0; p < length; p++ {
			if dir == "U" {
				pos.y = pos.y + 1
			} else if dir == "R" {
				pos.x = pos.x + 1
			} else if dir == "D" {
				pos.y = pos.y - 1
			} else if dir == "L" {
				pos.x = pos.x - 1
			}
			points = append(points, pos)
		}
	}
	return points
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(file string) {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(content), "\n")
	path1 := strings.Split(lines[0], ",")
	path2 := strings.Split(lines[1], ",")

	pos := point{0, 0}
	cable1 := points(path1, pos)
	pos = point{0, 0}
	cable2 := points(path2, pos)

	var intersect []point

	for _, i := range cable1 {
		for _, j := range cable2 {
			// if i.x == j.x && i.y == j.y {
			if i == j {
				intersect = append(intersect, i)
			}
		}
	}

	solution := 10000
	for _, p := range intersect {
		newx := Abs(p.x)
		newy := Abs(p.y)
		if newx+newy <= solution {
			solution = newx + newy
		}
		// if newx <= solution.x && newy <= solution.y {
		// 	solution.x, solution.y = newx, newy
		// }
	}
	fmt.Println(solution) // rigth 489
}

func contains(s []point, elem point) bool {
	for _, i := range s {
		if i == elem {
			return true
		}
	}
	return false
}

func part2(file string) {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(content), "\n")
	path1 := strings.Split(lines[0], ",")
	path2 := strings.Split(lines[1], ",")

	pos := point{0, 0}
	cable1 := points(path1, pos)
	pos = point{0, 0}
	cable2 := points(path2, pos)

	var intersect []point

	for _, i := range cable1 {
		for _, j := range cable2 {
			// if i.x == j.x && i.y == j.y {
			if i == j {
				intersect = append(intersect, i)
			}
		}
	}

	br := false
	var cross point
	// if cable cross, the shorter way has to be counted
	// this is nocht implemented
	way1 := make(map[point]int)
	way1[point{0, 0}] = 0
	for i, p := range cable1 {
		way1[p] = i + 1
	}
	way1len := 0
	step := 0
	for _, v := range cable1 {
		if _, exists := way1[v]; exists {
			way1len = way1[v]
		}
		for _, i := range intersect {
			if cable1[step].x == i.x && cable1[step].y == i.y {
				cross = cable1[step]
				br = true
				break
			}
		}
		if br {
			break
		}
		way1len++
	}

	br = false
	way2 := make(map[point]int)
	way2[point{0, 0}] = 0
	way2len := 0
	for step := 1; step < len(cable2); step++ {
		if _, exists := way2[cable2[step]]; exists {
			way1len = way2[cable2[step]]
		} else {
			way2len++
		}
		way2[cable2[step]] = step
		if cable2[step].x == cross.x && cable2[step].y == cross.y {
			way2len = way2len + step + 1
			br = true
			break
		}
		if br {
			break
		}
	}
	fmt.Println(way1len + way2len)
}

func main() {
	// part1("sample.txt")
	// part1("exp1.txt")
	// part1("exp2.txt")
	// part1("input.txt")
	// part2
	// fmt.Println("30!")
	// part2("sample.txt")
	// fmt.Println("---")
	// fmt.Println("610!")
	// part2("exp1.txt")
	// fmt.Println("---")
	// fmt.Println("410!")
	// part2("exp2.txt")
	part2("sample2.txt")
	fmt.Println("---")
	fmt.Println("not 178848 178846 178844 137012 178488 376180")
	part2("input.txt")
}
