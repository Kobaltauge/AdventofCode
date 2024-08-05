package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cube struct {
	x, y  int
	color string
}

func part1() {
	f, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	x := 0
	y := 0
	minx := 0
	maxx := 0
	miny := 0
	maxy := 0
	var field []cube

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	for fileScanner.Scan() {
		text := fileScanner.Text()

		line := strings.FieldsFunc(text, func(r rune) bool {
			if r == '(' || r == ')' || r == ' ' || r == '#' {
				return true
			}
			return false
		})
		dif := 0
		if line[0] == "U" {
			dif, _ = strconv.Atoi(line[1])
			for c := y; c >= y-dif; c-- {
				field = append(field, cube{x: x, y: c, color: line[2]})
			}
			if y-dif < miny {
				miny = y - dif
			}
		}
		if line[0] == "D" {
			dif, _ = strconv.Atoi(line[1])
			for c := y; c <= y+dif; c++ {
				field = append(field, cube{x: x, y: c, color: line[2]})
			}
			if y+dif > maxy {
				maxy = y + dif
			}
		}
		if line[0] == "L" {
			dif, _ = strconv.Atoi(line[1])
			for c := x; c >= x-dif; c-- {
				field = append(field, cube{x: x, y: c, color: line[2]})
			}
			if x-dif < minx {
				minx = x - dif
			}
		}
		if line[0] == "R" {
			dif, _ = strconv.Atoi(line[1])
			for c := x; c <= x+dif; c++ {
				field = append(field, cube{x: x, y: c, color: line[2]})
			}
			if x+dif < maxx {
				maxx = y + dif
			}
		}

		fmt.Println(x, y)
	}
	grid := make([][]cube{})

	f.Close()
}

// func part2() {
// 	knods := make(map[string]knot)
// 	f, err := os.Open("input.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	starts := make([]string, 0)

// 	fileScanner := bufio.NewScanner(f)
// 	fileScanner.Split(bufio.ScanLines)
// 	fileScanner.Scan()
// 	way := []rune(fileScanner.Text())
// 	fileScanner.Scan()
// 	for fileScanner.Scan() {
// 		x := fileScanner.Text()

// 		line := strings.FieldsFunc(x, func(r rune) bool {
// 			if r == '=' || r == '(' || r == ')' || r == ',' || r == ' ' {
// 				return true
// 			}
// 			return false
// 		})
// 		knods[line[0]] = knot{left: line[1], right: line[2]}
// 		if strings.HasSuffix(line[0], "A") {
// 			starts = append(starts, line[0])
// 		}

// 	}
// 	f.Close()
// 	res := make([]int, 0)
// 	for _, n := range starts {
// 		steps := 0
// 		for {
// 			for _, w := range way {
// 				if w == 76 {
// 					n = knods[n].left
// 					steps++
// 				} else {
// 					n = knods[n].right
// 					steps++
// 				}
// 				if strings.HasSuffix(n, "Z") {
// 					break
// 				}
// 			}
// 			if strings.HasSuffix(n, "Z") {
// 				res = append(res, steps)
// 				break
// 			}
// 		}
// 	}
// 	result := LCM(res[0], res[1], res...)
// 	fmt.Println(result)

func main() {
	fmt.Println("Part 1")
	part1()
	// fmt.Println("Part 2")
	// part2()
}
