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

	result := 0
	br := false
	var cross point
	// if cable cross, the shorter way has to be counted
	// this is nocht implemented
	for step := 1; step < len(cable1); step++ {
		for _, i := range intersect {
			if cable1[step].x == i.x && cable1[step].y == i.y {
				result = result + step + 1
				cross = cable1[step]
				br = true
				break
			}
		}
		if br {
			break
		}
	}

	br = false
	for step := 1; step < len(cable2); step++ {
		if cable2[step].x == cross.x && cable2[step].y == cross.y {
			result = result + step + 1
			br = true
			break
		}
		if br {
			break
		}
	}
	fmt.Println(result)
}

func main() {
	// part1("sample.txt")
	// part1("exp1.txt")
	// part1("exp2.txt")
	part1("input.txt")
	// part2
	fmt.Println("30!")
	part2("sample.txt")
	fmt.Println("---")
	fmt.Println("610!")
	part2("exp1.txt")
	fmt.Println("---")
	fmt.Println("410!")
	part2("exp2.txt")
	fmt.Println("---")
	fmt.Println("not 178848 178846 178844 137012 178488")
	part2("input.txt")
}
